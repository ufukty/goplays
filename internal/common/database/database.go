package database

import "fmt"

type Pool struct{}

func New(conn string) *Pool {
	fmt.Println("Constructing Database Pool")
	return &Pool{}
}
