package proto

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json: "email"`
	Password string        `json: "password"`
	Username string        `json: "username"`
}

// func signupPost(w http.ResponseWriter, r *http.Request) {
// 	log.Println("signup")
// }
