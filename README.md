# 1. Data Pipeline using Apache Airflow

This project involves the creation of a data pipeline using Apache Airflow. The pipeline consists of a DAG (Directed Acyclic Graph) designed to extract data from CSV, TSV, and fixed-width files, perform transformations on the data, and load the transformed data into a staging area. The DAG is scheduled to run on a daily basis at a specific time, and email notifications are configured to alert in case of any failures.

## Features

1. **Data Extraction**: The DAG leverages the CSV, TSV, and Fixed Width File operators provided by Apache Airflow to extract data from files in different formats. These operators handle the reading and parsing of the files, making it easy to extract data from diverse sources.

2. **Data Transformation**: A Python script is employed within the DAG to perform the necessary data transformations. This script can be customized based on the specific transformation requirements of the project. It allows for flexibility and adaptability in handling various data manipulation tasks.

3. **Staging Area**: The transformed data is loaded into a staging area using the Postgres operator in Apache Airflow. This operator facilitates the interaction with a PostgreSQL database and enables the insertion of the transformed data into the designated staging tables.

4. **Scheduling and Monitoring**: The DAG is scheduled to run on a daily basis at a specific time, ensuring regular and automated execution of the data pipeline. Apache Airflow's built-in scheduling capabilities make it easy to configure the desired execution frequency. Additionally, email notifications are configured to provide alerts in case of any failures during the execution of the pipeline.

## Installation and Configuration

1. Install Apache Airflow by following the official installation instructions provided by the Apache Airflow project.

2. Set up a PostgreSQL database or use an existing one to serve as the staging area for the transformed data. Make sure to configure the necessary credentials and connection details.

3. Configure the Apache Airflow environment by updating the necessary configurations in the Airflow configuration file (airflow.cfg). Set the appropriate values for parameters such as the executor, database connection, email settings, and timezone.

4. Create the required DAG file(s) in the Apache Airflow DAGs directory. This file should contain the DAG definition, task dependencies, and task implementations.

5. Customize the Python script used for data transformation within the DAG to match the specific requirements of your project. Update the transformations, data mappings, and any other relevant logic.

6. Start the Apache Airflow scheduler and web server to begin the execution and monitoring of the data pipeline. Use the command-line interface or the provided Airflow CLI commands to manage the execution and monitor the status of the DAG.

## Usage

1. Place the CSV, TSV, and fixed-width files containing the data to be processed in a designated directory accessible by the Apache Airflow environment.

2. The scheduled DAG will automatically trigger at the specified time and begin the data extraction, transformation, and loading process.

3. Monitor the execution of the DAG through the Apache Airflow web interface or command-line interface. Check the status of individual tasks, view logs, and track the progress of the data pipeline.

4. Upon successful execution, the transformed data will be loaded into the staging area in the configured PostgreSQL database.

5. In case of any failures during the execution, email notifications will be sent to the configured recipients, providing details of the encountered error and allowing for prompt investigation and resolution.

# 2. Monitoring Kubernetes Cluster using Prometheus and Grafana

This project aims to demonstrate the process of monitoring a Kubernetes cluster using Prometheus and Grafana. By following the steps outlined below, you will be able to create and configure a Kubernetes cluster using kind, deploy a sample deployment and service, scrape essential metrics using Prometheus, and visualize these metrics on Grafana dashboards.

## Prerequisites

Before getting started, ensure that you have the following prerequisites installed:

- Docker: Used for creating and running the Kubernetes cluster.
- Kubernetes (kubectl): The command-line tool for interacting with the Kubernetes cluster.
- kind (Kubernetes in Docker): A tool for running local Kubernetes clusters using Docker containers.
- Prometheus: The monitoring and alerting toolkit used to scrape and store metrics.
- Grafana: The open-source platform for visualizing and analyzing metrics.

## Steps

### 1. Setting up the Kubernetes Cluster

- Install Docker: Follow the official Docker installation guide for your operating system.
- Install kind: Use the official kind installation guide to install kind on your system.
- Create a Kubernetes cluster using kind: Execute the necessary commands to create a local Kubernetes cluster using kind.

### 2. Deploying a Sample Application

- Create a sample deployment and service: Use Kubernetes manifests to deploy a sample application on the Kubernetes cluster.

### 3. Configuring Prometheus

- Install Prometheus: Follow the official Prometheus installation guide to install Prometheus on the Kubernetes cluster.
- Configure Prometheus to scrape metrics: Create a Prometheus configuration file to specify the targets and metrics to scrape.

### 4. Visualizing Metrics with Grafana

- Install Grafana: Follow the official Grafana installation guide to install Grafana on the Kubernetes cluster.
- Configure Grafana: Set up Grafana to connect to Prometheus as a data source and import dashboards for visualization.
- Explore and customize Grafana dashboards: Access the Grafana UI, explore the available dashboards, and customize them to suit your needs.

## Resources

- [Official Docker Installation Guide](https://docs.docker.com/get-docker/)
- [Official kind Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [Official Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [Official Prometheus Installation Guide](https://prometheus.io/docs/prometheus/latest/installation/)
- [Official Grafana Installation Guide](https://grafana.com/docs/grafana/latest/installation/)

# 3. GO RESTful API Development

## Description

This project showcases the development of a RESTful API using the Go programming language. The main objective was to create a service that utilizes structs and slices to organize data and provides APIs for data manipulation. The project implemented commonly used HTTP methods such as GET, DELETE, UPDATE, and POST for data retrieval, deletion, updating, and creation.

To handle incoming requests and route them to the appropriate handlers, the project utilized Gorilla mux, a powerful HTTP router and dispatcher for Go. Gorilla mux allowed for defining multiple routes for different endpoints and mapping them to the corresponding handlers. The use of Gorilla mux made the process of handling and routing requests intuitive and efficient.

## Key Features

- Structs and Slices: Utilized structs and slices to organize and manage data efficiently.
- GET Method: Implemented GET API endpoints to retrieve data from the service.
- DELETE Method: Implemented DELETE API endpoints to remove data from the service.
- UPDATE Method: Implemented UPDATE API endpoints to modify existing data in the service.
- POST Method: Implemented POST API endpoints to create new data entries in the service.
- Gorilla mux: Utilized Gorilla mux as the HTTP router and dispatcher for handling incoming requests and mapping them to the appropriate handlers.
- Multiple Routes: Defined multiple routes to handle different endpoints and functionalities of the API.
- Intuitive Development: Developed the project with an emphasis on code readability and simplicity, making it easy to understand and maintain.

## Installation and Usage

1. Clone the repository: `git clone https://github.com/your-username/your-repo.git`
2. Navigate to the project directory: `cd your-repo`
3. Install the required dependencies: `go get -u github.com/gorilla/mux`
4. Run the application: `go run main.go`
5. The API endpoints are accessible at `http://localhost:8000/`




