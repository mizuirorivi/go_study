package main

import (
	"net/http"

	"github.com/ho-ryue-ji/session_study/router"
)

func main() {
	http.HandleFunc("/login", router.Login)
	http.HandleFunc("/register", router.Register)
	http.HandleFunc("/success", router.Success)
	http.HandleFunc("/logout", router.Logout)
	http.ListenAndServe(":8080", nil)
}
