package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
)

func (s *Server) handlerCreateLegoPartSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SetSerial  string `json:"set_serial_number"`
		PartSerial string `json:"part_serial_number"`
		Quantity   int32  `json:"quantity"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	legoPartSet, err := s.GetDBQueries().AddPartToSetBySerial(r.Context(), database.AddPartToSetBySerialParams{
		SetSerial:  params.SetSerial,
		PartSerial: params.PartSerial,
		Quantity:   params.Quantity,
	})
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("Couldn't add Lego part: %s to set: %s", params.PartSerial, params.SetSerial)
		WriteJSONResponse(w, http.StatusInternalServerError, msg)
		return
	}
	WriteJSONResponse(w, http.StatusOK, legoPartSet)
}

func (s *Server) handlerGetLegoPartForSet(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SetSerial string `json:"set_serial_number"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	legoPartsForSet, err := s.GetDBQueries().GetPartsOfSetBySerial(r.Context(), params.SetSerial)
	if err != nil {
		msg := fmt.Sprintf("Couldn't find Lego parts for set: %s", params.SetSerial)
		WriteJSONError(w, http.StatusNotFound, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoPartsForSet)
}

func (s *Server) handlerListLegoAllPartsForAllSets(w http.ResponseWriter, r *http.Request) {
	legoPartsInSets, err := s.GetDBQueries().GetAllPartsInAllSets(r.Context())
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Couldn't retrieve user's Lego parts and sets")
		return
	}

	WriteJSONResponse(w, http.StatusOK, legoPartsInSets)
}

func (s *Server) handlerDeleteLegoPartForSet(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SetSerial  string `json:"set_serial_number"`
		PartSerial string `json:"part_serial_number"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	err = s.GetDBQueries().RemovePartFromSet(r.Context(), database.RemovePartFromSetParams{
		SetSerial:  params.SetSerial,
		PartSerial: params.PartSerial,
	})
	if err != nil {
		msg := fmt.Sprintf("Couldn't delete lego part: %s for set: %s", params.PartSerial, params.SetSerial)
		WriteJSONResponse(w, http.StatusInternalServerError, msg)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Lego part for set deleted successfully")
}
