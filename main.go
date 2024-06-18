package main

import (
	"fmt"
	"net/http"
)

func main() {
	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Server Start")
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server Start")
	})

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "images")
	})

	mux.HandleFunc("/images/thumbnails", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "thumbnails")
	})

	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "user data")
	})

	mux.HandleFunc("/user/image", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "user image")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux, //Bisa menggunakan handler yang di comment
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
