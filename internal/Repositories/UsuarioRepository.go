package repositories

import (
	"context"

	models "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsuarioRepository interface {
	Inserir(ctx context.Context, usuario models.Usuario) (models.Usuario, error)
	PegarTodos(ctx context.Context) ([]models.Usuario, error)
	PegarUmPorId(ctx context.Context, id int) (models.Usuario, error)
	Atualizar(ctx context.Context, usuario models.Usuario) error
	Deletar(ctx context.Context, id int) error
}

type RepositoryDb struct {
	Connection *pgxpool.Pool
}

func (r *RepositoryDb) Inserir(ctx context.Context, usuario models.Usuario) (models.Usuario, error) {
	err := r.Connection.QueryRow(
		ctx,
		"INSERT INTO usuario (nome) VALUES ($1) RETURNING id, nome",
		usuario.Nome).Scan(&usuario.Id, &usuario.Nome)

	if err != nil {
		return models.Usuario{}, err
	}

	return usuario, nil
}

func (r *RepositoryDb) PegarTodos(ctx context.Context) ([]models.Usuario, error) {
	rows, err := r.Connection.Query(
		ctx,
		"SELECT * FROM usuario",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.Usuario

	for rows.Next() {
		var item models.Usuario
		if err := rows.Scan(item.Nome); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	erro := rows.Err()
	if erro != nil {
		return nil, erro
	}

	return items, nil
}

func (r *RepositoryDb) PegarUmPorId(ctx context.Context, id int) (models.Usuario, error) {
	var usuario = models.Usuario{Id: id}
	err := r.Connection.QueryRow(
		ctx,
		"SELECT id, nome FROM usuario  WHERE ID = $1",
		id).Scan(&usuario.Id, &usuario.Nome)

	if err != nil {
		return models.Usuario{}, err
	}

	return usuario, nil
}

func (r *RepositoryDb) Atualizar(ctx context.Context, usuario models.Usuario) error {
	_, err := r.Connection.Exec(
		ctx,
		"UPDATE usuario SET nome = $1 WHERE id = $2",
		usuario.Nome, usuario.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryDb) Deletar(ctx context.Context, id int) error {
	_, err := r.Connection.Exec(
		ctx,
		"DELETE FROM usuario WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
