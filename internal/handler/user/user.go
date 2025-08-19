package user

import (
	"fmt"
	"net/http"

	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/service"
)

type User struct {
	Logger  *logger.Logger
	Service *service.Service
}

func New(l *logger.Logger, s *service.Service) *User {
	fmt.Println("Constructing Handler/User")
	return &User{
		Logger:  l,
		Service: s,
	}
}

func (u User) Register(s *http.ServeMux) {
	fmt.Println("Registering routes for Handler/User")
}
