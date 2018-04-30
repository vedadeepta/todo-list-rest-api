package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vedadeepta/todo-list-rest-api/models"
	"github.com/vedadeepta/todo-list-rest-api/controllers/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var tempTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&tempTodo)
	if err != nil {
		fmt.Printf("Invalid request")
		utils.SendJson(w, http.StatusBadRequest, map[string]string {"error": "Invalid payload params"})
		return
	}
	for _, todo := range models.TodoSlice {
		if todo.ID == tempTodo.ID {
			fmt.Printf("Same Id todo exists")
			utils.SendJson(w, http.StatusBadRequest, map[string]string {"error": "Todo with same id exists"})
			return
		}
 	}
	models.TodoSlice = append(models.TodoSlice ,tempTodo)
	utils.SendJson(w, http.StatusOK, tempTodo)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, models.TodoSlice)
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
	if todo.ID == "" {
		utils.SendJson(w, http.StatusBadRequest, map[string]string {"error": "Todo not found"})
		return
	}
	utils.SendJson(w, http.StatusOK, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var tempTodo models.Todo
	flag := 0
	err := json.NewDecoder(r.Body).Decode(&tempTodo)
	if err != nil {
		fmt.Printf("Invalid request")
	}
	for i, todo := range models.TodoSlice {
		if tempTodo.ID == todo.ID {
			models.Update(&models.TodoSlice[i], tempTodo)
			flag = 1
			break
		}
	}
	if flag == 0 {
		utils.SendJson(w, http.StatusBadRequest, map[string]string {"error": "Todo not found"})
		return
	}
	utils.SendJson(w, http.StatusOK, map[string]string{"result": "ok"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flag := 0
	for i, todo := range models.TodoSlice {
		if todo.ID == params["id"] {
			models.TodoSlice = append(models.TodoSlice[ : i], models.TodoSlice[i+1 : ]...)
			flag = 1
			break
		}
	}
	if flag == 0 {
		utils.SendJson(w, http.StatusBadRequest, map[string]string {"error": "Todo not found"})
		return		
	}
	utils.SendJson(w, http.StatusOK, map[string]string{"result": "ok"})
}
