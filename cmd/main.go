package main

import (
	"github.com/niumandzi/nto2022/internal/config"
	contactRepository "github.com/niumandzi/nto2022/internal/repository/contact"
	contactUsecase "github.com/niumandzi/nto2022/internal/usecase/contact"
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
	contactRepo := contactRepository.NewContactRepository(db, logger)
	contactUsecase := contactUsecase.NewContacUsecase(contactRepo, timeoutContext, logger)

}
