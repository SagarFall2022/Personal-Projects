# The DAG object; we'll need this to instantiate a DAG
from airflow import DAG


# Operators; we need this to write tasks!
from airflow.operators.bash_operator import BashOperator


from airflow.utils.dates import days_ago

from datetime import timedelta




#defining DAG arguments
# You can override them on a per-task basis during operator initialization
default_args = {
    'owner': 'Sagar S Patil',
    'start_date': days_ago(0),
    'email': ['sagarspatil@somemail.com'],
    'email_on_failure': True,
    'email_on_retry': True,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}


# defining the DAG
dag = DAG(
    'ETL_toll_data',
    default_args=default_args,
    description='Apache Airflow Final Assignment',
    schedule_interval=timedelta(days=1),
)




# define the tasks


# define the task 'unzip_data'


unzip_data = BashOperator(
    task_id='unzip_data',
    bash_command='tar -xzf /home/project/airflow/dags/finalassignment/tolldata.tgz -C /home/project/airflow/dags/finalassignment/staging/',
    dag=dag,
)


# define the task 'extract_data_from_csv'


extract_data_from_csv = BashOperator(
    task_id='extract_data_from_csv',
    bash_command='cut -f1,2,3,4 -d"," /home/project/airflow/dags/finalassignment/staging/vehicle-data.csv > /home/project/airflow/dags/finalassignment/staging/csv_data.csv',
    dag=dag,
)




# define the task 'extract_data_from_tsv'


extract_data_from_tsv = BashOperator(
    task_id='extract_data_from_tsv',
    bash_command='cut -f5,6,7 /home/project/airflow/dags/finalassignment/staging/tollplaza-data.tsv > /home/project/airflow/dags/finalassignment/staging/tsv_data.csv',
    dag=dag,
)


# define the task 'extract_data_from_fixed_width'


extract_data_from_fixed_width = BashOperator(
    task_id='extract_data_from_fixed_width',
    bash_command='cut -c59-61,62-67 /home/project/airflow/dags/finalassignment/staging/payment-data.txt > /home/project/airflow/dags/finalassignment/staging/fixed_width_data.csv',
    dag=dag,
)


# define the task 'consolidate_data'


consolidate_data = BashOperator(
    task_id='consolidate_data',
    bash_command='paste -d "," /home/project/airflow/dags/finalassignment/staging/csv_data.csv /home/project/airflow/dags/finalassignment/staging/tsv_data.csv /home/project/airflow/dags/finalassignment/staging/fixed_width_data.csv | cut -d "," -f 1,2,3,4,5,6,7,8,9 > /home/project/airflow/dags/finalassignment/staging/extracted_data.csv',
    dag=dag,
)

transform_data = BashOperator(
    task_id='transform_data',
    bash_command='tr [:lower:] [:upper:] < /home/project/airflow/dags/finalassignment/staging/extracted_data.csv > /home/project/airflow/dags/finalassignment/staging/transformed_data.csv',
    dag=dag,
)

unzip_data >> extract_data_from_csv
unzip_data >> extract_data_from_tsv
unzip_data >> extract_data_from_fixed_width
extract_data_from_csv >> consolidate_data
extract_data_from_tsv >> consolidate_data
extract_data_from_fixed_width >> consolidate_data
consolidate_data >> transform_data