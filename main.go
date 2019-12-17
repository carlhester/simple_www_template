package main

import "fmt"
import "html/template"
import "net/http"
import "log"

type PageData struct {
	PageTitle string
}

func serveUI(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Panic(err)
	}
	data := fmt.Sprintf("%+v", req)
	page := PageData{PageTitle: data}
	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	serveAddress := ":8080"
	rMux := http.NewServeMux()
	rMux.HandleFunc("/", serveUI)
	log.Println(http.ListenAndServe(serveAddress, rMux))
}
