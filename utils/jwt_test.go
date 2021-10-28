package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"
)

func Test_generateAuthToken(t *testing.T) {
	UserID := "someNumber"

	key, errGen := rsa.GenerateKey(rand.Reader, 2048)
	if errGen != nil {
		t.Errorf("could not generate private key file: %v", errGen)
		return
	}

	signedData, err := GenerateAuthToken(UserID, key)
	if err != nil {
		t.Errorf("generateAuthToken() error = %v", err)
		return
	}

	parsed, err := ParseAuthToken(signedData, &key.PublicKey)
	if err != nil {
		t.Errorf("parseAuthToken() error = %v", err)
		return
	}

	if parsed.UserID != UserID {
		t.Errorf("parsed.UserID got = %v, want %v", parsed.UserID, UserID)
		return
	}

	if parsed.ExpiresAt < time.Now().Add(time.Hour*24*6).Unix() {
		t.Errorf("parsed.ExpiresAt should be more than 1 week got = %d", parsed.ExpiresAt)
	}

}

func Test_generateAuthTokenWithDifferentKey(t *testing.T) {
	UserID := "someNumber"

	key, errGen := rsa.GenerateKey(rand.Reader, 2048)
	if errGen != nil {
		t.Errorf("could not generate private key file: %v", errGen)
		return
	}
	signedData, err := GenerateAuthToken(UserID, key)
	if err != nil {
		t.Errorf("generateAuthToken() error = %v", err)
		return
	}

	differentKey, errGen := rsa.GenerateKey(rand.Reader, 2048)
	if errGen != nil {
		t.Errorf("could not generate private key file: %v", errGen)
		return
	}

	_, err = ParseAuthToken(signedData, &differentKey.PublicKey)
	if err == nil {
		t.Error("parseAuthToken() wanted not nil error")
		return
	}

}
