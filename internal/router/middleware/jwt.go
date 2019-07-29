package middleware

import (
	"fmt"
	"net/http"
	"time"

	"miu-auth-api-v1/internal/platform"
	"miu-auth-api-v1/internal/revokedToken"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper      func(c echo.Context) bool
	jwtExtractor func(echo.Context) (string, error)
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

func JWT(key interface{}, rts revokedToken.Store) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c, rts)
}

func JWTWithConfig(config JWTConfig, rts revokedToken.Store) echo.MiddlewareFunc {
	extractor := jwtFromHeader("Authorization", "Bearer")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)
			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, platform.NewHttpError(err))
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return c.JSON(http.StatusForbidden, platform.NewHttpError(ErrJWTInvalid))
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				accountID := uint(claims["id"].(float64))
				jti, err := uuid.Parse(claims["jti"].(string))
				if err != nil {
					c.JSON(http.StatusForbidden, platform.NewHttpError(ErrJWTInvalid))
				}

				et, err := rts.GetByJti(jti)

				if et != nil {
					return c.JSON(http.StatusForbidden, platform.NewHttpError(ErrJWTInvalid))
				}

				if err != nil {
					return c.JSON(http.StatusForbidden, platform.NewHttpError(ErrJWTInvalid))
				}

				exp := time.Unix(int64(claims["exp"].(float64)), 0)
				c.Set("accountId", accountID)
				c.Set("tokenJti", jti)
				c.Set("tokenExp", exp)
				c.Set("token", token.Raw)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, platform.NewHttpError(ErrJWTInvalid))
		}
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
