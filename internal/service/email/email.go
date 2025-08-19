package email

import (
	"fmt"

	"github.com/ufukty/play/internal/common/config"
	"github.com/ufukty/play/internal/common/logger"
)

type Email struct{}

func New(cfg *config.Config, l *logger.Logger) *Email {
	fmt.Println("Constructing Service/Email")
	return &Email{}
}
