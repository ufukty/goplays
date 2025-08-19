package order

import (
	"fmt"

	"github.com/ufukty/goplays/internal/common/cache"
	"github.com/ufukty/goplays/internal/common/logger"
	"github.com/ufukty/goplays/internal/repository"
	"github.com/ufukty/goplays/internal/service/email"
	"github.com/ufukty/goplays/internal/service/user"
)

type Order struct{}

func New(r *repository.Repository, user *user.User, email *email.Email, cache *cache.Cache, l *logger.Logger) *Order {
	fmt.Println("Constructing Service/Order")
	return &Order{}
}
