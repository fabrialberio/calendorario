package auth

import (
	"calendorario/pkg/database"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type userClaims struct {
	database.User
	jwt.RegisteredClaims
}

func init() {
	privateKeyBytes, err := os.ReadFile("private_key.pem")
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	publicKeyBytes, err := os.ReadFile("public_key.pem")
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}
}

func generateJWT(userId int64, role database.Role) (string, error) {
	claims := userClaims{
		User: database.User{
			ID:   userId,
			Role: role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			Issuer:    "calendorario",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateJWT(tokenString string) (*userClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
