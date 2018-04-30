package models

type Todo struct {
	ID string
	Name string
	Description string
}

var TodoSlice [] Todo

func Update(todo *Todo, temp Todo) {
	if(temp.Name != "") {
		todo.Name = temp.Name
	}
	if(temp.Description != "") {
		todo.Description = temp.Description
	}
}