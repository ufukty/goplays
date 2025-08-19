package main

import (
	"net/http"

	"github.com/ufukty/play/internal/common/cache"
	"github.com/ufukty/play/internal/common/config"
	"github.com/ufukty/play/internal/common/database"
	"github.com/ufukty/play/internal/common/logger"
	"github.com/ufukty/play/internal/handler"
	"github.com/ufukty/play/internal/repository"
	"github.com/ufukty/play/internal/service"
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
