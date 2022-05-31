package router

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ho-ryue-ji/session_study/db"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		sessionid, err := r.Cookie("sessionid")
		if err != nil {
			log.Fatal(err)
		}
		user, err := db.Get(sessionid.Value)
		if err != nil {
			log.Fatal(err)
		}
		if user.SessionId == sessionid.Value {
			t, _ := template.ParseFiles("pages/success.gtpl")
			t.Execute(w, nil)
		} else {
			t, _ := template.ParseFiles("pages/login.gtpl")
			t.Execute(w, nil)
		}

	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
