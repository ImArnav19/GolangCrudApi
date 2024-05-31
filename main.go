package main

import (
	"encoding/json" //json
	"fmt"           //main
	"log"           //log results
	"math/rand"     //random values
	"net/http"      //http routes
	"strconv"       //string conv

	"github.com/gorilla/mux" //configure routes
)

type Movie struct {
	ID       string    `json:"id"` //id is as json element
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //struct for movies
}

type Director struct { //struct for director
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //getMovies json.NewEncoder(w).Encode()
	json.NewEncoder(w).Encode(movies)                  //movies as struct

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting up headers

	params := mux.Vars(r) //get params value

	for index, item := range movies { //for-each loop ,
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //appends current index to next index
			break
		}

	}

	json.NewEncoder(w).Encode(movies) //goog for sending json stuff
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie) //Decoder for r.Body

	movie.ID = strconv.Itoa(rand.Intn(100000)) //assign ID

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") //basic comb of delete and create

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000))

			movies = append(movies, movie)
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "234", Title: "Movie 1", Director: &Director{Firstname: "Virat", Lastname: "Kohli"}})
	movies = append(movies, Movie{ID: "2", Isbn: "432", Title: "2 Movie", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") //routes for all with mux
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") //learnt a funckin append() function

	fmt.Printf("Webserver running on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) //hhtp.ListenAndServer(":8000",r)
}
