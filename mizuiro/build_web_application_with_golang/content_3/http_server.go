package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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

// func (srv *http.Server) myServe(l net.Listener) error {
// 	defer l.Close()
// 	var tempDelay time.Duration // how long to sleep on accept  failure
// 	for {
// 		rw, e := l.Accept()
// 		if e != nil {
// 			if ne, ok := e.(net.Error); ok && ne.Temporary() {
// 				if tempDelay == 0 {
// 					tempDelay = 5 * time.Millisecond
// 				} else {
// 					tempDelay *= 2
// 				}
// 				if max := 1 * time.Second; tempDelay > max {
// 					tempDelay = max
// 				}
// 				log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
// 				time.Sleep(tempDelay)
// 				continue
// 			}
// 			return e
// 		}
// 		tempDelay = 0
// 		c, err := srv.newConn(rw)
// 		if err != nil {
// 			continue
// 		}
// 		go c.serve()
// 	}
// }
func main() {
	http.HandleFunc("/", sayhelloName)
	// server := &http.Server{Addr: ":9090", Handler: nil}
	// ln, err := net.Listen("tcp", server.Addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = server.myServe(ln)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
