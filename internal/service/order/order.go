package order

import (
	"fmt"

	"github.com/ufukty/play/internal/common/cache"
	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/repository"
	"github.com/ufukty/play/internal/service/email"
	"github.com/ufukty/play/internal/service/user"
)

type Order struct{}

func New(r *repository.Repository, user *user.User, email *email.Email, cache *cache.Cache, l *logger.Logger) *Order {
	fmt.Println("Constructing Service/Order")
	return &Order{}
}
