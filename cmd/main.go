package main

import (
	"fmt"
	"github.com/niumandzi/nto2022/internal/config"
	"github.com/niumandzi/nto2022/internal/middleware"
	contactRepository "github.com/niumandzi/nto2022/internal/repository/contact"
	hotelRepository "github.com/niumandzi/nto2022/internal/repository/hotel"
	"github.com/niumandzi/nto2022/internal/usecase"
	contactUsecase "github.com/niumandzi/nto2022/internal/usecase/contact"
	hotelUsecase "github.com/niumandzi/nto2022/internal/usecase/hotel"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/niumandzi/nto2022/pkg/sqlitedb"
	"time"
)

// @title Go + Gin + Swaggo
// @version 1.0
// @description Test

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	//ctx := context.Background()

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
	hotelRepo := hotelRepository.NewHotelRepository(db, logger)

	contact := contactUsecase.NewContacUsecase(contactRepo, timeoutContext, logger)
	hotel := hotelUsecase.NewHotelUsecase(hotelRepo, timeoutContext, logger)

	cases := usecase.NewUsecases(contact, hotel)

	r := middleware.NewClient(cases)

	if err := r.Run(); err != nil {
		fmt.Println("ERROR CLIENT")
	}

	//gui.Index(ctx, cases)
}
