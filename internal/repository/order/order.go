package order

import (
	"fmt"

	"github.com/ufukty/goplays/internal/common/database"
	"github.com/ufukty/goplays/internal/common/logger"
)

type Order struct{}

func New(db *database.Pool, l *logger.Logger) *Order {
	fmt.Println("Constructing Repository/Order")
	return &Order{}
}
