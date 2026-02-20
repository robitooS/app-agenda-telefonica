package service

import (
	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/repository"
	"log"
	"os"
	"time"
	"fmt"
)

type contatoService struct {
	repo    repository.ContatoRepository
	logPath string
}

func NewContatoService(repo repository.ContatoRepository, logPath string) ContatoService {
	return &contatoService{repo: repo, logPath: logPath}
}

func (s *contatoService) Create(contato *entity.Contato) error {
	return s.repo.Create(contato)
}

func (s *contatoService) FindAll() ([]*entity.Contato, error) {
	return s.repo.FindAll()
}

func (s *contatoService) FindByID(id int64) (*entity.Contato, error) {
	return s.repo.FindByID(id)
}

func (s *contatoService) Update(contato *entity.Contato) error {
	return s.repo.Update(contato)
}

func (s *contatoService) Delete(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	file, err := os.OpenFile(s.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Falha ao abrir o arquivo de log %s: %v", s.logPath, err)
		return nil // Não bloqueia a exclusão se o log falhar
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%s - Contato ID %d excluído.\n", time.Now().Format("2006-01-02 15:04:05"), id)
	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Falha ao escrever no arquivo de log %s: %v", s.logPath, err)
	}

	return nil
}
