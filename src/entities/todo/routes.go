package todo

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	group := engine.Group("/todos")
	group.GET("/", GetTodosHandler)
	group.GET("/:id", GetTodoHandler)
	group.POST("/", InsertTodoHandler)
	group.DELETE("/:id", RemoveTodoHandler)
	group.PUT("/:id", UpdateTodoHandler)
}
