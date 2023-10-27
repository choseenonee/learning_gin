package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/niumandzi/nto2022/internal/middleware/handler"
	"github.com/niumandzi/nto2022/internal/usecase"
)

func NewClient(cases *usecase.UseCases) *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler(cases)

	r.POST("/create_hotel", h.CreateHotel)

	return r
}
