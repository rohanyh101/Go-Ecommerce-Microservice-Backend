package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Fatalf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be non-empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Error("error hashing password")
	}

	if ok := ComparePasswords(hash, []byte("password")); !ok {
		t.Error("expected password to match with hash")
	}

	if ok := ComparePasswords(hash, []byte("wrong_password")); ok {
		t.Error("expected password to not match with hash")
	}
}
