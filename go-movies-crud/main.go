package main

//For this project, I wont be using databases but I will be using structs and slices
import (
	"encoding/json"
	"fmt"
	"log" // I need to encode the data into json when I send it to postman
	"math/rand"
	"strconv"

	//To create random values
	"net/http" // Needed to create server

	"github.com/gorilla/mux"
)

//Struct is basically like an object in JS

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"` // unique number associated to the movie
	Title    string    `json:"title"`
	Director *Director `json:"director"` // Here * is the pointer as usual and as I create a Director then it will associted with Movie Struct

}

//Movies and Director would be associated i.e every movie has a director

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `josn:"lastname"`
}

var movies []Movie //here we have created a variable which have a slice of MOvie struct

func getMovies(w http.ResponseWriter, r *http.Request) { // Here passing of the request which from POSTMAN and w is the response which we send from this func

	w.Header().Set("Content-Type", "application/json") // We are sending the response in JSON format
	json.NewEncoder(w).Encode(movies)                  //We are encoding the w of the movies slice

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //Here we are appending the movies from the index+1 so we are deleting using the append
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies { // I am using blank identifier here as I wont be using index here and GO will give error

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Here just for practice purpose I am deleting the record and then inserting new record and If I was using DB then it would be different
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	//I am just creating two data manually below so that I can test if the APTs are working as expected
	movies = append(movies, Movie{ID: "1", Isbn: "42586", Title: "Movie one", Director: &Director{Firstname: "Sagar", Lastname: "Patil"}})
	movies = append(movies, Movie{ID: "2", Isbn: "42587", Title: "Movie two", Director: &Director{Firstname: "Shashank", Lastname: "Parameshwaran"}})
	r.HandleFunc("/movies", getMovies).Methods("GET") //I need to create five different routes and five different functions
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("STring the server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // To start the server is so easy this one line is enough
}
