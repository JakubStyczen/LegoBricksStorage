package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONResponse_Success(t *testing.T) {
	// Arrange
	data := map[string]string{"message": "success"}
	recorder := httptest.NewRecorder()
	status := http.StatusOK

	// Act
	err := WriteJSONResponse(recorder, status, data)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if recorder.Code != status {
		t.Errorf("Expected status %d, got %d", status, recorder.Code)
	}
	var responseData map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseData); err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}
	if responseData["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", responseData["message"])
	}
}

func TestWriteJSONResponse_MarshalError(t *testing.T) {
	// Arrange
	data := make(chan int) // Channels cannot be marshalled to JSON
	recorder := httptest.NewRecorder()
	status := http.StatusOK

	// Act
	err := WriteJSONResponse(recorder, status, data)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
