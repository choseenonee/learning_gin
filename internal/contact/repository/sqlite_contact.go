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
	result, err := s.Conn.ExecContext(
		ctx,
		"INSERT INTO contact (ContactType, Name, Number, Email) VALUES (?, ?, ?, ?)",
		contact.ContactType,
		contact.Name,
		contact.Number,
		contact.Email,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *sqliteContactRepository) Get(ctx context.Context, contactId int) (model.Contact, error) {
	row := s.Conn.QueryRowContext(ctx, "SELECT * FROM contact WHERE id = ?", contactId)

	var contact model.Contact

	err := row.Scan(
		&contact.Id,
		&contact.ContactType,
		&contact.Name,
		&contact.Number,
		&contact.Email,
	)

	if err != nil {
		return model.Contact{}, err
	}

	return contact, nil
}

func (s sqliteContactRepository) GetByType(ctx context.Context, contactType string) ([]model.Contact, error) {
	rows, err := s.Conn.QueryContext(ctx, "SELECT * FROM contact WHERE contact_type = ?", contactType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact
	for rows.Next() {
		var contact model.Contact
		err := rows.Scan(&contact.Id, &contact.ContactType, &contact.Name, &contact.Number, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s sqliteContactRepository) GetAll(ctx context.Context) ([]model.Contact, error) {
	rows, err := s.Conn.QueryContext(ctx, "SELECT * FROM contact")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact
	for rows.Next() {
		var contact model.Contact
		err := rows.Scan(&contact.Id, &contact.ContactType, &contact.Name, &contact.Number, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s sqliteContactRepository) Update(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	_, err := s.Conn.ExecContext(
		ctx,
		"UPDATE contact SET ContactType=?, Name=?, Number=?, Email=? WHERE Id=?",
		contactInput.ContactType,
		contactInput.Name,
		contactInput.Number,
		contactInput.Email,
		contactId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s sqliteContactRepository) Delete(ctx context.Context, contactId int) error {
	_, err := s.Conn.ExecContext(context.Background(), `DELETE FROM contact WHERE id = ?`, contactId)
	if err != nil {
		return err
	}
	return nil
}
