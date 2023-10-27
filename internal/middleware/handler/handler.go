package handler

import "github.com/niumandzi/nto2022/internal/usecase"

type Handler struct {
	cases *usecase.UseCases
}

func NewHandler(cases *usecase.UseCases) *Handler {
	return &Handler{cases: cases}
}
