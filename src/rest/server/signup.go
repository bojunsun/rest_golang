package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/proto"
)

type cred struct {
	Email    string
	Password string
}

func signupPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var signupData cred
	err := decoder.Decode(&signupData)
	if err != nil {
		panic(err)
	}

	fmt.Println(signupData)

	if user, err := proto.NewUser(signupData.Email, signupData.Password); err != nil {
	} else {

		if err = user.SignUp(); err != nil {

		} else {
			var js []byte
			js, _ = json.Marshal(signupData)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}

}
