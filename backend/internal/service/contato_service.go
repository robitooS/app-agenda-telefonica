package service

import (
	"context"
	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/errors"
	"github.com/robitooS/backend/internal/repository"
)

type contatoService struct {
	repo    repository.ContatoRepository
}

func NewContatoService(repo repository.ContatoRepository) ContatoService {
	return &contatoService{
		repo: repo, 
	}
}

func (s *contatoService) Create(ctx context.Context, contato *entity.Contato) error {
	if err := s.repo.Create(ctx, contato); err != nil {
		return errors.WrapErrorf(err, "servico: falha ao criar contato")
	}
	return nil
}

func (s *contatoService) FindAll(ctx context.Context) ([]*entity.Contato, error) {
	contatos, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.WrapErrorf(err, "servico: falha ao buscar todos os contatos")
	}
	return contatos, nil
}

func (s *contatoService) FindWithFilters(ctx context.Context, nome string, numero string) ([]*entity.Contato, error) {
	contatos, err := s.repo.FindWithFilters(ctx, nome, numero)
	if err != nil {
		return nil, errors.WrapErrorf(err, "servico: falha ao buscar contatos com filtros")
	}
	return contatos, nil
}

func (s *contatoService) FindByID(ctx context.Context, id int64) (*entity.Contato, error) {
	contato, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.WrapErrorf(err, "servico: falha ao buscar contato por ID %d", id)
	}
	return contato, nil
}

func (s *contatoService) Update(ctx context.Context, contato *entity.Contato) error {
	if err := s.repo.Update(ctx, contato); err != nil {
		return errors.WrapErrorf(err, "servico: falha ao atualizar contato %d", contato.ID)
	}
	return nil
}

func (s *contatoService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.WrapErrorf(err, "servico: falha ao deletar contato %d", id)
	}
	return nil
}
