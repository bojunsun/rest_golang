package main

import (
	"github.com/gorilla/mux"
	"net/http"
	//"rest/proto"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", signupPost).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
