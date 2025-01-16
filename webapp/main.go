package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/router/rotas"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Println("rodando o app")

	rotas.Configurar(r)
	fmt.Println("gerou a porta")
	log.Fatal(http.ListenAndServe(":3000", r))
}
