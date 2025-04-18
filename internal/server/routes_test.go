package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

func TestHandler(t *testing.T) {
	s := &Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

type MockDB struct {
	MockHealth          func() map[string]string
	MockGetDBQueries    func() *database.Queries
	MockClose           func() error
	MockCreateUser      func(ctx any, p database.CreateUserParams) (database.User, error)
	MockGetUserByAPIKey func(ctx any, apiKey string) (database.User, error)
}

// --- implementacja interfejsu Service ---

func (m *MockDB) Health() map[string]string {
	if m.MockHealth != nil {
		return m.MockHealth()
	}
	return map[string]string{"status": "mock", "message": "mock health"}
}

func (m *MockDB) GetDBQueries() *database.Queries {
	if m.MockGetDBQueries != nil {
		return m.MockGetDBQueries()
	}
	return nil
}

func (m *MockDB) Close() error {
	if m.MockClose != nil {
		return m.MockClose()
	}
	return nil
}

// --- dodatkowe metody, jeżeli testy wymagają konkretnej logiki ---

func (m *MockDB) CreateUser(ctx any, p database.CreateUserParams) (database.User, error) {
	if m.MockCreateUser != nil {
		return m.MockCreateUser(ctx, p)
	}
	return database.User{}, errors.New("mock CreateUser not implemented")
}

func (m *MockDB) GetUserByAPIKey(ctx any, key string) (database.User, error) {
	if m.MockGetUserByAPIKey != nil {
		return m.MockGetUserByAPIKey(ctx, key)
	}
	return database.User{}, errors.New("mock GetUserByAPIKey not implemented")
}
func TestCreateUserHandler(t *testing.T) {
	mock := &MockDB{}

	mock.MockCreateUser = func(ctx any, p database.CreateUserParams) (database.User, error) {
		return database.User{
			ID:        uuid.New(),
			Name:      p.Name,
			Age:       p.Age,
			ApiKey:    "mock-api-key",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	// Proxy typu *database.Queries, które wywołuje metody MockDB
	mock.MockGetDBQueries = func() *database.Queries {
		return &database.Queries{
			// tutaj nic nie robimy – bo handler i tak nie wywołuje bezpośrednio metod Queries,
			// tylko my to przechwytujemy w mocku (taki hack).
		}
	}

	s := &Server{db: mock}

	ts := httptest.NewServer(http.HandlerFunc(s.handlerCreateUser))
	defer ts.Close()

	body := `{"Name": "John", "Age": 25}`
	resp, err := http.Post(ts.URL, "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", resp.StatusCode)
	}

	var userResp User
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &userResp)
	if err != nil {
		t.Fatalf("error unmarshaling response: %v", err)
	}

	if userResp.Name != "John" || userResp.Age != 25 {
		t.Errorf("unexpected user data: %+v", userResp)
	}
}

// func TestGetUserHandler(t *testing.T) {
// 	mockDB := &MockDB{} // <- zamień na swój mock interfejsu DB
// 	apiKey := "123-test-api-key"
// 	mockUser := database.User{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		Name:      "Anna",
// 		Age:       30,
// 		ApiKey:    apiKey,
// 	}

// 	mockDB.MockGetUserByAPIKey = func(_ any, key string) (database.User, error) {
// 		if key == apiKey {
// 			return mockUser, nil
// 		}
// 		return database.User{}, sql.ErrNoRows
// 	}

// 	s := &Server{db: mockDB}

// 	req := httptest.NewRequest("GET", "/users", nil)
// 	req.Header.Set("Authorization", "ApiKey "+apiKey)

// 	w := httptest.NewRecorder()
// 	s.handlerGetUser(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK, got %v", resp.StatusCode)
// 	}

// 	var userResp User
// 	data, _ := io.ReadAll(resp.Body)
// 	err := json.Unmarshal(data, &userResp)
// 	if err != nil {
// 		t.Fatalf("error unmarshaling: %v", err)
// 	}

// 	if userResp.Name != "Anna" {
// 		t.Errorf("expected user name Anna, got %v", userResp.Name)
// 	}
// }
