package models

import (
	"fmt"

	"github.com\rpucca\rpucca\db"
)

type structProduto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []structProduto {
	var int_id, int_quantidade int
	var str_nome, str_descricao string
	var flo_preco float64
	linha_produto := structProduto{}   //instancia da estrutura produtos
	array_produto := []structProduto{} //array de produtos
	db := db.ConectaComBancoDeDados()

	fmt.Println("Início BuscaTodosOsProdutos")
	selectDeTodosOsProdutos, registro := db.Query("select * from produtos")
	if registro != nil {
		panic(registro.Error())
	}

	for selectDeTodosOsProdutos.Next() {
		defer fmt.Println("Meio BuscaTodosOsProdutos")

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

	fmt.Println("Fim BuscaTodosOsProdutos")
	defer db.Close() //com defer o close é adiada até que a função exemplo() seja concluída.
	return array_produto
}
