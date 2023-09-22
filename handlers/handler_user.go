package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/phyrexxxxx/exhibition/internal/database"
	"github.com/phyrexxxxx/exhibition/models"
	"github.com/phyrexxxxx/exhibition/utils"
)

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
	utils.RespondWithJSON(w, http.StatusCreated, models.DBUserToUser(user))
}

func (apiCfg *HandlerApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.DBUserToUser(user))
}
