package todo

import (
	"net/http"
	"todolist/src/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodosHandler(ctx *gin.Context) {
	todos, err := GetTodos()
	if err != nil {
		utils.Error.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
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
		if err == mongo.ErrNoDocuments {
			ctx.Status(http.StatusNotFound)
			return
		}
		utils.Error.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
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
		utils.Error.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
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
		utils.Error.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
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
		utils.Error.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
}
