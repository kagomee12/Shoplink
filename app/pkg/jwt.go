package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTSecret string
type JWTIssuer string

func NewJWTSecret() JWTSecret {
    return JWTSecret(os.Getenv("JWT_SECRET"))
}

func NewJWTIssuer() JWTIssuer {
    return JWTIssuer(os.Getenv("JWT_ISSUER"))
}

type JWTService interface {
	GenerateToken(userID uint, username string) (string, error)
	ValidateToken(token string) (*JWTClaims, error)
}

type JWTServiceImpl struct {
	JWTSecret string
	JWTIssuer string
}

func NewJWTService(secret JWTSecret, issuer JWTIssuer) *JWTServiceImpl {
	return &JWTServiceImpl{
		JWTSecret: string(secret),
		JWTIssuer: string(issuer),
	}
}

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j *JWTServiceImpl) GenerateToken(userID uint, username string) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.JWTIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTServiceImpl) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err

}
