package app

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	DB *sql.DB
	Fiber *fiber.App
}

func NewServer() *Server {
	return &Server{
		DB: NewDatabase(),
		Fiber: fiber.New(),
	}
}

func (server *Server ) Start() error {
	return server.Fiber.Listen(":3000")
}