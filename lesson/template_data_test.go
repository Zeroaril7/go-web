package lesson

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))
	t.ExecuteTemplate(w, "data.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Zero",
		"UserInfo": map[string]interface{}{
			"Address": "Jalan in aja dulu ",
			"Phone":   "0872635171238",
		},
	})
}

func TestTemplateData(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

type UserInfo struct {
	Address string
	Phone   string
}

type PageInfo struct {
	Title    string
	Name     string
	UserInfo UserInfo
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))
	t.ExecuteTemplate(w, "data.gohtml", PageInfo{
		Title: "Template Data Struct",
		Name:  "Zero",
		UserInfo: UserInfo{
			Address: "Jalan in aja dulu ",
			Phone:   "08263618232155",
		},
	})
}

func TestTemplateStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
