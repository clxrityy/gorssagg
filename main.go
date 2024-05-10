package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!");

	godotenv.Load(".env");

	portString := os.Getenv("PORT");

	if (portString == "") {
		log.Fatal("PORT environment variable not set")
	}

	router := chi.NewRouter();

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}));

	v1Router := chi.NewRouter();

	v1Router.Get("/healthz", readyHandler);
	v1Router.Get("/error", errorHandler);


	router.Mount("/v1", v1Router);

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Println("Server is running on port: ", portString);

	err := srv.ListenAndServe();

	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("PORT: ", portString)
}