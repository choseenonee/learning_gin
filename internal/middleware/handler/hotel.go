package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/niumandzi/nto2022/model"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error" example:"error"`
}

// @Summary Create a new hotel
// @Description Create a new hotel with the input payload
// @ID create-hotel
// @Accept  json
// @Produce  json
// @Param   input body model.Hotel true "Hotel Payload"
// @Success 200 {object} map[string]int "Successfully created hotel with ID"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} map[string]string "Internal server error"
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
