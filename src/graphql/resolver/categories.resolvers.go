package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-template/src/graphql/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, name *string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, name *string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddBacknumberCategory(ctx context.Context, backnumberID string, categoryID string) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveBacknumberCategory(ctx context.Context, backnumberID string, categoryID string) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategory(ctx context.Context, id string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListBacknumberByCategory(ctx context.Context, tagID string) ([]*model.Works, error) {
	panic(fmt.Errorf("not implemented"))
}
