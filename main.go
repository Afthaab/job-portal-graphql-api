package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/afthaab/job-portal-graphql/auth"
	"github.com/afthaab/job-portal-graphql/database"
	"github.com/afthaab/job-portal-graphql/graph"
	"github.com/afthaab/job-portal-graphql/middlewares"
	"github.com/afthaab/job-portal-graphql/repository"
	"github.com/afthaab/job-portal-graphql/service"
)

const defaultPort = "8080"

func main() {

	svc, _, err := StartApp()
	if err != nil {
		log.Info().Err(err).Msg("error in startapp")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Svc: svc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}

func StartApp() (service.UserService, middlewares.Mid, error) {

	// =========================================================================
	// initializing the authentication support
	log.Info().Msg("main started : initializing the authentication support")

	//reading the private key file
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in reading auth private key : %w", err) // %w is used for error wraping
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in parsing auth private key : %w", err) // %w is used for error wraping
	}
	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in reading auth public key : %w", err) // %w is used for error wraping
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in parsing auth public key : %w", err) // %w is used for error wraping
	}

	// initializing the authentication layer
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in constructing auth %w", err)
	}

	// initializing middleware layer
	mid, err := middlewares.NewMiddleware(a)
	if err != nil {
		return &service.Service{}, mid, fmt.Errorf("error in constructing middleware %w", err) // %w is used for error wraping
	}

	// starting database connection
	db, err := database.ConnectToDatabase()
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in opening the database connection : %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("database is not connected: %w", err)
	}
	repo, err := repository.NewRepository(db)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("repository not initialized: %w", err)
	}
	svc, err := service.NewService(repo)
	if err != nil {
		return &service.Service{}, middlewares.Mid{}, fmt.Errorf("service layer not initialized: %w", err)
	}

	return svc, mid, nil
}
