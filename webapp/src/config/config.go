package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ApiUrl representa a URL para comunicação com a API
	ApiUrl = "127.0.0.1"

	// Porta onde a aplicação web vai estar rodando
	Porta = 0

	// HashKey é utilizada para autenticar o cookie
	HashKey []byte

	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
