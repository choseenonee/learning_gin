package contact

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/niumandzi/nto2022/model"
	"regexp"
)

func (c ContactUsecase) UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := validation.ValidateStruct(&contactInput,
		validation.Field(&contactInput.ContactType, validation.Required, validation.In("worker", "private_client", "legal_client")),
		validation.Field(&contactInput.Name, validation.Required),
		validation.Field(&contactInput.Number, validation.Required, validation.Match(regexp.MustCompile(`^\+\d{1,2}\s\(\d{3}\)\s\d{3}-\d{2}-\d{2}$`))),
		validation.Field(&contactInput.Email, validation.Required, is.Email),
	)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}

	err = c.contactRepo.Update(ctx, contactId, contactInput)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}
	return nil
}
