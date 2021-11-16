package repository

import (
	"TubesBigData/database"
	"TubesBigData/model"
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(toDo *model.ToDo) (interface{}, error) {
	// Add the current timestamp
	toDo.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	// Insert to DB
	inserted, err := database.Collection.InsertOne(context.Background(), toDo)
	if err != nil {
		log.Fatalln("Insert:", err)
		return "", err
	}

	return inserted.InsertedID, nil
}

func GetOne(id string) (*model.ToDo, error) {
	toDo := model.ToDo{}

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln("Hex conversion:", err)
		return nil, errors.New("hex conversion")
	}

	err = database.Collection.FindOne(context.TODO(), bson.M{"_id": idHex}).Decode(&toDo)
	if err != nil {
		log.Fatalln("GetOne:", err)
		return nil, err
	}

	return &toDo, nil
}

func DeleteOne(id string) (int, error) {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, errors.New("hex conversion")
	}

	result, err := database.Collection.DeleteOne(context.TODO(), bson.M{"_id": idHex})
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func UpdateOne(id string, newToDo model.ToDo) (int, error) {

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln("Hex conversion:", err)
		return 0, errors.New("hex conversion")
	}

	// Get the bson.M that is required to send
	updateBson := model.PrepareBsonUpdateTodo(newToDo)

	result, err := database.Collection.UpdateOne(context.TODO(), bson.M{"_id": idHex}, updateBson)

	if err != nil {
		log.Fatalln("UpdateOne:", err)
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func DeleteMultiple(listToDelete []string) (int, error) {
	listToDeleteHex := []primitive.ObjectID{}
	for i := 0; i < len(listToDelete); i++ {
		idHex, err := primitive.ObjectIDFromHex(listToDelete[i])
		if err != nil {
			log.Fatalln("Hex conversion:", err)
			return 0, errors.New("hex conversion")
		}
		listToDeleteHex = append(listToDeleteHex, idHex)
	}

	result, err := database.Collection.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": listToDeleteHex}})
	if err != nil {
		log.Fatalln("Database delete many:", err)
		return 0, err
	}

	return int(result.DeletedCount), err
}

func GetAll(params string) ([]model.ToDo, error) {
	toDos := model.ToDos{}
	filter := bson.D{}
	if params != "" {
		b, _ := strconv.ParseBool(params)
		if b {
			filter = bson.D{{"status", b}}
		} else {
			filter = bson.D{{"status", bson.D{{"$ne", true}}}}
		}
	}
	findOptions := options.Find()
	// Sort by `deadline` field descending
	findOptions.SetSort(bson.D{{"deadline", 1}})
	cursor, err := database.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatalln("DB Read:", err)
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor failure:", err)
		return nil, err
	}

	if err = cursor.All(context.TODO(), &toDos.ToDos); err != nil {
		log.Fatalln("Cursor decode fail:", err)
		return nil, err
	}

	cursor.Close(context.TODO())

	return toDos.ToDos, nil
}
