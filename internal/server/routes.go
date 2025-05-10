package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Post("/users", s.handlerCreateUser)
	r.Get("/users", s.middlewareAuth(s.handlerGetUser))

	r.Post("/lego/sets", s.middlewareAuth(s.handlerCreateLegoSet))
	r.Get("/lego/sets/{serial_number}", s.handlerGetLegoSet)
	r.Get("/lego/sets", s.handlerListLegoSets)
	r.Patch("/lego/sets/{serial_number}", s.middlewareAuth(s.handlerUpdateLegoSet))
	r.Delete("/lego/sets/{serial_number}", s.middlewareAuth(s.handlerDeleteLegoSet))

	r.Post("/lego/user/set", s.middlewareAuth(s.handlerCreateUserSet))
	r.Get("/lego/user/set", s.middlewareAuth(s.handlerGetUserSet))
	r.Get("/lego/user/sets", s.middlewareAuth(s.handlerListUserSets))
	r.Delete("/lego/user/set", s.middlewareAuth(s.handlerDeleteUserSet))

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	_ = WriteJSONResponse(w, 200, resp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	_ = WriteJSONResponse(w, 200, s.db.Health())
}
