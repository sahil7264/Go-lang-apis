package routes

import (
	"main/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies/delete", controller.DeleteAllMovie).Methods("DELETE")

	r.HandleFunc("/api/movies/update/{id}", controller.MarkWatched).Methods("PUT")
	r.HandleFunc("/api/movies/create", controller.InsertOneMovie).Methods("POST")
	r.HandleFunc("/api/movies/delete/{id}", controller.DeleteMovie).Methods("DELETE")

	return r
}
