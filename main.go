package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(web http.ResponseWriter, route *http.Request) {
	web.Header().Set("Content-Type", "application/json")
	json.NewEncoder(web).Encode(movies)
}

func deleteMovie(web http.ResponseWriter, route *http.Request) {
	web.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(web).Encode(movies)
}

func getMovie(web http.ResponseWriter, route *http.Request) {
	web.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(web).Encode(item)
			return
		}
	}
}

func createMovie(web http.ResponseWriter, route *http.Request) {
	web.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(route.Body).Decode(&movie)
	movie.ID = (strconv.Itoa(rand.Intn(100000000)))
	movies = append(movies, movie)
	json.NewEncoder(web).Encode(movie)
}

func updateMovie(web http.ResponseWriter, route *http.Request) {
	web.Header().Set("Content-Type", "application/json")
	params := mux.Vars(route)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(route.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(web).Encode(movie)
			return
		}
	}
}

func main() {
	route := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}
