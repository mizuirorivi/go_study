package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("###########################")
	fmt.Println("[*] Form Data ", r.Form)
	fmt.Println("[*] path", r.URL.Path)
	fmt.Println("[*] scheme", r.URL.Scheme)
	fmt.Println("[*] url: ", r.Form["url_long"])
	fmt.Println("[*] Form Data detailed")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(w, "key: %s, val: %s\n", k, strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello mizuiro_rivi!")
}

type login_info struct {
	username string
	password string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		fmt.Fprintf(w, "post_data:")
		fmt.Fprintf(w, body)
		fmt.Println("body:", body)
	}
}
func setCokkie(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "mizuiro_rivi", Expires: expiration}
	http.SetCookie(w, &cookie)
	t, _ := template.ParseFiles("setcokkie.gtpl")
	t.Execute(w, nil)
}

func readCokkie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("readCokkie")
	fmt.Fprintf(w, "Hey!! this is Cokkies!!!!! bunbun hello youtubeb!!!!")
	cookies := r.Cookies()
	for _, cookie := range cookies {
		fmt.Fprintf(w, "%s:%s\n", cookie.Name, cookie.Value)
	}
}
func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/setCookie", setCokkie)
	http.HandleFunc("/readCookie", readCokkie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
