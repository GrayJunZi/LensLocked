package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at</p><a href='mailto:test@123.com'>test@123.com</a>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", nil)
}
