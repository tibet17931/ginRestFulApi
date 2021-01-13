package service

import (
	"errors"
	"ginRestFulApi/src/models/db"
	"ginRestFulApi/src/models/entity"

	"github.com/goonode/mogo"
	"labix.org/v2/mgo/bson"
)

//Create is to register new user
func Create(user *(entity.User)) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err == nil {
		return errors.New("Already Exist")
	}
	userModel := mogo.NewDoc(user).(*(entity.User))
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
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
