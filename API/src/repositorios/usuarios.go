package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO USUARIOS (NOME, NICK, EMAIL, SENHA) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM USUARIOS WHERE NOME LIKE ? OR NICK LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorId traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorId(ID uint64) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM USUARIOS WHERE ID = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar atualiza os dados de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE USUARIOS SET NOME = ?, NICK = ?, EMAIL = ? WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM USUARIOS WHERE ID = ?",
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

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT ID, SENHA FROM USUARIOS WHERE EMAIL = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

// Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(usuarioID uint64, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO SEGUIDORES (USUARIO_ID, SEGUIDOR_ID) VALUES (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

// PararDeSeguir permite que um usuário pare de seguir outro
func (repositorio Usuarios) PararDeSeguir(usuarioID uint64, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM SEGUIDORES WHERE USUARIO_ID = ? AND SEGUIDOR_ID = ? ",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT U.ID, U.NOME, U.NICK, U.EMAIL, U.CRIADOEM
		FROM USUARIOS U
		INNER JOIN SEGUIDORES S ON U.ID = S.SEGUIDOR_ID
		WHERE S.USUARIO_ID = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT U.ID, U.NOME, U.NICK, U.EMAIL, U.CRIADOEM
		FROM USUARIOS U
		INNER JOIN SEGUIDORES S ON U.ID = S.USUARIO_ID
		WHERE S.SEGUIDOR_ID = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSenha traz a senha de um usuário pelo id
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT SENHA FROM USUARIOS WHERE ID = ?", usuarioID)
	if erro != nil {
		return "", erro
	}

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil

}

// Atualizar atualiza os dados de um usuário no banco de dados
func (repositorio Usuarios) AtualizarSenha(ID uint64, senhaNova string) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE USUARIOS SET SENHA = ? WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senhaNova, ID); erro != nil {
		return erro
	}

	return nil
}
