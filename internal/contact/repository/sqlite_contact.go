package repository

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/niumandzi/nto2022/internal/domain"
	"github.com/niumandzi/nto2022/model"
)

type sqliteContactRepository struct {
	Conn *sql.DB
}

func NewSqliteContactRepository(conn *sql.DB) domain.ContactRepository {
	return &sqliteContactRepository{conn}
}

func (s sqliteContactRepository) Create(ctx context.Context, contact model.Contact) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteContactRepository) Get(ctx context.Context, contactId int) (model.Contact, error) {
	row := s.Conn.QueryRowContext(ctx, "SELECT id, contact_type, name, number, email FROM contact WHERE id = ?", contactId)

	var contact model.Contact

	err := row.Scan(
		&contact.Id,
		&contact.ContactType,
		&contact.Name,
		&contact.Number,
		&contact.Email,
	)

	// Проверяем ошибки
	if err != nil {
		return model.Contact{}, err
	}

	return contact, nil
}

func (s sqliteContactRepository) GetByType(ctx context.Context, contactType string) (model.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (s sqliteContactRepository) GetAll(ctx context.Context) ([]model.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (s sqliteContactRepository) Update(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	//TODO implement me
	panic("implement me")
}

func (s sqliteContactRepository) Delete(ctx context.Context, contactId int) error {
	//TODO implement me
	panic("implement me")
}
