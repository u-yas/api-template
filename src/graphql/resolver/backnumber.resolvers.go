package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-api-sample/src/graphql/generated"
	"go-api-sample/src/graphql/model"
)

func (r *mutationResolver) CreateBacknumber(ctx context.Context, input model.CreateBacknumberInput) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBacknumber(ctx context.Context, input model.UpdateBacknumberInput) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBacknumber(ctx context.Context, input model.DeleteBacknumberInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListTopPageBacknumbers(ctx context.Context) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListMyAllBacknumbers(ctx context.Context) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListSuggestBacknumbers(ctx context.Context) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListBacknumbersFromWorksID(ctx context.Context, worksID string) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListFilterTagsBacknumbers(ctx context.Context, tag string) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListSearchBacknumbers(ctx context.Context, title string) ([]*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMySpecificBacknumbers(ctx context.Context, backnumberID string) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPublicSpecificBacknumber(ctx context.Context, backnumberID string) (*model.Backnumber, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
