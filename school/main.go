package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"school/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	s, err := graph.NewGraphQLServer(goDotEnvVariable("STUDENT_SERVICE"),goDotEnvVariable("TEACHER_SERVICE"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.New(s.TOExecutableSchema()))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func goDotEnvVariable(key string) string {
	// load env file
	err := godotenv.Load("env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}
	return os.Getenv(key)
}