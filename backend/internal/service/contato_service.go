package service

import (
	"context"
	"github.com/robitooS/backend/internal/entity"
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
	return s.repo.Create(ctx, contato)
}

func (s *contatoService) FindAll(ctx context.Context) ([]*entity.Contato, error) {
	return s.repo.FindAll(ctx)
}

func (s *contatoService) FindWithFilters(ctx context.Context, nome string, numero string) ([]*entity.Contato, error) {
	return s.repo.FindWithFilters(ctx, nome, numero)
}

func (s *contatoService) FindByID(ctx context.Context, id int64) (*entity.Contato, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *contatoService) Update(ctx context.Context, contato *entity.Contato) error {
	return s.repo.Update(ctx, contato)
}

func (s *contatoService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
