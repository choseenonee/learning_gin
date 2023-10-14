package main

import (
	"github.com/niumandzi/nto2022/internal/contact/repository"
	"github.com/niumandzi/nto2022/internal/contact/usecase"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/niumandzi/nto2022/pkg/sqlitedb"
	"time"
)

func main() {
	filePath := "./nto2022.db"
	driverName := "sqlite3"
	timeout := 2

	//ctx := context.Background()

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

	timeoutContext := time.Duration(timeout) * time.Second
	contactRepo := repository.NewSqliteContactRepository(db, logger)
	contactUseCase := usecase.NewContacUsecase(contactRepo, timeoutContext, logger)

	//contact, err := contactUseCase.GetContactsByType(ctx, "worker")
	//if err != nil {
	//	logger.Fatalf(err.Error())
	//}
	//
	//for _, contact := range contact {
	//	fmt.Printf("ID: %d, Type: %s, Name: %s, Number: %s, Email: %s\n",
	//		contact.Id, contact.ContactType, contact.Name, contact.Number, contact.Email)
	//}
}
