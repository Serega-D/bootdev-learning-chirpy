package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
    "crypto/rand"
    "encoding/hex"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/alexedwards/argon2id"
)

// HashPassword хеширует строку пароля
func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

// CheckPasswordHash сравнивает пароль с хешем из базы
func CheckPasswordHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

// MakeJWT создает подписанный токен для userID
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	signingKey := []byte(tokenSecret)

	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// ValidateJWT проверяет подпись токена и возвращает userID из Subject
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claimsStruct := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		return uuid.Nil, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	return userID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header found")
	}

	// Ожидаем формат: "Bearer TOKEN_STRING"
	splitHandler := strings.Split(authHeader, " ")
	if len(splitHandler) < 2 || splitHandler[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitHandler[1], nil
}


func MakeRefreshToken() (string, error) {
    bytes := make([]byte, 32)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header found for GetAPIKey function")
	}
	
	// Expected format: Authorization: ApiKey THE_KEY_HERE
	splitHandler := strings.Split(authHeader, " ")
	if len(splitHandler) < 2 || splitHandler[0] != "ApiKey" {
		return "", errors.New("malformed authorization header for GetAPIKey function")
	}

	return splitHandler[1], nil
}

//  Add a func GetAPIKey(headers http.Header) (string, error) 
// 	to your auth package. It should extract the api key from the 
// 	Authorization header, which is expected to be in this format:

// Authorization: ApiKey THE_KEY_HERE
