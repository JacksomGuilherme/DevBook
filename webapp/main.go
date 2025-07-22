package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	r := router.Gerar()

	utils.CarregarTemplates()

	fmt.Printf("Listening at: %s:%d\n", "127.0.0.1", 3030)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", 3030), r))
}
