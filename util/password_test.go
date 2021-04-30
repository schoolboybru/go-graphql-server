package util

import "testing"

func TestHashPassword(t *testing.T) {
	password, err := HashPassword("tEsTin@gPasSW!0rd")

	if err != nil {
		t.Errorf("password hashing failed, got %v",err.Error())
	} else {
		t.Logf("password hashing sucess, got %v", password)
	}
}

func TestValidPasswordHash(t *testing.T) {
	password, err := HashPassword("testingPassword")

	if err != nil {
		t.Errorf("password hashing failed, got %v", err.Error())
	}

	isValid := CheckPasswordHash("testingPassword", password)

	if isValid == false {
		t.Errorf("checking password hash failed, got %v", isValid)
	}

	t.Logf("checking password hash sucess")
}