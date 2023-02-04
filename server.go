package main

import (
	"log"
	"myapp/config"
	"myapp/directives"
	"myapp/graph"
	"myapp/middleware"
	"myapp/migration"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	//add migration
	migration.MigrateTable()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//Defer closing Database
	db := config.GetDb()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
