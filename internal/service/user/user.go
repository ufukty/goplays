package user

import (
	"fmt"

	"github.com/ufukty/goplays/internal/common/cache"
	"github.com/ufukty/goplays/internal/common/logger"
	"github.com/ufukty/goplays/internal/repository"
	"github.com/ufukty/goplays/internal/service/email"
)

type User struct{}

func New(r *repository.Repository, email *email.Email, cache *cache.Cache, l *logger.Logger) *User {
	fmt.Println("Constructing Service/User")
	return &User{}
}
