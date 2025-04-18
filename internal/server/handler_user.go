package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JakubStyczen/LegoBricksStorage/internal/auth"
	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

func (s *Server) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
		Age  int32
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := s.GetDBQueries().CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Age:       params.Age,
	})
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	WriteJSONResponse(w, http.StatusOK, databaseUserToUser(user))
}

func (s *Server) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		WriteJSONError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}

	user, err := s.GetDBQueries().GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		WriteJSONError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	WriteJSONResponse(w, http.StatusOK, databaseUserToUser(user))
}
