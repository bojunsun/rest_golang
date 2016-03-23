package main

import (
	"encoding/json"
	//"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"rest/proto"
)

func SignupPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var signupData proto.User
	err := decoder.Decode(&signupData)
	if err != nil {
		panic(err)
	}

	var res interface{}

	if _, err := proto.GetUser(signupData.Email); err == nil {
		res = proto.NewError("User Already Exist")
	} else {
		if user, err := proto.NewUser(signupData.Email, signupData.Password, signupData.Username); err != nil {
			res = proto.NewError("NewUser err")
		} else {
			if err = user.SignUp(); err != nil {
				res = proto.NewError("User SignUp err")
			} else {
				res, _ = proto.NewResult("true", bson.M{"_id": user.ID})
			}
		}
	}
	var js []byte
	js, _ = json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginData proto.User
	err := decoder.Decode(&loginData)
	if err != nil {
		panic(err)
	}

	var res interface{}

	if user, err := proto.GetUser(loginData.Email); err != nil {
		res = proto.NewError("GetUser err")
	} else if user.LogIn(loginData.Password); err != nil {
		res = proto.NewError("User Login err")
	} else {
		session, _ := store.Get(r, generalSession())
		//store login(true/flase) and email in session
		session.Values["login"] = true
		session.Values["email"] = user.Email
		session.Save(r, w)
		res, _ = proto.NewResult("true", bson.M{"_id": user.ID})
	}

	var js []byte
	js, _ = json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	if checkLogin(r) {
		session, _ := store.Get(r, generalSession())
		session.Values["login"] = false
		session.Save(r, w)
		res, _ = proto.NewResult("true", nil)
	} else {
		res = proto.NewError("User has not login")
	}
	convertJson(res, w)
}

func UserselfGet(w http.ResponseWriter, r *http.Request) {

	var res interface{}

	if checkLogin(r) {
		session, _ := store.Get(r, generalSession())
		email := session.Values["email"].(string)

		if user, err := proto.GetUser(email); err != nil {
			res = proto.NewError("GetUSer err")
		} else {
			res, _ = proto.NewResult("true", *user)
		}
	} else {
		res = proto.NewError("User has not login")
	}

	var js []byte
	js, _ = json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
