package main

import (
	"log"
	"net/http"
)

func signupPost(w http.ResponseWriter, r *http.Request) {
	log.Println("signup")
	log.Println("whatever")
}
