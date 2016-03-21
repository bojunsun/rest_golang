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
