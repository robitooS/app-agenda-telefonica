package service

import "github.com/robitooS/backend/internal/entity"

type ContatoService interface {
	Create(contato *entity.Contato) error
	FindAll() ([]*entity.Contato, error)
	FindByID(id int64) (*entity.Contato, error)
	Update(contato *entity.Contato) error
	Delete(id int64) error
}
