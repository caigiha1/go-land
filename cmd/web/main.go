package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "http service address")

	flag.Parse()

	log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
	log.New(os.Stderr, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	log.Printf("Starting server at port 4000")
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
