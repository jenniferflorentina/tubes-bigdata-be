package domain

import (
	"context"
	"time"

	"github.com/tubes-bigdata/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func Find(id string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		restErr := utils.NotFound("id wrong.")
		return nil, restErr
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = usersC.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

func GetAll() ([]*User, *utils.RestErr) {
	var users []*User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cur, err := usersC.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result *User
		err = cur.Decode(&result)
		if err != nil {
			restErr := utils.NotFound("failed to get all.")
			return users, restErr
		}
		result.Password = ""
		users = append(users, result)
	}
	return users, nil
}

func Delete(id string) *utils.RestErr {
	usersC := db.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		restErr := utils.NotFound("id wrong.")
		return restErr
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		restErr := utils.NotFound("failed to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return restErr
	}
	return nil
}

func Update(id string, user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		restErr := utils.NotFound("id wrong.")
		return nil, restErr
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"name": user.Name,
		"email": user.Email}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(id)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
