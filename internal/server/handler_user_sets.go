package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

func (s *Server) handlerCreateUserSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
		Price        string `json:"price"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	userSet, err := s.GetDBQueries().AddUserSet(r.Context(), database.AddUserSetParams{
		ID:           uuid.New(),
		SerialNumber: params.SerialNumber,
		Price:        params.Price,
		UserID:       user.ID,
	})
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create lego set")
		return
	}
	WriteJSONResponse(w, http.StatusOK, userSet)
}

func (s *Server) handlerGetUserSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	userSet, err := s.GetDBQueries().GetUserSetBySerialNumber(r.Context(), database.GetUserSetBySerialNumberParams{
		SerialNumber: params.SerialNumber,
		ID:           user.ID,
	})
	if err != nil {
		WriteJSONError(w, http.StatusNotFound, "Couldn't find lego set")
		return
	}

	WriteJSONResponse(w, http.StatusOK, userSet)
}

func (s *Server) handlerListUserSets(w http.ResponseWriter, r *http.Request, user database.User) {
	userSets, err := s.GetDBQueries().ListUserSets(r.Context(), user.ID)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve lego sets")
		return
	}

	WriteJSONResponse(w, http.StatusOK, userSets)
}

func (s *Server) handlerDeleteUserSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	err = s.GetDBQueries().DeleteUserSetBySerial(r.Context(), database.DeleteUserSetBySerialParams{
		SerialNumber: params.SerialNumber,
		ID:           user.ID,
	})
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't delete lego set")
		return
	}

	WriteJSONResponse(w, http.StatusOK, "User Lego set deleted successfully")
}
