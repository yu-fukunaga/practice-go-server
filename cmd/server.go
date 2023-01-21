package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"practice-server/graph"
	"practice-server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
)

const defaultPort = "8080"

var sampleCmd = &cobra.Command{
	Use:   "api-server",
	Short: "API server start",
	Long:  "API server start",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("SAMPLE_PORT")
		fmt.Println("port",port)
		if port == "" {
			port = defaultPort
		}

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	},
}
