package handler

import (
	"fmt"
	"net/http"

	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/handler/order"
	"github.com/ufukty/play/internal/handler/user"
	"github.com/ufukty/play/internal/service"
)

type Handler struct {
	User  *user.User
	Order *order.Order
}

func New(l *logger.Logger, s *service.Service) *Handler {
	fmt.Println("Constructing Handler")
	var (
		user  = user.New(l, s)
		order = order.New(l, s)
	)
	return &Handler{
		User:  user,
		Order: order,
	}
}

func (h Handler) Register(s *http.ServeMux) {
	fmt.Println("Registering routes")
	h.Order.Register(s)
	h.User.Register(s)
}
