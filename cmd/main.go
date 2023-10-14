package main

import (
	"github.com/niumandzi/nto2022/internal/config"
	"github.com/niumandzi/nto2022/internal/repository/contact"
	"github.com/niumandzi/nto2022/pkg/config"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/niumandzi/nto2022/pkg/sqlitedb"
	"time"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	cfg, err := config.InitConfig()

	db, err := sqlitedb.NewClient(cfg.DriverName, cfg.FilePath)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = sqlitedb.CreateTables(db)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	timeoutContext := time.Duration(cfg.Timeout) * time.Second
	contactRepo := contact.NewContactRepository(db, logger)
	contactUseCase := contact.NewContacUsecase(contactRepo, timeoutContext, logger)

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
