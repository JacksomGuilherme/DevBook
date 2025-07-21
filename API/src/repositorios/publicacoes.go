package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO PUBLICACOES (TITULO, CONTEUDO, AUTOR_ID) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Publicacoes) BuscarPublicacoes() {

}

func (repositorio Publicacoes) BuscarPorId(ID uint64) (modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT ID, TITULO, CONTEUDO, AUTOR_ID, CURTIDAS, CRIADAEM FROM PUBLICACOES WHERE ID = ?",
		ID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Atualizar(publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE PUBLICACOES SET TITULO = ?, SET CONTEUDO = ?) WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Deletar() {

}
