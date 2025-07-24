package main

import (
	"fmt"
	"log"
	"net/http"
	cookies "webapp/src/Cookies"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	r := router.Gerar()

	utils.CarregarTemplates()

	fmt.Printf("Listening at: %s:%d\n", "http://127.0.0.1", config.Porta)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
