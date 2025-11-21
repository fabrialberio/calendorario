package session

import (
	"calendorario/pkg/database"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrCookieNotFound = errors.New("cookie not found")
	ErrCookieExpired  = errors.New("cookie is expired")
)

const sessionCookieName = "jwt"

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}

	return hash, nil
}

func GetAuthenticatedUser(r *http.Request) (*database.User, error) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil, ErrCookieNotFound
	}

	claims, err := validateJWT(cookie.Value)
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, ErrCookieExpired
	} else if err != nil {
		return nil, err
	}

	return &claims.User, nil
}

func SetAuthenticatedUser(w http.ResponseWriter, user *database.User, password []byte) error {
	err := bcrypt.CompareHashAndPassword(user.PasswordHash, password)
	if err != nil {
		return err
	}

	token, err := generateJWT(user.ID, user.Role)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func UnsetAuthenticatedUser(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
