package lesson

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Redirect...")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://music.youtube.com/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
