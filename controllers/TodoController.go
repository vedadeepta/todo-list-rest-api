package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vedadeepta/go-rest-api/models"
	"github.com/vedadeepta/go-rest-api/controllers/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var tempTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&tempTodo)
	if err != nil {
		fmt.Printf("Invalid request")
	}
	models.TodoSlice = append(models.TodoSlice ,tempTodo)	
	w.WriteHeader(200)
	response, _ := json.Marshal(tempTodo)
	w.Write(response)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var allTodos = models.TodoSlice
	response, _ := json.Marshal(allTodos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	for _, t := range models.TodoSlice {
		if t.ID == params["id"] {
			todo = t
			break
		}
	}
	utils.SendJson(w, 200, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var tempTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&tempTodo)
	if err != nil {
		fmt.Printf("Invalid request")
	}
	for i, todo := range models.TodoSlice {
		if tempTodo.ID == todo.ID {
			models.Update(&models.TodoSlice[i], tempTodo)
			break
		}
	}
	utils.SendJson(w, 200, map[string]string{"result": "ok"})
}