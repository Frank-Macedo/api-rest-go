package main

import (
	"github.com/Frank-Macedo/api-rest-go/database"
	"github.com/Frank-Macedo/api-rest-go/models"
	"github.com/Frank-Macedo/api-rest-go/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
