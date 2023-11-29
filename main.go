package main

import (
	"fmt"           //biblioteca para impressões
	"html/template" //biblioteca para processar e renderizar templates HTML
	"net/http"      //biblioteca funcionalidades para criar servidores HTTP

	"github.com\rpucca\rpucca\models"
)

var templ = template.Must(template.ParseGlob("templates/*.html")) //encapsula todos os templates (*.html) renderizando e retornando o template e msg de erro se houver.

func main() {
	http.HandleFunc("/", index) //acessa a raiz ("/") do servidor, e executa a função index
	fmt.Println(templ)
	http.ListenAndServe(":8000", nil) //sobe o servidor porta 8080
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	templ.ExecuteTemplate(w, "Index", todosOsProdutos)
}
