package helper

import (
	"fmt"
	"gostore/entity"
	authError "gostore/utils/helper/domain/errorModel"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id       int
	Username string
	Email    string
	Role     string
	jwt.RegisteredClaims
}

type Key string

const (
	GOSTORE_USERID    Key = "user-store-id"
	GOSTORE_USERNAME  Key = "user-store-user"
	GOSTORE_USEREMAIL Key = "user-store-email"
	GOSTORE_USERROLE  Key = "user-store-role"
)

func GenerateToken(user entity.UserResponse, jwtExpTime time.Time) (string, error) {
	// Create token claim
	claims := &JWTClaim{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-store",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
		},
	}

	// Generate JWT Token
	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenAlgorithm.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func GetTokenFromHeader(w http.ResponseWriter, r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return "", false
	}

	return strings.Replace(authHeader, "Bearer ", "", -1), true
}

func ValidateJWT(ts string) (*JWTClaim, bool, error) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(ts, claims, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		v, _ := err.(jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorExpired:
			return nil, false, authError.ErrExpToken
		default:
			return nil, false, authError.ErrUnauthorized
		}
	}

	if !token.Valid {
		return nil, false, authError.ErrInvalidToken
	}

	return token.Claims.(*JWTClaim), true, nil
}

func VerifyRefreshToken(s entity.Session) (*JWTClaim, error) {
	payload, ok, err := ValidateJWT(s.Token)
	if !ok || err != nil {
		return &JWTClaim{}, err
	}

	return payload, nil
}
