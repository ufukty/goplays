package main

import (
	"net/http"

	"github.com/ufukty/goplays/internal/common/cache"
	"github.com/ufukty/goplays/internal/common/config"
	"github.com/ufukty/goplays/internal/common/database"
	"github.com/ufukty/goplays/internal/common/logger"
	"github.com/ufukty/goplays/internal/handler"
	"github.com/ufukty/goplays/internal/repository"
	"github.com/ufukty/goplays/internal/service"
)

func main() {
	var (
		cfg   = config.ReadConfig("")
		l     = logger.New("play")
		db    = database.New("psql://...")
		cache = cache.New()
	)

	var (
		r = repository.New(l, db)
		s = service.New(l, cfg, cache, r)
		h = handler.New(l, s)
	)

	router := http.NewServeMux()
	h.Register(router)
	http.ListenAndServe(":8080", router)
}
