package server

import (
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/auth"
	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (s *Server) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			WriteJSONResponse(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := s.GetDBQueries().GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			WriteJSONResponse(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
