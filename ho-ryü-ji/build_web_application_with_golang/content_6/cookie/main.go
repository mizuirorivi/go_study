package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "ho-ryu-ji", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Setting Cookie...")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cookie set ↓")
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/readCookie", readCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
