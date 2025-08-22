package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin vai renderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "login", nil)
}

// CarregarPaginaDeCadastroDeUsuario vai renderizar a tela de cadastro de usuários
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro", nil)
}

// CarregarPaginaPrincipal vai renderizar a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaDeEdicaoDePublicacao carrega a pagina de edição de uma publicação
func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.ApiUrl, publicacaoID)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodDeErro(w, response)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: erro.Error()})
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao", publicacao)
}

// CarregarPaginaDeUsuarios carrega a pagina com os usuários que atendem o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.ApiUrl, nomeOuNick)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodDeErro(w, response)
		return
	}

	var usuarios []modelos.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios", usuarios)
}
