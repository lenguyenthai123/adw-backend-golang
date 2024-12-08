package jwt

import (
	"testing"
	"time"
	"trekkstay/pkgs/jwt/constant"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	provider := NewJWT("secret")

	t.Run("Valid token", func(t *testing.T) {
		// Create a valid token
		claims := jwt.MapClaims{
			"user_id": "123",
			"role":    "user",
			"exp":     time.Now().Add(time.Second * time.Duration(50)).Unix(), // Set the expiry time
			"iat":     time.Now().Unix(),                                      // Set the token creation time
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := jwtToken.SignedString([]byte("secret"))

		// Call the Validate function
		payload, err := provider.Validate(tokenString)

		// Check if the payload and error are as expected
		assert.NoError(t, err)
		assert.Equal(t, "123", payload.UserID)
	})

	t.Run("Invalid token", func(t *testing.T) {
		// Create an invalid token
		tokenString := "invalid_token"

		// Call the Validate function
		payload, err := provider.Validate(tokenString)

		// Check if the error is as expected
		assert.Equal(t, constant.ErrCannotUnmarshalToken, err)
		assert.Nil(t, payload)
	})
}

func TestGenerateToken(t *testing.T) {
	provider := NewJWT("secret")

	// Test case 1: Generate token with valid data and expiry
	t.Run("Generate token with valid data and expiry", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_id": "user1",
			"role":    "user",
		}
		expiry := 3600

		token, err := provider.Generate(payload, expiry)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Verify the generated token
		if token["token"] == "" {
			t.Errorf("Expected non-empty token, got empty token")
		}
		if token["expiry"] != expiry {
			t.Errorf("Expected expiry %d, got %d", expiry, token["expiry"])
		}
		if token["created_at"].(time.Time).IsZero() {
			t.Errorf("Expected non-zero creation time, got zero time")
		}
	})

	// Test case 2: Generate token with empty data and expiry
	t.Run("Generate token with empty data and expiry", func(t *testing.T) {
		payload := map[string]interface{}{}
		expiry := 0

		token, err := provider.Generate(payload, expiry)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Verify the generated token
		if token["token"] == "" {
			t.Errorf("Expected non-empty token, got empty token")
		}
		if token["expiry"] != expiry {
			t.Errorf("Expected expiry %d, got %d", expiry, token["expiry"])
		}
		if token["created_at"].(time.Time).IsZero() {
			t.Errorf("Expected non-zero creation time, got zero time")
		}
	})
}
