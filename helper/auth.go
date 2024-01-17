package helper

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id       int
	Username string
	Role     string
	jwt.RegisteredClaims
}

type Key string

const (
	GOSTORE_USERID   Key = "go-store-id"
	GOSTORE_USERNAME Key = "go-store-user"
	GOSTORE_USERROLE Key = "go-store-role"
)

func GetTokenFromHeader(w http.ResponseWriter, r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return ""
	}

	return strings.Replace(authHeader, "Bearer ", "", -1)
}

func ValidateJWT(ts string) (*jwt.Token, error) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(ts, claims, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return token, err
}

func RoleBasedAccessControl(r *http.Request, allowedRole ...string) bool {
	role := r.Context().Value(GOSTORE_USERROLE)
	allow := false

	for _, allowed := range allowedRole {
		if role == allowed {
			allow = true
			break
		}
	}

	return allow
}
