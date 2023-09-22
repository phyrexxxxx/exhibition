package auth

import (
	"fmt"
	"net/http"

	"github.com/phyrexxxxx/exhibition/internal/database"
	"github.com/phyrexxxxx/exhibition/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *AuthApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetApiKey(r.Header)
		if err != nil {
			// 401 Status Unauthorized
			utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			// 400 Status Bad Request
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Cannot Get User: %v", err))
			return
		}

		handler(w, r, user)
	}
}
