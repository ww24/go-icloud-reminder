package setup

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

var key = make([]byte, 256)

func init() {
	size, err := rand.Read(key)
	if err != nil {
		log.Fatalf("failed to generate key: %+v", err)
	}
	if size < len(key) {
		log.Fatalf("failed to generate key (length shortage): %+v", err)
	}
}

func GenerateToken(message []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	token := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(token)
}

func VerifyToken(message []byte, csrfToken string) bool {
	token, err := base64.RawURLEncoding.DecodeString(csrfToken)
	if err != nil {
		return false
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expected := mac.Sum(nil)
	return hmac.Equal(token, expected)
}

func RandomBytes() ([]byte, error) {
	data := make([]byte, 128)
	size, err := rand.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %+v", err)
	}
	if size < len(data) {
		return nil, fmt.Errorf("failed to generate key (length shortage): %+v", err)
	}
	return data, nil
}
