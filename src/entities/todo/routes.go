package todo

import "github.com/gorilla/mux"

func Routes(router *mux.Router) {
	router.HandleFunc("/todos", GetTodosHandler).Methods("GET")
	router.HandleFunc("/todos", InsertTodoHandler).Methods("POST")
	router.HandleFunc("/todos/{id}", GetTodoHandler).Methods("GET")
	router.HandleFunc("/todos/{id}", RemoveTodoHandler).Methods("DELETE")
	router.HandleFunc("/todos/{id}", UpdateTodoHandler).Methods("PUT")
}
