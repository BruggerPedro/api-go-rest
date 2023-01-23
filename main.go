package main

import (
	"fmt"

	"github.com/BruggerPedro/api-go-rest/database"
	"github.com/BruggerPedro/api-go-rest/models"
	"github.com/BruggerPedro/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConnectDB()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
