package go_movies_crud

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

// Director This director struct is associated with Director in the movie
//because every movie has a director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// a slice of movies
var movies []Movie
/*5 routes: get all, get by Id, create, update, delete
5 functions: get movies, get movie, create, update, delete
5 Endpoints: /movies, /movies/id, /movies, /movies/id, /movies
 5 methods: GET, GET, POST, PUT DELETE
 */

func getMovies(w http.ResponseWriter, r *http.Request)  {
	//struct needs to be able to convert the json coming into his own format
	w.Header().Set("Content-Type", "application/json")
	//encode w: i.e, encode your response into json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request)  {
	 w.Header().Set("Content-Type", "application/json")
	 //id passed from post man will go as a param into this function
	 //the param will be present inside mux.vars.
	 // r-request
	 params:=mux.Vars(r)
	for index, item:= range movies{
		//id sent from postman/frontend was obtained as a request, then it was stored in params, hence the logic below
		if item.ID==params["id"] {

		}
	}
}


func main() {
	r:=mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "The Movie", Director: &Director {Firstname: "Young", Lastname: "Cho"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23456", Title: "The Engineer", Director: &Director{Firstname: "Amy", Lastname: "Kai"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting server and Port at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}






