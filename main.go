package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaketran42/go-graphql/gql"
	"github.com/jaketran42/go-graphql/postgres"
	"github.com/jaketran42/go-graphql/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	// Initialize our api
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	// Create a new connection to database
	db, err := postgres.New(
		postgres.ConStr("localhost", 5432, "jakeTran", "go_graphql_db"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create root query for graphql
	rootQuery := gql.NewRoot(db)
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	// Middlewares
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, 
		middleware.StripSlashes,    
		middleware.Recoverer, 
	)

	router.Post("/graphql", s.GraphQL())

	return router, db
}