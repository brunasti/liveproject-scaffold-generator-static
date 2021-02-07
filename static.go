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

	mux.HandleFunc("/healthcheck", healthCheck)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appName + " -------- end")

}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/healthcheck [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "ok")
	if err != nil {
		log.Printf("/healthcheck - Error processing [%v][%d]\n", err, n)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ [%v]\n", r.URL.Path[1:])
	files := []string{"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html"}

	templateFiles, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("/ [%v] - Error retrieving files [%v]\n", r.URL.Path[1:], err)
		return
	}
	templates := template.Must(templateFiles, err)

	// TODO Create content
	var threads = [1]Thread{}

	err = templates.ExecuteTemplate(w, "layout", threads)

	if err != nil {
		log.Printf("/ [%v] - Error executing template [%v]\n", r.URL.Path[1:], err)
		return
	}

}
