package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id        primitive.ObjectID `form:"id" bson:"_id"`
	Title     string             `form:"title" bson:"title"`
	Completed bool               `form:"completed" bson:"completed"`
}
