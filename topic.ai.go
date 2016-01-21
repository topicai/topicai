package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if fmt.Sprint(r.URL) == "/weakand" {
		http.Redirect(w, r, "https://github.com/wangkuiyi/weakand", http.StatusFound)
	} else {
		fmt.Println("Hello")
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
