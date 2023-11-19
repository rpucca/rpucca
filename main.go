package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq" //biblioteca postgree
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=P0stAdm host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		println("NOK")
		panic(err.Error())
	} else {
		println("OK")
		return db
	}
	println("Fim")
}

var templ = template.Must(template.ParseGlob("templates/*.html")) //encapsula todos os templates (*.html) renderizando e retornando o template e msg de erro se houver.

type str_Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
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
