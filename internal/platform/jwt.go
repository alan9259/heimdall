package platform

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWTToken(id int32) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().UTC().Add(time.Hour * 72).Unix()
	claims["aud"] = "MIU"
	claims["jti"] = uuid.New()
	claims["iat"] = time.Now().UTC().Unix()
	claims["iss"] = "MIU"
	t, _ := token.SignedString(JWTSecret)
	return t
}
