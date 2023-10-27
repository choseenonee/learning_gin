package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/niumandzi/nto2022/model"
	"net/http"
	"time"
)

// @Summary create hotel
// @ID create-hotel
// @Produce json
// @Success 200 {int} id
// @Router /create_hotel [post]

func (h *Handler) CreateHotel(c *gin.Context) {
	var input model.Hotel

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)

	id, err := h.cases.Hotel.CreateHotel(ctx, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
