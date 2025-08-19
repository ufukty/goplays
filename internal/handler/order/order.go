package order

import (
	"fmt"
	"net/http"

	"github.com/ufukty/goplays/internal/common/logger"
	"github.com/ufukty/goplays/internal/service"
)

type Order struct {
	Logger  *logger.Logger
	Service *service.Service
}

func New(l *logger.Logger, s *service.Service) *Order {
	fmt.Println("Constructing Handler/Order")
	return &Order{
		Logger:  l,
		Service: s,
	}
}

func (u Order) Register(s *http.ServeMux) {
	fmt.Println("Registering routes for Handler/Order")
}
