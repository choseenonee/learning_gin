package contact

import (
	"github.com/niumandzi/nto2022/internal/repository"
	"github.com/niumandzi/nto2022/pkg/logging"
	"time"
)

type ContactUsecase struct {
	contactRepo    repository.ContactRepository
	contextTimeout time.Duration
	logger         logging.Logger
}

func NewContacUsecase(contact repository.ContactRepository, timeout time.Duration, logger logging.Logger) ContactUsecase {
	return ContactUsecase{
		contactRepo:    contact,
		contextTimeout: timeout,
		logger:         logger,
	}
}
