package proto

import (
	"gopkg.in/mgo.v2/bson"
	"rest/io"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json: "email"`
	Password string        `json: "password"`
	Username string        `json: "username"`
}

func NewUser(email string, password string) (user *User, err error) {
	user = &User{Email: email, Password: password}
	return
}

func (user *User) SignUp() (err error) {
	user.ID = bson.NewObjectId()
	mc := io.NewMongoClient()
	defer mc.Close()
	mc.Insert("users", user)
	return
}
