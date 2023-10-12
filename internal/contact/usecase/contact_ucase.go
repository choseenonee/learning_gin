package usecase

import (
	"context"
	"github.com/niumandzi/nto2022/internal/domain"
	"github.com/niumandzi/nto2022/model"
	"time"
)

type contactUsecase struct {
	contactRepo    domain.ContactRepository
	contextTimeout time.Duration
}

func NewContacUsecase(contact domain.ContactRepository, timeout time.Duration) domain.ContactUsecase {
	return &contactUsecase{
		contactRepo:    contact,
		contextTimeout: timeout,
	}
}

func (c *contactUsecase) CreateContact(ctx context.Context, contact model.Contact) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *contactUsecase) GetContact(ctx context.Context, contactId int) (model.Contact, error) {
	contact, err := c.contactRepo.Get(ctx, contactId)
	if err != nil {
		return model.Contact{}, err
	}
	return contact, nil
}

func (c *contactUsecase) GetContactByType(ctx context.Context, contactType string) (model.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *contactUsecase) GetAllContacts(ctx context.Context) ([]model.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *contactUsecase) UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	//TODO implement me
	panic("implement me")
}

func (c *contactUsecase) DeleteContact(ctx context.Context, contactId int) error {
	//TODO implement me
	panic("implement me")
}
