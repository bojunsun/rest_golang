package main

import (
	"encoding/json"
	"net/http"
	"rest/proto"
)

func signupPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var signupData proto.User
	err := decoder.Decode(&signupData)
	if err != nil {
		panic(err)
	}

	var res interface{}

	if user, err := proto.NewUser(signupData.Email, signupData.Password, signupData.Username); err != nil {
		res = proto.NewError("NewUser err")
	} else {
		if err = user.SignUp(); err != nil {
			res = proto.NewError("User SignUp err")
		} else {
			res, _ = proto.NewResult("true", *user)
		}
	}
	var js []byte
	js, _ = json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
