package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	
	"github.com/camisetags/monster_duel_api_go/generated"
	"github.com/camisetags/monster_duel_api_go/resolvers"
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

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 
