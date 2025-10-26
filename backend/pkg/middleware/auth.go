package middleware

import (
	user_memory_cache "movepilot/pkg/cache/user_memory"
	"movepilot/pkg/constants"
	user_repo "movepilot/pkg/repositories/user"
	"movepilot/pkg/validate"

	"context"
	"movepilot/pkg/jwt"
	"movepilot/pkg/output"
	"net/http"
)

func AuthCachedMiddleware(repo user_repo.Repository, cache *user_memory_cache.Cache) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(constants.AUTH_TOKEN_HEADER)
			if len(token) == 0 {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token required"})
				return
			}

			parsed, err := jwt.Parse(token)
			if err != nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: err.Error()})
				return
			}

			id, ok := parsed["uuid"].(string)
			if !ok || !validate.ValidateUUID(id) {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token invalid"})
				return
			}

			if usr := cache.Get(id); usr != nil {
				ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			usr, err := repo.GetByUUID(r.Context(), id)
			if err != nil || usr == nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth failed"})
				return
			}

			cache.Set(id, usr)

			ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthAlwaysFreshMiddleware(repo user_repo.Repository, cache *user_memory_cache.Cache) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(constants.AUTH_TOKEN_HEADER)
			if len(token) == 0 {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token required"})
				return
			}

			parsed, err := jwt.Parse(token)
			if err != nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: err.Error()})
				return
			}

			id, ok := parsed["uuid"].(string)
			if !ok || !validate.ValidateUUID(id) {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token invalid"})
				return
			}

			usr, err := repo.GetByUUID(r.Context(), id)
			if err != nil || usr == nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth failed"})
				return
			}

			cache.Set(id, usr)

			ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
