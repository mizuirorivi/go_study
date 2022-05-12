package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
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
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		fmt.Println("[*] Get request")
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./login_after.gtpl")
		t.Execute(w, token)
	} else {
		fmt.Println("[*] Method: POST")
		if r.Form.Get("token") == "" {
			fmt.Fprintf(w, "token is empty")
		}
		// bufbody := new(bytes.Buffer)
		// bufbody.ReadFrom(r.Body)
		// body := bufbody.String()
		// fmt.Fprintf(w, "post_data:")
		// fmt.Fprintf(w, body)
		// fmt.Println("body:", body)
	}
}
func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
