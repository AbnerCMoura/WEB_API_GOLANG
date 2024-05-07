package services

import (
	"context"
	"errors"

	models "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Models"
	repositories "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Repositories"
)

var ErrNomeVazio = errors.New("Nome não pode ser vazio")

type UsuarioService struct {
	UsuarioRepository repositories.UsuarioRepository
}

func (s UsuarioService) Inserir(ctx context.Context, usuario models.Usuario) (models.Usuario, error) {
	if usuario.Nome == "" {
		return models.Usuario{}, ErrNomeVazio
	}

	return s.UsuarioRepository.Inserir(ctx, usuario)
}

func (s UsuarioService) Delete(ctx context.Context, id int) error {
	return s.UsuarioRepository.Deletar(ctx, id)
}

func (s UsuarioService) FindOneByID(ctx context.Context, id int) (models.Usuario, error) {
	return s.UsuarioRepository.PegarUmPorId(ctx, id)
}

func (s UsuarioService) FindAll(ctx context.Context) ([]models.Usuario, error) {
	usuarios, err := s.UsuarioRepository.PegarTodos(ctx)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (s UsuarioService) Update(ctx context.Context, params *models.Usuario) error {
	post, err := params.Validate()
	if err != nil {
		return err
	}

	err = s.UsuarioRepository.Atualizar(ctx, post)

	if err != nil {
		return err
	}

	return nil
}
