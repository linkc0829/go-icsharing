package resolvers

import (
	"context"
	"github.com/linkc0829/go-ics/internal/graph/generated"
)

//Resolver contains db element
type Resolver struct{
	db mongo.MongoDB
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
