package lesson

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	typeFile := r.URL.Query().Get("type")

	switch typeFile {
	case "html":
		http.ServeFile(w, r, "./src/index.html")
		break
	case "js":
		http.ServeFile(w, r, "./src/index.js")
	default:
		fmt.Fprint(w, "File Not Found")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFile),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed src/index.html
var html string

//go:embed src/index.js
var js string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	typeFile := r.URL.Query().Get("type")

	switch typeFile {
	case "html":
		fmt.Fprint(w, html)
		break
	case "js":
		fmt.Fprint(w, js)
	default:
		fmt.Fprint(w, "File Not Found")
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
