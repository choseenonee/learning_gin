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
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	id, err := c.contactRepo.Create(ctx, contact)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *contactUsecase) GetContact(ctx context.Context, contactId int) (model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contact, err := c.contactRepo.Get(ctx, contactId)
	if err != nil {
		return model.Contact{}, err
	}
	return contact, nil
}

func (c *contactUsecase) GetContactsByType(ctx context.Context, contactType string) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contacts, err := c.contactRepo.GetByType(ctx, contactType)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (c *contactUsecase) GetAllContacts(ctx context.Context) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contacts, err := c.contactRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (c *contactUsecase) UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Update(ctx, contactId, contactInput)
	if err != nil {
		return err
	}
	return nil
}

func (c *contactUsecase) DeleteContact(ctx context.Context, contactId int) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Delete(ctx, contactId)
	if err != nil {
		return err
	}
	return nil
}
