package usecase

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

type ContactUseсase interface {
	CreateContact(ctx context.Context, contact model.Contact) (int, error)
	GetContact(ctx context.Context, contactId int) (model.Contact, error)
	GetContactsByType(ctx context.Context, contactType string) ([]model.Contact, error)
	GetAllContacts(ctx context.Context) ([]model.Contact, error)
	UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error
	DeleteContact(ctx context.Context, contactId int) error
}