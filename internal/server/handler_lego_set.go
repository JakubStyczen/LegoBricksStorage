package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s *Server) handlerCreateLegoSet(w http.ResponseWriter, r *http.Request) {
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
		SerialNumber: params.SerialNumber,
		Name:         params.Name,
		Price:        params.Price,
		Theme:        params.Theme,
		Year:         params.Year,
		TotalParts:   params.TotalParts,
	})
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create lego set")
		return
	}
	WriteJSONResponse(w, http.StatusOK, legoSet)
}

func (s *Server) handlerGetLegoSet(w http.ResponseWriter, r *http.Request) {
	serialNumber := chi.URLParam(r, "serial_number")

	if serialNumber == "" {
		WriteJSONError(w, http.StatusBadRequest, "Serial number is required")
		return
	}

	legoSet, err := s.GetDBQueries().GetLegoSetBySerial(r.Context(), serialNumber)
	if err != nil {
		WriteJSONError(w, http.StatusNotFound, "Couldn't find lego set")
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoSet)
}

func (s *Server) handlerListLegoSets(w http.ResponseWriter, r *http.Request) {
	legoSets, err := s.GetDBQueries().ListLegoSets(r.Context())
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve lego sets")
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoSets)
}

func (s *Server) handlerUpdateLegoSet(w http.ResponseWriter, r *http.Request) {
	serialNumber := chi.URLParam(r, "serial_number")

	if serialNumber == "" {
		WriteJSONError(w, http.StatusBadRequest, "Serial number is required")
		return
	}

	type parameters struct {
		Name       string `json:"name"`
		Price      string `json:"price"`
		Theme      string `json:"theme"`
		Year       int32  `json:"year"`
		TotalParts int32  `json:"total_parts"`
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	err := s.GetDBQueries().UpdateLegoSet(r.Context(), database.UpdateLegoSetParams{
		SerialNumber: serialNumber,
		Name:         params.Name,
		Price:        params.Price,
		Theme:        params.Theme,
		Year:         params.Year,
		TotalParts:   params.TotalParts,
	})
	if err != nil {
		log.Println("update error:", err)
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't update lego set")
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego set updated successfully")
}

func (s *Server) handlerDeleteLegoSet(w http.ResponseWriter, r *http.Request) {
	serialNumber := chi.URLParam(r, "serial_number")

	if serialNumber == "" {
		WriteJSONError(w, http.StatusBadRequest, "Serial number is required")
		return
	}
	err := s.GetDBQueries().DeleteLegoSet(r.Context(), serialNumber)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't delete lego set")
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego set deleted successfully")
}
