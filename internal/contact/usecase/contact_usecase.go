package usecase

import (
	"context"
	"github.com/niumandzi/nto2022/internal/domain"
	"github.com/niumandzi/nto2022/model"
	"github.com/niumandzi/nto2022/pkg/logging"
	"time"
)

type contactUsecase struct {
	contactRepo    domain.ContactRepository
	contextTimeout time.Duration
	logger         logging.Logger
}

func NewContacUsecase(contact domain.ContactRepository, timeout time.Duration, logger logging.Logger) domain.ContactUsecase {
	return &contactUsecase{
		contactRepo:    contact,
		contextTimeout: timeout,
		logger:         logger,
	}
}

func (c *contactUsecase) CreateContact(ctx context.Context, contact model.Contact) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	id, err := c.contactRepo.Create(ctx, contact)
	if err != nil {
		c.logger.Error(err.Error())
		return 0, err
	}
	return id, nil
}

func (c *contactUsecase) GetContact(ctx context.Context, contactId int) (model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contact, err := c.contactRepo.Get(ctx, contactId)
	if err != nil {
		c.logger.Error(err.Error())
		return model.Contact{}, err
	}
	return contact, nil
}

func (c *contactUsecase) GetContactsByType(ctx context.Context, contactType string) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contacts, err := c.contactRepo.GetByType(ctx, contactType)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	return contacts, nil
}

func (c *contactUsecase) GetAllContacts(ctx context.Context) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contacts, err := c.contactRepo.GetAll(ctx)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	return contacts, nil
}

func (c *contactUsecase) UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Update(ctx, contactId, contactInput)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}
	return nil
}

func (c *contactUsecase) DeleteContact(ctx context.Context, contactId int) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Delete(ctx, contactId)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}
	return nil
}
