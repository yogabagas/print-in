package middlewares

import (
	"errors"
	"fmt"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/transport/rest/handler/response"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type MiddlewareImpl struct{}

type Middleware interface {
	AuthenticationMiddleware(next http.Handler) http.Handler
}

func NewMiddleware() Middleware {
	return &MiddlewareImpl{}
}

var res = response.NewJSONResponse()

// AuthenticationMiddleware validates the JWT token.
func (mi *MiddlewareImpl) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		isRegister := r.URL.Path == "/v1/users" && r.Method == http.MethodPost

		if token == "" && !isRegister {
			res.SetError(response.ErrUnauthorized).SetMessage(errors.New("An Authorization Header is required").Error()).Send(w)
			return
		}

		if !mi.parseJwt(token) && !isRegister {
			res.SetError(response.ErrUnauthorized).SetMessage(errors.New("Invalid Authorized Token").Error()).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mi *MiddlewareImpl) parseJwt(authorizationHeader string) (valid bool) {
	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) == 2 {
		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(config.GlobalCfg.App.JwtSecret), nil
		})

		return token.Valid
	}
	return false
}
