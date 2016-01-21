package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	goTagsTemplate = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.ImportPath}} git {{.RepoPath}}">
</head>
<body>
</body>
</html>
`
)

var (
	temp *template.Template
)

type GoTagMap struct {
	ImportPath string
	RepoPath   string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	switch fmt.Sprint(r.URL) {
	case "/weakand", "/weakand/":
		temp.Execute(w, GoTagMap{
			ImportPath: "topic.ai/weakand",
			RepoPath:   "https://github.com/wangkuiyi/weakand"})
	case "/phoenix", "/phoenix/":
		temp.Execute(w, GoTagMap{
			ImportPath: "topic.ai/phoenix",
			RepoPath:   "https://github.com/wangkuiyi/phoenix"})
	default:
		fmt.Fprintf(w, "Hello! Welcome to topic.ai!")
	}
}

func main() {
	addr := flag.String("addr", ":80", "Listening address")
	flag.Parse()

	if t, e := template.New("gotags").Parse(goTagsTemplate); e != nil {
		log.Fatal(e)
	} else {
		temp = t
	}

	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(*addr, nil)
}
