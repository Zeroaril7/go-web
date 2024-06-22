package lesson

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed src
var rsc embed.FS

func TestFileServerEmbed(t *testing.T) {

	src, _ := fs.Sub(rsc, "src")
	fileServer := http.FileServer(http.FS(src))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

func TestFileServer(t *testing.T) {
	dir := http.Dir("../src")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
