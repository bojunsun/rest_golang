package proto

import (
	"errors"
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

func GetUser(email string) (user *User, err error) {
	mc := io.NewMongoClient()
	defer mc.Close()
	user = &User{}
	if err = mc.GetOne("users", bson.M{"email": email}, user); err != nil {
		err = errors.New("User not found")
	}
	return
}

func (user *User) LogIn(password string) (err error) {
	if user.Password != password {
		err = errors.New("Wrong password.")
		return
	}
	return
}
