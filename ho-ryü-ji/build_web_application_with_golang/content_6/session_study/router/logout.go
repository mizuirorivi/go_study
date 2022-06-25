package router

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/ho-ryue-ji/session_study/db"
)

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionid, err := r.Cookie("sessionid")
	if err != nil {
		log.Fatal(err)
	}
	sessionid.MaxAge = -1
	http.SetCookie(w, sessionid)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout method:", r.Method)
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	case "POST":
		DeleteSession(w, r)
		// db.Delete(sessionid.Value)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
