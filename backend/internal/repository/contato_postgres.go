package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/errors"
)

type ContatoPostgres struct {
	db *sql.DB
}

func NewContatoPostgres(db *sql.DB) *ContatoPostgres {
	return &ContatoPostgres{db: db}
}

func (r *ContatoPostgres) Create(ctx context.Context, contato *entity.Contato) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao iniciar transacao para criar contato")
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO Contato (ID, NOME, IDADE) VALUES ($1, $2, $3)",
		contato.ID, contato.Nome, contato.Idade)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao inserir contato")
	}

	for _, telefone := range contato.Telefones {
		_, err := tx.ExecContext(ctx, "INSERT INTO Telefone (IDCONTATO, ID, NUMERO) VALUES ($1, $2, $3)",
			telefone.IDContato, telefone.ID, telefone.Numero)
		if err != nil {
			return errors.WrapErrorf(err, "repositorio: falha ao inserir telefone para o contato %d", contato.ID)
		}
	}

	return tx.Commit()
}

func (r *ContatoPostgres) FindAll(ctx context.Context) ([]*entity.Contato, error) {
	return r.FindWithFilters(ctx, "", "")
}

func (r *ContatoPostgres) FindWithFilters(ctx context.Context, nome string, numero string) ([]*entity.Contato, error) {
	query := `
		SELECT c.ID, c.NOME, c.IDADE, t.IDCONTATO, t.ID, t.NUMERO 
		FROM Contato c 
		LEFT JOIN Telefone t ON c.ID = t.IDCONTATO 
		WHERE 1=1
	`
	var args []interface{}
	argCount := 1

	if nome != "" {
		query += fmt.Sprintf(" AND c.NOME ILIKE $%d", argCount)
		args = append(args, "%"+nome+"%")
		argCount++
	}

	if numero != "" {
		query += fmt.Sprintf(" AND EXISTS (SELECT 1 FROM Telefone t2 WHERE t2.IDCONTATO = c.ID AND t2.NUMERO LIKE $%d)", argCount)
		args = append(args, "%"+numero+"%")
		argCount++
	}

	query += " ORDER BY c.ID, t.ID"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.WrapErrorf(err, "repositorio: falha ao consultar contatos com filtros")
	}
	defer rows.Close()

	contactsMap := make(map[int64]*entity.Contato)
	var contacts []*entity.Contato

	for rows.Next() {
		var (
			contatoID         sql.NullInt64
			contatoNome       sql.NullString
			contatoIdade      sql.NullInt32
			telefoneIDContato sql.NullInt64
			telefoneID        sql.NullInt64
			telefoneNumero    sql.NullString
		)
		err := rows.Scan(&contatoID, &contatoNome, &contatoIdade, &telefoneIDContato, &telefoneID, &telefoneNumero)
		if err != nil {
			return nil, errors.WrapErrorf(err, "repositorio: falha ao escanear linha da consulta de contatos com filtros")
		}

		if _, ok := contactsMap[contatoID.Int64]; !ok {
			contato := &entity.Contato{
				ID:    contatoID.Int64,
				Nome:  contatoNome.String,
				Idade: int(contatoIdade.Int32),
			}
			contactsMap[contatoID.Int64] = contato
			contacts = append(contacts, contato)
		}

		if telefoneID.Valid {
			contactsMap[contatoID.Int64].Telefones = append(contactsMap[contatoID.Int64].Telefones, entity.Telefone{
				IDContato: telefoneIDContato.Int64,
				ID:        telefoneID.Int64,
				Numero:    telefoneNumero.String,
			})
		}
	}
	return contacts, nil
}

func (r *ContatoPostgres) FindByID(ctx context.Context, id int64) (*entity.Contato, error) {
	contato := &entity.Contato{}
	var (
		contatoID         sql.NullInt64
		contatoNome       sql.NullString
		contatoIdade      sql.NullInt32
		telefoneIDContato sql.NullInt64
		telefoneID        sql.NullInt64
		telefoneNumero    sql.NullString
	)

	rows, err := r.db.QueryContext(ctx, "SELECT c.ID, c.NOME, c.IDADE, t.IDCONTATO, t.ID, t.NUMERO FROM Contato c LEFT JOIN Telefone t ON c.ID = t.IDCONTATO WHERE c.ID = $1 ORDER BY t.ID", id)
	if err != nil {
		return nil, errors.WrapErrorf(err, "repositorio: falha ao consultar contato por ID %d", id)
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		err := rows.Scan(&contatoID, &contatoNome, &contatoIdade, &telefoneIDContato, &telefoneID, &telefoneNumero)
		if err != nil {
			return nil, errors.WrapErrorf(err, "repositorio: falha ao escanear linha da consulta de contato por ID %d", id)
		}

		if !found {
			contato.ID = contatoID.Int64
			contato.Nome = contatoNome.String
			contato.Idade = int(contatoIdade.Int32)
			found = true
		}

		if telefoneID.Valid {
			contato.Telefones = append(contato.Telefones, entity.Telefone{
				IDContato: telefoneIDContato.Int64,
				ID:        telefoneID.Int64,
				Numero:    telefoneNumero.String,
			})
		}
	}

	if !found {
		return nil, errors.ErrNotFound
	}

	return contato, nil
}

func (r *ContatoPostgres) Update(ctx context.Context, contato *entity.Contato) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao iniciar transacao para atualizar contato %d", contato.ID)
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, "UPDATE Contato SET NOME = $1, IDADE = $2 WHERE ID = $3",
		contato.Nome, contato.Idade, contato.ID)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao atualizar contato %d", contato.ID)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM Telefone WHERE IDCONTATO = $1", contato.ID)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao deletar telefones do contato %d antes da atualizacao", contato.ID)
	}

	for _, telefone := range contato.Telefones {
		_, err := tx.ExecContext(ctx, "INSERT INTO Telefone (IDCONTATO, ID, NUMERO) VALUES ($1, $2, $3)",
			contato.ID, telefone.ID, telefone.Numero)
		if err != nil {
			return errors.WrapErrorf(err, "repositorio: falha ao inserir telefone %d para o contato %d durante a atualizacao", telefone.ID, contato.ID)
		}
	}

	return tx.Commit()
}

func (r *ContatoPostgres) Delete(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao iniciar transacao para deletar contato %d", id)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "DELETE FROM Telefone WHERE IDCONTATO = $1", id)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao deletar telefones do contato %d", id)
	}

	res, err := tx.ExecContext(ctx, "DELETE FROM Contato WHERE ID = $1", id)
	if err != nil {
		return errors.WrapErrorf(err, "repositorio: falha ao deletar contato %d", id)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}

	return tx.Commit()
}
