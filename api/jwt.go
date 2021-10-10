package api

import (
	"fmt"
	"os"
	"time"

	"github.com/anfelo/chillhacks_platform/utils/errors"
	"github.com/golang-jwt/jwt/v4"
)

var signingKey = os.Getenv("JWTSECRET")

func GenerateJWT(username, uID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = uID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(token string) (string, *errors.RestErr) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", errors.NewInternatServerError("internal server error")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username := claims["username"]
		return fmt.Sprintf("%v", username), nil
	}

	return "", errors.NewUnauthorizedError("unauthorized request")

}
