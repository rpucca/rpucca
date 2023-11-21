package db

import {
	"database/sql"  //biblioteca para execução de consultas SQL
	_ "github.com/lib/pq" //biblioteca postgree
}

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=P0stAdm host=localhost sslmode=disable"
	db, conecta := sql.Open("postgres", conexao) //retorna dados da conexão ou erro
	if conecta != nil {
		println("NOK - conectaComBancoDeDados")
		panic(conecta.Error())
	} else {
		println("OK -conectaComBancoDeDados")
		return db
	}
}