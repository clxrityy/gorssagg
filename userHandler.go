package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/clxrityy/gorssagg/internal/database"
	"github.com/google/uuid"
);

func (apiCfg *apiConfig)createUserHandler(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:name`
	}

	params := paramaters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig)getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseUserToUser(user))
}