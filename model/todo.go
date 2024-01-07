package model

import (
	"encoding/json"
	"os"
)

type Todo struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	IsCompleted bool   `json:"isCompleted"`
}

type TodosType []Todo

var Todos = TodosType{}

const jsonFileName = "todos.json"

// ...

func loadTodosFromJSON() error {
	file, err := os.ReadFile(jsonFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Todos)
	if err != nil {
		return err
	}

	return nil
}

func saveTodosToJSON() error {
	data, err := json.MarshalIndent(Todos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func newTodo(text string) (r *Todo) {
	return &Todo{
		Id:          len(Todos) + 1,
		IsCompleted: false,
		Text:        text,
	}
}

func (t TodosType) AddTodo(text string) TodosType {
	todo := newTodo(text)
	Todos = append(Todos, *todo)
	saveTodosToJSON()
	return Todos
}
func (t TodosType) DeleteTodo(id int) TodosType {
	for i, todo := range Todos {
		if todo.Id == id {
			// Delete the todo by slicing it out
			Todos = append(Todos[:i], Todos[i+1:]...)
			saveTodosToJSON()
			return Todos
		}
	}
	return Todos
}

func (t TodosType) MarkComplete(id int) TodosType {
	for i, todo := range Todos {
		if todo.Id == id {
			// Mark the todo as complete
			Todos[i].IsCompleted = true
			saveTodosToJSON()
			return Todos
		}
	}
	return Todos
}

func InitTodo() error {
	err := loadTodosFromJSON()
	if err != nil {
		return err
	}
	return nil
}
