package service

import (
	"fmt"

	"github.com/ufukty/goplays/internal/common/cache"
	"github.com/ufukty/goplays/internal/common/config"
	"github.com/ufukty/goplays/internal/common/logger"
	"github.com/ufukty/goplays/internal/repository"
	"github.com/ufukty/goplays/internal/service/email"
	"github.com/ufukty/goplays/internal/service/order"
	"github.com/ufukty/goplays/internal/service/user"
)

type Service struct {
	User  *user.User
	Email *email.Email
	Order *order.Order
}

func New(l *logger.Logger, cfg *config.Config, cache *cache.Cache, r *repository.Repository) *Service {
	fmt.Println("Constructing Service")
	var (
		email = email.New(cfg, l)
		user  = user.New(r, email, cache, l)
		order = order.New(r, user, email, cache, l)
	)
	return &Service{
		Email: email,
		User:  user,
		Order: order,
	}
}
