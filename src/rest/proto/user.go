package proto

import (
	"gopkg.in/mgo.v2/bson"
	"rest/io"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json: "email" bson: "email"`
	Password string        `json: "password" bson: "password"`
	Username string        `json: "username" bson: "username"`
}

func NewUser(email string, password string, username string) (user *User, err error) {
	user = &User{Email: email, Password: password, Username: username}
	return
}

func (user *User) SignUp() (err error) {
	user.ID = bson.NewObjectId()
	mc := io.NewMongoClient()
	defer mc.Close()
	mc.Insert("users", user)
	return
}
