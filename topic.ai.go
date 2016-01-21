package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	switch fmt.Sprint(r.URL) {
	case "/weakand":
		http.Redirect(w, r, "https://github.com/wangkuiyi/weakand", http.StatusFound)
	case "/phoenix":
		http.Redirect(w, r, "https://github.com/wangkuiyi/phoenix", http.StatusFound)
	default:
		fmt.Fprintf(w, "Hello! Welcome to topic.ai!")
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
