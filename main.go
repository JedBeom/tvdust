package main

import (
	"net/http"
	"os"
)

func Run() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", index())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func main() {
	loadConfig()

	if len(os.Args) > 1 {
		Run()
	} else {
		go Run()
		open()
	}
}
