package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roh4nyh/ecom/types"
)

type mockUserStore struct{}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}

func (s *mockUserStore) CeateUser(u *types.User) error {
	return nil
}

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// crete a user payload
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Password:  "password",
		}

		m, _ := json.Marshal(payload)

		// create a dummy request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(m))
		if err != nil {
			t.Fatal(err)
		}

		// record the response
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", handler.handleRegister)

		// call the handler
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should register a user", func(t *testing.T) {
		// crete a user payload
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@gmail.com",
			Password:  "password",
		}

		m, _ := json.Marshal(payload)

		// create a dummy request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(m))
		if err != nil {
			t.Fatal(err)
		}

		// record the response
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", handler.handleRegister)

		// call the handler
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}
