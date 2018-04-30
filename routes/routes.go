package routes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vedadeepta/todo-list-rest-api/controllers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API for todolist, check Readme for further instructions")
}

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/create", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/update", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}