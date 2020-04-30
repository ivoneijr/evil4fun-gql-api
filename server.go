package main

import (
	"graphql-api/db"
	"graphql-api/graph"
	"graphql-api/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db.Ping()
	db.AutoMigrate()
	//db.Seed()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//UsersRepo: database.UsersRepo{DB: DB}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
