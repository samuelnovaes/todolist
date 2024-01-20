package todo

import (
	"encoding/json"
	"net/http"
	"todolist/src/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := GetTodos()
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, todos)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	todo, err := GetTodo(id)
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, todo)
}

func InsertTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	err = InsertTodo(todo)
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendOk(w)
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	err = RemoveTodo(id)
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendOk(w)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	todo.Id = id

	err = UpdateTodo(todo)
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendOk(w)
}
