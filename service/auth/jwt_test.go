package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Fatalf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be non-empty")
	}
}
