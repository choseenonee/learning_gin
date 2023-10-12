package main

import (
	"context"
	"fmt"
	"github.com/niumandzi/nto2022/internal/contact/repository"
	"github.com/niumandzi/nto2022/internal/contact/usecase"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/niumandzi/nto2022/pkg/sqlitedb"
	"time"
)

func main() {
	filePath := "./nto2022.db"
	driverName := "sqlite3"

	ctx := context.Background()

	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	db, err := sqlitedb.NewClient(driverName, filePath)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = sqlitedb.CreateTables(db)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	timeoutContext := time.Duration(10) * time.Second
	contactRepo := repository.NewSqliteContactRepository(db)
	contactUseCase := usecase.NewContacUsecase(contactRepo, timeoutContext)
	contact, err := contactUseCase.GetContact(ctx, 1)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	fmt.Println("ID:", contact.Id)
	fmt.Println("Type:", contact.ContactType)
	fmt.Println("Name:", contact.Name)
	fmt.Println("Number:", contact.Number)
	fmt.Println("Email:", contact.Email)
}
