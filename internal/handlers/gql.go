package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	gql "github.com/linkc0829/go-ics/internal/graph/generated"
	"github.com/linkc0829/go-ics/internal/graph/resolvers"
	"github.com/linkc0829/go-ics/internal/mongodb"

	"log"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(db *mongodb.MongoDB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			DB: db,
		},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	log.Println("conntect to grathql server @ " + path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}