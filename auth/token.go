package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func getPrivateToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key, found := os.LookupEnv("JWT_PRIVATE_KEY")
	if found {

		return key
	}
	panic("JWT_PRIVATE_TOKEN não foi definido")
}

var jwtPrivateToken = getPrivateToken()

func GenerateToken(claims *JwtClaims, expirationTIme time.Time) (string, error) {
	claims.ExpiresAt = expirationTIme.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string) (bool, *JwtClaims) {
	claims := &JwtClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func getTokenFromString(tokenString string, claims *JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("erro ao verificar o token")
		}
		return []byte(jwtPrivateToken), nil
	})
}
