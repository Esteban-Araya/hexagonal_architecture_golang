package jwt

import (
	"Project/internal/api"
	"context"
	"net/http"
)

func (j JwtMiddleware) MiddlewareIsJwtValid(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		map_claims, err := j.JwtService.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			api.Error(err, http.StatusNotAcceptable).Send(w)
			return
		}

		if err = j.JwtService.IsTokenExpire(map_claims["exp"].(float64)); err != nil {
			api.Error(err, http.StatusNotAcceptable).Send(w)
			return
		}
		user, err := j.JwtService.GetUser(map_claims["user_id"].(float64))
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}