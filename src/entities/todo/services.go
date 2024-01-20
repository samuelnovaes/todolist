package todo

import (
	"todolist/src/clients"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func coll() *mongo.Collection {
	return clients.DB.Collection("todos")
}

func GetTodos() ([]Todo, error) {
	var todos []Todo

	cursor, err := coll().Find(clients.Ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(clients.Ctx, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodo(id primitive.ObjectID) (*Todo, error) {
	var todo Todo

	err := coll().FindOne(clients.Ctx, bson.D{{
		Key:   "_id",
		Value: id,
	}}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func InsertTodo(todo Todo) error {
	todo.Id = primitive.NewObjectID()

	_, err := coll().InsertOne(clients.Ctx, todo)
	if err != nil {
		return err
	}

	return nil
}

func RemoveTodo(id primitive.ObjectID) error {
	_, err := coll().DeleteOne(clients.Ctx, bson.D{{
		Key:   "_id",
		Value: id,
	}})
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodo(todo Todo) error {
	_, err := coll().UpdateByID(clients.Ctx, todo.Id, bson.D{{
		Key:   "$set",
		Value: todo,
	}})
	if err != nil {
		return err
	}

	return nil
}
