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

// Criar permite incluir uma nova publicação no banco de dados
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

// BuscarPublicacoes traz as publicações dos usuários seguidos e do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT P.*, U.nick FROM publicacoes P 
		INNER JOIN usuarios U ON U.id = P.autor_id 
		INNER JOIN seguidores S ON P.autor_id = S.usuario_id 
		WHERE U.id = ? OR S.seguidor_id = ?
		ORDER BY P.ID DESC`,
		usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

// BuscarPorId traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorId(ID uint64) (modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`SELECT P.*, U.NICK FROM PUBLICACOES P 
		 INNER JOIN USUARIOS U ON U.ID = P.AUTOR_ID 
		 WHERE P.ID = ? `,
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
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Atualizar permite alterar uma publicação do banco de dados
func (repositorio Publicacoes) Atualizar(ID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE PUBLICACOES SET TITULO = ?, CONTEUDO = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, ID)
	if erro != nil {
		return erro
	}

	return nil
}

// Deletar permite remover uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM PUBLICACOES WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz todas as publicações de um usuário específico do banco de dados
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`SELECT P.*, U.NICK FROM PUBLICACOES P 
		 INNER JOIN USUARIOS U ON U.ID = P.AUTOR_ID 
		 WHERE P.AUTOR_ID = ? `,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida da publicação no banco de dados
func (repositorio Publicacoes) Curtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE PUBLICACOES SET CURTIDAS = CURTIDAS + 1 WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida da publicação no banco de dados
func (repositorio Publicacoes) Descurtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE PUBLICACOES SET CURTIDAS = 
			CASE WHEN CURTIDAS > 0 THEN CURTIDAS - 1 ELSE 0 END 
		WHERE ID = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
