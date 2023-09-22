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

func (apiCfg *HandlerApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	// 201 Status Created
	utils.RespondWithJSON(w, http.StatusCreated, models.DBFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *HandlerApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DBFeedFollowsToFeedFollows(feedFollows))
}
