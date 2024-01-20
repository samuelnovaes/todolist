package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodosHandler(ctx *gin.Context) {
	todos, err := GetTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func GetTodoHandler(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	todo, err := GetTodo(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func InsertTodoHandler(ctx *gin.Context) {
	var todo Todo

	err := ctx.Bind(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = InsertTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func RemoveTodoHandler(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = RemoveTodo(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func UpdateTodoHandler(ctx *gin.Context) {
	var todo Todo

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.Bind(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	todo.Id = id

	err = UpdateTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
