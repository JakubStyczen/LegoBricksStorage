package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) getURLParam(w http.ResponseWriter, r *http.Request, paramName string) string {
	paramValue := chi.URLParam(r, paramName)

	if paramValue == "" {
		WriteJSONError(w, http.StatusBadRequest, paramName+" is required")
		return ""
	}
	return paramValue
}

func (s *Server) handlerCreateLegoSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
		Name         string `json:"name"`
		Price        string `json:"price"`
		Theme        string `json:"theme"`
		Year         int32  `json:"year"`
		TotalParts   int32  `json:"total_parts"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	legoSet, err := s.GetDBQueries().CreateLegoSet(r.Context(), database.CreateLegoSetParams{
		ID:           uuid.New(),
		SerialNumber: params.SerialNumber,
		Name:         params.Name,
		Price:        params.Price,
		Theme:        params.Theme,
		Year:         params.Year,
		TotalParts:   params.TotalParts,
		UserID:       user.ID,
	})
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create lego set")
		return
	}
	WriteJSONResponse(w, http.StatusOK, databaseLegoSetToLegoSet(legoSet))
}

func (s *Server) handlerGetLegoSet(w http.ResponseWriter, r *http.Request) {
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

	legoSet, err := s.GetDBQueries().GetLegoSetBySerial(r.Context(), params.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("Couldn't find Lego SN: %s", params.SerialNumber)
		WriteJSONError(w, http.StatusNotFound, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, databaseLegoSetToLegoSet(legoSet))
}

func (s *Server) handlerListLegoSets(w http.ResponseWriter, r *http.Request) {
	legoSets, err := s.GetDBQueries().ListLegoSets(r.Context())
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve lego sets")
		return
	}

	WriteJSONResponse(w, http.StatusOK, ConvertDBObjListToObjList(legoSets, databaseLegoSetToLegoSet))
}

func (s *Server) handlerUpdateLegoSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
		Name         string `json:"name"`
		Price        string `json:"price"`
		Theme        string `json:"theme"`
		Year         int32  `json:"year"`
		TotalParts   int32  `json:"total_parts"`
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	err := s.GetDBQueries().UpdateLegoSet(r.Context(), database.UpdateLegoSetParams{
		SerialNumber: params.SerialNumber,
		Name:         params.Name,
		Price:        params.Price,
		Theme:        params.Theme,
		Year:         params.Year,
		TotalParts:   params.TotalParts,
		UserID:       user.ID,
	})
	if err != nil {
		log.Println("Update error:", err)
		msg := fmt.Sprintf("Couldn't update Lego set SN: %s", params.SerialNumber)
		WriteJSONError(w, http.StatusInternalServerError, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego set updated successfully")
}

func (s *Server) handlerDeleteLegoSet(w http.ResponseWriter, r *http.Request, user database.User) {
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
	err = s.GetDBQueries().DeleteLegoSet(r.Context(), params.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("Couldn'td delte Lego set SN: %s", params.SerialNumber)
		WriteJSONResponse(w, http.StatusInternalServerError, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego set deleted successfully")
}
