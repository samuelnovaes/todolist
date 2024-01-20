package main

import (
	"os"
	"todolist/src/clients"
	"todolist/src/entities/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	err := clients.ConnectMongoDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	todo.Routes(engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
