package main

import (
	"net/http"

	"github.com/ho-ryu-ji/session/router"
)

func main() {
	http.HandleFunc("/login", router.Login)
	http.HandleFunc("/register", router.Register)
	http.ListenAndServe(":8080", nil)
}
