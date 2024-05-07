package models

import "errors"

type Usuario struct {
	Id   int
	Nome string
}

var (
	ErrNomeInvalido = errors.New("Dado passado é inválido")
)

func (p *Usuario) Validate() (Usuario, error) {
	if p.Nome == "" {
		return Usuario{}, ErrNomeInvalido
	}

	return Usuario{
		Nome: p.Nome,
	}, nil
}
