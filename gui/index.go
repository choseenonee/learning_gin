package gui

import (
	"context"
	"fmt"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func Index(ctx context.Context, cases *usecase.UseCases) {
	contactData := model.Contact{
		ContactType: "worker",
		Name:        "John Doe",
		Number:      "+7 (999) 856-23-23",
		Email:       "john.doe@example.com",
	}
	id, err := cases.Contact.CreateContact(ctx, contactData)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(id)

	contact, err := cases.Contact.GetContact(ctx, 1)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("ID: %d, Type: %s, Name: %s, Number: %s, Email: %s\n",
		contact.Id, contact.ContactType, contact.Name, contact.Number, contact.Email)

	contactsByType, err := cases.Contact.GetContactsByType(ctx, "worker")
	if err != nil {
		println(err.Error())
	}
	for _, contactsByType := range contactsByType {
		fmt.Printf("ID: %d, Type: %s, Name: %s, Number: %s, Email: %s\n",
			contactsByType.Id, contactsByType.ContactType, contactsByType.Name, contactsByType.Number, contactsByType.Email)
	}

	contactsALL, err := cases.Contact.GetAllContacts(ctx)
	if err != nil {
		println(err.Error())
	}
	for _, contactsALL := range contactsALL {
		fmt.Printf("ID: %d, Type: %s, Name: %s, Number: %s, Email: %s\n",
			contactsALL.Id, contactsALL.ContactType, contactsALL.Name, contactsALL.Number, contactsALL.Email)
	}

	updateInput := model.UpdateContactInput{
		ContactType: "worker",
		Name:        "John Doe",
		Number:      "1234567890",
		Email:       "john.doe@example.com",
	}
	err = cases.Contact.UpdateContact(ctx, 5, updateInput)
	if err != nil {
		println(err.Error())
	}

	err = cases.Contact.DeleteContact(ctx, id)
	if err != nil {
		println(err.Error())
	}
}
