package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Vermelha, bem linda", Preco: 39, Quantidade: 5},
		{"Vestido", "Branco", 79, 5},
		{"Fone", "Preto", 155, 6},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
