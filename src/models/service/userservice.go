package service

import (
	"context"
	"ginRestFulApi/src/models/entity"

	"ginRestFulApi/src/models/db"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func UserCollection(c *mongo.Database) *mongo.Collection {
	collection = c.Collection("user")
	return collection
}

//Create is to register new user
func Create(user *(entity.User)) string {

	collection = UserCollection(db.GetConnection())

	collection.InsertOne(context.TODO(), user)
	// // conn.
	// defer conn.Session.Close()

	// doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	// err := doc.FindOne(bson.M{"email": user.Email}, doc)
	// if err == nil {
	// 	return errors.New("Already Exist")
	// }
	// userModel := mogo.NewDoc(user).(*(entity.User))
	// err = mogo.Save(userModel)
	// if vErr, ok := err.(*mogo.ValidationError); ok {
	// 	return vErr
	// }
	return "ok"
}

// // Delete a user from DB
// func Delete(email string) error {
// 	user, _ := userservice.FindByEmail(email)
// 	conn := db.GetConnection()
// 	defer conn.Session.Close()
// 	err := user.Remove()
// 	return err
// }

// //Find user
// func Find(user *(entity.User)) (*entity.User, error) {
// 	conn := db.GetConnection()
// 	defer conn.Session.Close()

// 	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
// 	err := doc.FindOne(bson.M{"email": user.Email}, doc)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return doc, nil
// }

// //Find user from email
// func FindByEmail(email string) (*entity.User, error) {
// 	conn := db.GetConnection()
// 	defer conn.Session.Close()

// 	user := new(entity.User)
// 	user.Email = email
// 	return userservice.Find(user)
// }
