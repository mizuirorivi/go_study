package router

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Success(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success method:", r.Method)
	switch r.Method {
	case "GET":
		pageid, err := LoginBySessionid(w, r)
		switch pageid {
		case "login":
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		case "success":
			t, _ := template.ParseFiles("pages/success.gtpl")
			t.Execute(w, nil)
		default:
			log.Fatal(err)
		}
	}
}
