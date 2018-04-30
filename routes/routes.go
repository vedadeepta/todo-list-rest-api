package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vedadeepta/todo-list-rest-api/controllers"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/create", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}