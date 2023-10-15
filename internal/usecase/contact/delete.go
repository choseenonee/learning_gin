package contact

import "context"

func (c ContactUsecase) DeleteContact(ctx context.Context, contactId int) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Delete(ctx, contactId)
	if err != nil {
		return err
	}
	return nil
}
