package todo

import (
	"todolist/src/clients"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}

func TodoColl() *mongo.Collection {
	return clients.DB.Collection("todos")
}
