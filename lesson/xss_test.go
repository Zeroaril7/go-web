package lesson

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "xss.html", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Hai Template</p><script>alert('Hei')</script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL, nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscape(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TemplateDisableAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "xss.html", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<h1>Hai Template</h1>"),
	})
}

func TestTemplateDisableAutoEscape(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL, nil)
	rec := httptest.NewRecorder()

	TemplateDisableAutoEscape(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestTemplateDisableAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(TemplateDisableAutoEscape),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "xss.html", map[string]interface{}{
		"Title": "XSS",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestXSS(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+"/?body=<p>alert</p>", nil)
	rec := httptest.NewRecorder()

	TemplateXSS(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestTemplateXSS(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
