package main

import (
	"database/sql"  //biblioteca para execução de consultas SQL
	"fmt"           //biblioteca para impressões
	"html/template" //biblioteca para processar e renderizar templates HTML
	"net/http"      //biblioteca funcionalidades para criar servidores HTTP

	_ "github.com/lib/pq" //biblioteca postgree
)

var templ = template.Must(template.ParseGlob("templates/*.html")) //encapsula todos os templates (*.html) renderizando e retornando o template e msg de erro se houver.

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=P0stAdm host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao) //retorna dados da conexão ou erro
	if err != nil {
		println("NOK - conectaComBancoDeDados")
		panic(err.Error())
	} else {
		println("OK -conectaComBancoDeDados")
		return db
	}
}

func main() {
	http.HandleFunc("/", index) //acessa a raiz ("/") do servidor, e executa a função index
	fmt.Println(templ)
	http.ListenAndServe(":8000", nil) //sobe o servidor porta 8080
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	/*
		lst_produtos := []str_Produto{
			{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		}
	*/

	templ.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
