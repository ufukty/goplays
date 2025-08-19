package repository

import (
	"fmt"

	"github.com/ufukty/play/internal/common/database"
	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/repository/order"
	"github.com/ufukty/play/internal/repository/user"
)

type Repository struct {
	User  *user.User
	Order *order.Order
}

func New(l *logger.Logger, db *database.Pool) *Repository {
	fmt.Println("Constructing Repository")
	var (
		order = order.New(db, l)
		user  = user.New(db, l)
	)
	return &Repository{
		User:  user,
		Order: order,
	}
}
