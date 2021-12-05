package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtWrapper struct {
	SecretKey string
	Subject   string
}

type JwtClaim struct {
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken() (signedToken string, err error) {
	claims := &JwtClaim{
		jwt.StandardClaims{
			Subject:   j.Subject,
			IssuedAt:  time.Now().Local().Unix(),
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return
}
