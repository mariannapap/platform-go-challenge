package main

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "platform-go-challenge/docs" // swag init generates this file
	"platform-go-challenge/handlers"
)

// @title GWI Challenge API
// @version 1.0
// @description This is a sample server for the GWI Challenge API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email youremail@domain.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	router := mux.NewRouter()

	userFavouritesService := handlers.NewUserFavouritesService()

	router.HandleFunc("/users/{id}/favourites", userFavouritesService.GetUserFavourites).Methods("GET")
	router.HandleFunc("/users/{id}/favourites", userFavouritesService.AddAssetToFavourites).Methods("POST")
	router.HandleFunc("/users/{id}/favourites/{asset_id}", userFavouritesService.RemoveAssetFromFavourites).Methods("DELETE")
	router.HandleFunc("/users/{id}/favourites/{asset_id}", userFavouritesService.EditAssetDescription).Methods("PUT")

	// Add the Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", router)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return
	}
}
