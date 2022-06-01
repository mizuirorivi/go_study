package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ho-ryue-ji/session_study/db"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pages/login.gtpl")
	t.Execute(w, nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

		sessionid, _ := r.Cookie("sessionid")

		if sessionid != nil {
			_, err := db.Get(sessionid.Value)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Redirect(w, r, "/register", http.StatusFound)
					return
				}
				log.Fatal(err)
			}

			t, _ := template.ParseFiles("pages/success.gtpl")
			t.Execute(w, nil)
			return
		}
		ViewLogin(w, r)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
