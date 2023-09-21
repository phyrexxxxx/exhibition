package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/phyrexxxxx/exhibition/auth"
	"github.com/phyrexxxxx/exhibition/config"
	"github.com/phyrexxxxx/exhibition/internal/database"
	"github.com/phyrexxxxx/exhibition/utils"
)

// use Embedding to prevent "cannot define new methods on non-local type config.ApiConfig" error
type HandlerApiConfig struct {
	*config.ApiConfig
}

func (apiCfg *HandlerApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// The json:"name" tag is used to specify how the field should be serialized when converting the struct to JSON
	type parameters struct {
		Name string `json:"name"`
	}

	// creates a new JSON decoder that reads from the request body
	decoder := json.NewDecoder(r.Body)

	// decoding JSON data into a Go struct called `params`
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// 400 Bad Request
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	// 201 Status Created
	utils.RespondWithJSON(w, http.StatusCreated, config.DBUserToUser(user))
}

func (apiCfg *HandlerApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
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

	utils.RespondWithJSON(w, http.StatusOK, config.DBUserToUser(user))
}
