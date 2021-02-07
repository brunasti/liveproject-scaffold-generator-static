package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func main() {
	fmt.Println(appName + " ------------")

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	//server.ListenAndServe()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appName + " -------- end")

}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ [%v]\n", r.URL.Path[1:])
	files := []string{"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html"}

	templateFiles, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("/ [%v] - Error processing [%v]\n", r.URL.Path[1:], err)
		return
	}
	templates := template.Must(templateFiles, err)

	var threads = [1]Thread{}

	templates.ExecuteTemplate(w, "layout", threads)

}
