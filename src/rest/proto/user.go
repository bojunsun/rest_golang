package proto

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"rest/io"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json: "email" bson: "email"`
	Password string        `json: "password" bson: "password"`
	Username string        `json: "username" bson: "username"`
}

func md5Encode(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func NewUser(email string, password string, username string) (user *User, err error) {
	user = &User{Email: email, Password: md5Encode(password), Username: username}
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
	if user.Password != md5Encode(password) {
		err = errors.New("Wrong password.")
		return
	}
	return
}
