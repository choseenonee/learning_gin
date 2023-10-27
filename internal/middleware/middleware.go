package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/niumandzi/nto2022/internal/middleware/handler"
	"github.com/niumandzi/nto2022/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewClient(cases *usecase.UseCases) *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler(cases)

	r.POST("/create_hotel", h.CreateHotel)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
