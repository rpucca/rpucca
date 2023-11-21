package main

import {
	"fmt"           //biblioteca para impressões
	"html/template" //biblioteca para processar e renderizar templates HTML
	"net/http"      //biblioteca funcionalidades para criar servidores HTTP
}

var templ = template.Must(template.ParseGlob("templates/*.html")) //encapsula todos os templates (*.html) renderizando e retornando o template e msg de erro se houver.

func main() {
	http.HandleFunc("/", index) //acessa a raiz ("/") do servidor, e executa a função index
	fmt.Println(templ)
	http.ListenAndServe(":8000", nil) //sobe o servidor porta 8080
}

func index(w http.ResponseWriter, r *http.Request) {
	var int_id, int_quantidade int
	var str_nome, str_descricao string
	var flo_preco float64
	linha_produto := strProduto{}   //instancia da estrutura produtos
	array_produto := []strProduto{} //array de produtos
	db := conectaComBancoDeDados()

	fmt.Println("Início")
	selectDeTodosOsProdutos, registro := db.Query("select * from produtos")
	if registro != nil {
		panic(registro.Error())
	}

	for selectDeTodosOsProdutos.Next() {
		defer fmt.Println("Meio")

		registro = selectDeTodosOsProdutos.Scan(&int_id, &str_nome, &str_descricao, &flo_preco, &int_quantidade)
		if registro != nil {
			panic(registro.Error())
		}

		linha_produto.Nome = str_nome
		linha_produto.Descricao = str_descricao
		linha_produto.Preco = flo_preco
		linha_produto.Quantidade = int_quantidade

		array_produto = append(array_produto, linha_produto)
	}
	/*
		lst_produtos := []str_Produto{
			{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		}
	*/

	templ.ExecuteTemplate(w, "Index", array_produto)
	defer db.Close()  //com defer o close é adiada até que a função exemplo() seja concluída.
	fmt.Println("Fim")

}


