package user

import (
	"fmt"

	"github.com/ufukty/play/internal/common/cache"
	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/repository"
	"github.com/ufukty/play/internal/service/email"
)

type User struct{}

func New(r *repository.Repository, email *email.Email, cache *cache.Cache, l *logger.Logger) *User {
	fmt.Println("Constructing Service/User")
	return &User{}
}
