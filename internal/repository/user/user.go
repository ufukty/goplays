package user

import (
	"fmt"

	"github.com/ufukty/play/internal/common/database"
	"github.com/ufukty/play/internal/common/logger"
)

type User struct{}

func New(db *database.Pool, l *logger.Logger) *User {
	fmt.Println("Constructing Repository/User")
	return &User{}
}
