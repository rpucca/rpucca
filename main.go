package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/*.html")) //encapsula todos os templates (*.html) renderizando e retornando o template e msg de erro se houver.

type str_Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index) //acessa a raiz ("/") do servidor, e executa a função index
	fmt.Println(templ)
	http.ListenAndServe(":8000", nil) //sobe o servidor porta 8080
}

func index(w http.ResponseWriter, r *http.Request) {
	lst_produtos := []str_Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	templ.ExecuteTemplate(w, "Index", lst_produtos)
}
