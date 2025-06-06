package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port    int
	addr_ip string
	db      Service
}

func NewServer() *http.Server {
	addr_ip := os.Getenv("ADDR")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:    port,
		addr_ip: addr_ip,
		db:      NewService(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%d", NewServer.addr_ip, NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) GetDBQueries() *database.Queries {
	return s.db.GetDBQueries()
}
