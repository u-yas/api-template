package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-template/src/graphql/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateWorks(ctx context.Context, works *model.CreateWorksInput) (*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWorks(ctx context.Context, works *model.UpdateWorksInput) (*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWorks(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetWorks(ctx context.Context, id string) (*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMyWorks(ctx context.Context, id string) (*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListSearchWorks(ctx context.Context, search string) ([]*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListMyWorks(ctx context.Context) ([]*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListRankingWorks(ctx context.Context) ([]*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListSuggestWorks(ctx context.Context) ([]*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}
