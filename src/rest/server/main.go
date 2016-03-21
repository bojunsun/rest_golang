package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("rest"))

func generalSession() (name string) {
	name = "rest_common"
	return
}

func checkLogin(r *http.Request) bool {
	session, _ := store.Get(r, generalSession())
	if session.Values["login"] == nil || session.Values["login"].(bool) == false {
		return false
	} else {
		return true
	}

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", SignupPost).Methods("POST")
	router.HandleFunc("/login", LoginPost).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
