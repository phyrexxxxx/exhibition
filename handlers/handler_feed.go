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

func (apiCfg *HandlerApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// 400 Bad Request
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	// 201 Status Created
	utils.RespondWithJSON(w, http.StatusCreated, models.DBFeedToFeed(feed))
}

func (apiCfg *HandlerApiConfig) HandlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Cannot get all feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DBFeedsToFeeds(feeds))
}
