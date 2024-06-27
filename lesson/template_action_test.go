package lesson

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Address struct {
	Street string
	City   string
}

type User struct {
	Name    string
	Score   int
	Hobbies []string
	Address Address
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action/if.html"))

	t.ExecuteTemplate(w, "if.html", User{
		Name: "Budi",
	})
}

func TestActionIf(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateActionIf(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action/operator.html"))

	t.ExecuteTemplate(w, "operator.html", User{
		Score: 80,
	})
}

func TestActionOperator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateActionOperator(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action/range.html"))

	t.ExecuteTemplate(w, "range.html", User{
		Hobbies: []string{
			"Baca Buku",
			"Main Game",
		},
	})
}

func TestActionRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateActionRange(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action/with.html"))

	t.ExecuteTemplate(w, "with.html", User{
		Name: "Budi",
		Address: Address{
			Street: "Jalan in aja dulu",
			City:   "Surabaya",
		},
	})
}

func TestActionWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateActionWith(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
