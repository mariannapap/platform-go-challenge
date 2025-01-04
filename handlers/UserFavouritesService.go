package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"platform-go-challenge/models"
)

type UserFavouritesService struct {
	users map[string]*models.User
}

func NewUserFavouritesService() *UserFavouritesService {
	return &UserFavouritesService{
		users: make(map[string]*models.User),
	}
}

// GetUserFavourites godoc
// @Summary Get user's favourites
// @Description Get the list of user's favourite assets
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {array} models.Asset
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/favourites [get]
func (s *UserFavouritesService) GetUserFavourites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, ok := s.users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(user.Favourites)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// AddAssetToFavourites godoc
// @Summary Add asset to user's favourites
// @Description Add an asset to the user's list of favourite assets
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param asset body models.Asset true "Asset"
// @Success 200 {array} models.Asset
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/favourites [post]
func (s *UserFavouritesService) AddAssetToFavourites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var asset models.Asset
	err := json.NewDecoder(r.Body).Decode(&asset)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, ok := s.users[userID]
	if !ok {
		user = &models.User{ID: userID}
		s.users[userID] = user
	}

	user.Favourites = append(user.Favourites, asset)
	err = json.NewEncoder(w).Encode(user.Favourites)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// RemoveAssetFromFavourites godoc
// @Summary Remove asset from user's favourites
// @Description Remove an asset from the user's list of favourite assets
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param asset_id path string true "Asset ID"
// @Success 200 {array} models.Asset
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/favourites/{asset_id} [delete]
func (s *UserFavouritesService) RemoveAssetFromFavourites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	assetID := vars["asset_id"]

	user, ok := s.users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	for i, asset := range user.Favourites {
		if asset.ID == assetID {
			user.Favourites = append(user.Favourites[:i], user.Favourites[i+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(user.Favourites)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// EditAssetDescription godoc
// @Summary Edit asset description
// @Description Edit the description of an asset in the user's list of favourite assets
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param asset_id path string true "Asset ID"
// @Param asset body models.Asset true "Asset"
// @Success 200 {array} models.Asset
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/favourites/{asset_id} [put]
func (s *UserFavouritesService) EditAssetDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	assetID := vars["asset_id"]

	var updatedAsset models.Asset
	err := json.NewDecoder(r.Body).Decode(&updatedAsset)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, ok := s.users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	for i, asset := range user.Favourites {
		if asset.ID == assetID {
			user.Favourites[i].Description = updatedAsset.Description
			break
		}
	}

	err = json.NewEncoder(w).Encode(user.Favourites)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
