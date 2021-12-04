package main

import (
	"fmt"
	"log"

	"github.com/marius004/url-shortener/api"
	"github.com/marius004/url-shortener/database"
	"github.com/marius004/url-shortener/internal"
	"github.com/marius004/url-shortener/internal/services"
)

type Server struct {
	db     *database.Database
	config *internal.Config
	logger *log.Logger

	services *services.Services
}

func (s *Server) Serve() {
	api := &api.API{}
	fmt.Println(api)
}

func NewServer(config *internal.Config, db *database.Database, logger *log.Logger) *Server {
	var (
		userService     = db.UserService()
		shortUrlService = db.ShortUrlService()

		services = services.New(userService, shortUrlService)
	)

	return &Server{
		config:   config,
		db:       db,
		logger:   logger,
		services: services,
	}
}
