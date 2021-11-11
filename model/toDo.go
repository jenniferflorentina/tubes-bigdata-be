package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"desc" bson:"desc,omitempty"`
	SubTodo     []string           `json:"subTodo" bson:"subTodo,omitempty"`
	CreatedOn   int64              `json:"on" bson:"on,emitempty"`
}

type ToDos struct {
	ToDos []ToDo `json: "todos"`
}

func PrepareBsonUpdateTodo(newToDo ToDo) bson.M {
	bsonData := bson.M{}

	if newToDo.Title != "" {
		bsonData["title"] = newToDo.Title
	}

	if newToDo.Description != "" {
		bsonData["desc"] = newToDo.Description
	}

	if len(newToDo.SubTodo) != 0 {
		bsonData["subTodo"] = newToDo.SubTodo
	}

	preparedBson := bson.M{"$set": bsonData}

	return preparedBson
}
