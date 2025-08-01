package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type Erro struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// TratarStatusCodDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodDeErro(w http.ResponseWriter, r *http.Response) {
	var erro Erro
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
