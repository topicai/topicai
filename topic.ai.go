package main

import (
	"flag"
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	switch fmt.Sprint(r.URL) {
	case "/weakand", "/weakand/":
		http.Redirect(w, r, "https://github.com/wangkuiyi/weakand", http.StatusFound)
	case "/phoenix", "/phoenix/":
		http.Redirect(w, r, "https://github.com/wangkuiyi/phoenix", http.StatusFound)
	default:
		fmt.Fprintf(w, "Hello! Welcome to topic.ai!")
	}
}

func main() {
	addr := flag.String("addr", ":80", "Listening address")
	flag.Parse()

	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(*addr, nil)
}
