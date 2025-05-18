package server

import (
	"encoding/json"
	"fmt"
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
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't add Lego set to user")
		return
	}
	WriteJSONResponse(w, http.StatusOK, databaseUserSetToUserSet(userSet))
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
		msg := fmt.Sprintf("Couldn't find Lego set SN: %s for user", params.SerialNumber)
		WriteJSONError(w, http.StatusNotFound, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, databaseUserSetToUserSet(userSet))
}

func (s *Server) handlerListUserSets(w http.ResponseWriter, r *http.Request, user database.User) {
	userSets, err := s.GetDBQueries().ListUserSets(r.Context(), user.ID)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve user's Lego sets")
		return
	}

	WriteJSONResponse(w, http.StatusOK, ConvertDBObjListToObjList(userSets, databaseUserSetToUserSet))
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
		msg := fmt.Sprintf("Couldn't delete lego set SN: %s for user", params.SerialNumber)
		WriteJSONResponse(w, http.StatusInternalServerError, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "User's Lego set deleted successfully")
}
