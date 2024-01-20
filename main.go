package main

import (
	"fmt"
	"net"
	"net/http"
	"todolist/src/clients"
	"todolist/src/entities/todo"

	"github.com/gorilla/mux"
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
	srv := &mux.Router{}
	todo.Routes(srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running...")

	err = http.Serve(listener, srv)
	if err != nil {
		panic(err)
	}
}
