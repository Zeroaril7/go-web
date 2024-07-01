package lesson

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	if err := myTemplates.ExecuteTemplate(w, "upload.html", nil); err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload_success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("resources"))))

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/photo_6055474313777234042_y.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	w := multipart.NewWriter(body)
	w.WriteField("name", "Test")
	file, _ := w.CreateFormFile("file", "test.png")
	file.Write(uploadFileTest)
	w.Close()

	req := httptest.NewRequest(http.MethodGet, mainURL+pathUpload, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)

	bodyRes, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(bodyRes))
}
