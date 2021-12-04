package main

import (
	"log"

	"github.com/marius004/url-shortener/database"
	"github.com/marius004/url-shortener/internal"
	"github.com/marius004/url-shortener/internal/services"
	"github.com/marius004/url-shortener/models"
)

type Server struct {
	db     *database.Database
	config *internal.Config
	logger *log.Logger

	services *services.Services
}

func (s *Server) Serve() {
	s.db.AutoMigrate(&models.User{})
}

func NewServer(config *internal.Config, db *database.Database, logger *log.Logger) *Server {
	var (
		userService = db.UserService()
		services    = services.New(userService)
	)

	return &Server{
		config:   config,
		db:       db,
		logger:   logger,
		services: services,
	}
}
