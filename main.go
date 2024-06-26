package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/clxrityy/gorssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!");

	godotenv.Load(".env");

	portString := os.Getenv("PORT");

	if (portString == "") {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL");

	if (dbURL == "") {
		log.Fatal("DB_URL environment variable not set")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
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

	v1Router.Post("/users", apiCfg.createUserHandler);
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.getUserHandler));

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.createFeedHandler));
	v1Router.Get("/feeds", apiCfg.getFeedsHandler);

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.createFeedFollowsHandler));
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.getFeedFollowsHandler));
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.deleteFeedFollowHandler));

	router.Mount("/v1", v1Router);

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Println("Server is running on port: ", portString);

	err = srv.ListenAndServe();

	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("PORT: ", portString)
}