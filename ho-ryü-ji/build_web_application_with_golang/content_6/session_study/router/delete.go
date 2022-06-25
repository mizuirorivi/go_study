package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ho-ryue-ji/session_study/db"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete method:", r.Method)
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	case "POST":
		sessionid, err := r.Cookie("sessionid")
		if err != nil {
			log.Fatal(err)
		}
		user, err := db.Get("SessionId", sessionid.Value)
		if err != nil {
			log.Fatal(err)
		}

		user.Life = 0
		db.Update(sessionid.Value, user)
		DeleteSession(w, r)
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}
