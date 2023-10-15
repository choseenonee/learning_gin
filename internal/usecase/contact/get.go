package contact

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (c ContactUsecase) GetContact(ctx context.Context, contactId int) (model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contact, err := c.contactRepo.GetById(ctx, contactId)
	if err != nil {
		c.logger.Error(err.Error())
		return model.Contact{}, err
	}
	return contact, nil
}
