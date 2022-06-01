package router

import (
	"fmt"
	"net/http"
	"text/template"
)

func Success(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success: method:", r.Method)
	t, _ := template.ParseFiles("pages/success.gtpl")
	t.Execute(w, nil)
}
