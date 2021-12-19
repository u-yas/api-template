package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-template/graph/generated"
	"api-template/graph/model"
	"context"
	"fmt"
)

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

func (r *queryResolver) ListFavoriteUser(ctx context.Context) ([]*model.PublicUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListSearchUser(ctx context.Context, name string) ([]*model.PublicUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetSpecificUser(ctx context.Context, id string) (*model.PublicUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMe(ctx context.Context) (*model.PrivateUser, error) {
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
