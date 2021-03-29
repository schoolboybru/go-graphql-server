package main

import (
	"github.com/graphql-go/handler"
	"github.com/schoolboybru/go-graphql-server/graph"
	"log"
	"net/http"
)

func main() {

	h := handler.New(&handler.Config{
		Schema:       &graph.MainSchema,
		Pretty:       true,
		GraphiQL:     true,
		Playground:   true,
	})


	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
