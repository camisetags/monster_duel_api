package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	
	"monster_duel_api/generated"
	"monster_duel_api/resolvers"
)

const defaultPort = "8080"

// Server struct that contains server logic
type Server struct {}

// Run runs the server
func (s Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	schema := generated.NewExecutableSchema(generated.Config{
		// Resolvers: &resolvers.Resolver{},
		Resolvers: &resolvers.Resolver{},
	})

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(schema))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 
