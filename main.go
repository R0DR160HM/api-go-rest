package main

import (
	"fmt"

	"github.com/R0DR160HM/api-go-rest/database"
	"github.com/R0DR160HM/api-go-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
