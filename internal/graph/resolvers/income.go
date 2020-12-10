package resolvers

import (
	"context"

	"github.com/linkc0829/go-ics/internal/graph/generated"
	"github.com/linkc0829/go-ics/internal/graph/models"
)


func (r *mutationResolver) CreateIncome(ctx context.Context, input models.IncomeInput) (*models.Income, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateIncome(ctx context.Context, id string, input models.IncomeInput) (*models.Income, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteIncome(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) LikeIncome(ctx context.Context, id string) (int, error) {
	panic("not implemented")
}