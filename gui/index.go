package gui

import (
	"context"
	"fmt"
	"github.com/niumandzi/nto2022/internal/usecase"
)

func Index(ctx context.Context, cases *usecase.UseCases) {
	contacts, err := cases.Contact.GetAllContacts(ctx)
	if err != nil {
		println(err.Error())
	}

	for _, contacts := range contacts {
		fmt.Printf("ID: %d, Type: %s, Name: %s, Number: %s, Email: %s\n",
			contacts.Id, contacts.ContactType, contacts.Name, contacts.Number, contacts.Email)
	}
}
