package http_server

import (
	"fmt"
	"html/template"
	"net/http"
)

const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>System Design</title>
</head>
<body>
    <h1>{{.Title}}</h1>
</body>
</html>
`

type FrontendData struct {
	Title string
}

func handleRootPath(w http.ResponseWriter, _ *http.Request) {
	data := FrontendData{
		Title: "System Dashboard",
	}
	template, err := template.New("webpage").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StartHttpServer() {
	fmt.Println("hello world")
	http.HandleFunc("/", handleRootPath)
	http.ListenAndServe(":8080", nil)
}
