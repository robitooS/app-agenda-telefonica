package repository

import (
	"context"

	"github.com/robitooS/backend/internal/entity"
)

type ContatoRepository interface {
	Create(ctx context.Context ,contato *entity.Contato) error
	FindAll(ctx context.Context) ([]*entity.Contato, error)
	FindWithFilters(ctx context.Context, nome string, numero string) ([]*entity.Contato, error)
	FindByID(ctx context.Context, id int64) (*entity.Contato, error)
	Update(ctx context.Context, contato *entity.Contato) error
	Delete(ctx context.Context, id int64) error
}
