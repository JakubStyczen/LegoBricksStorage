package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

func (s *Server) handlerCreateLegoPart(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SerialNumber string `json:"serial_number"`
		Quantity     int32  `json:"quantity"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		Price        string `json:"price"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	legoPart, err := s.GetDBQueries().CreatePart(r.Context(), database.CreatePartParams{
		ID:           uuid.New(),
		SerialNumber: params.SerialNumber,
		Name:         params.Name,
		Quantity:     params.Quantity,
		Color:        params.Color,
		Price:        params.Price,
	})
	if err != nil {
		log.Println(err)
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create lego part")
		return
	}
	WriteJSONResponse(w, http.StatusOK, legoPart)
}

func (s *Server) handlerGetLegoPart(w http.ResponseWriter, r *http.Request) {
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
	legoPart, err := s.GetDBQueries().GetPartByNumber(r.Context(), params.SerialNumber)
	if err != nil {
		WriteJSONError(w, http.StatusNotFound, "Couldn't find lego part")
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoPart)
}

func (s *Server) handlerListLegoParts(w http.ResponseWriter, r *http.Request) {
	legoParts, err := s.GetDBQueries().ListParts(r.Context())
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve lego parts")
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoParts)
}

func (s *Server) handlerUpdateLegoPart(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		SerialNumber string `json:"serial_number"`
		Quantity     int32  `json:"quantity"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		Price        string `json:"price"`
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	// todo: check if part exists
	err := s.GetDBQueries().UpdatePart(r.Context(), database.UpdatePartParams{
		SerialNumber: params.SerialNumber,
		Name:         params.Name,
		Quantity:     params.Quantity,
		Color:        params.Color,
		Price:        params.Price,
	})
	if err != nil {
		log.Println("update error:", err)
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't update lego part")
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego part updated successfully")
}

func (s *Server) handlerDeleteLegoPart(w http.ResponseWriter, r *http.Request, user database.User) {
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
	err = s.GetDBQueries().DeletePart(r.Context(), params.SerialNumber)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't delete lego part")
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego part deleted successfully")
}
