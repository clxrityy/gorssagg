package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/clxrityy/gorssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := paramaters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
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
		respondWithError(w, 400, fmt.Sprintf("Couldn't follow feed: %v", err))
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}