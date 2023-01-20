package main

import (
	"fmt"

	"github.com/BruggerPedro/api-go-rest/models"
	"github.com/BruggerPedro/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
