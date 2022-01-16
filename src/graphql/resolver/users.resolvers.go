package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-api-sample/src/db"
	"go-api-sample/src/graphql/model"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.PrivateUser, error) {
	now := time.Now() // ①
	uid := uuid.NewString()

	dbC := make(chan bool)
	esC := make(chan bool)

	dbInput := func() {
		user := db.User{
			UUID:    uid,
			Email:   input.Email,
			Name:    input.Email,
			Suspend: false,
		}
		err := user.Insert(ctx, r.DbCon, boil.Infer())
		if err != nil {
			fmt.Println(err)
			dbC <- true
		}
		dbC <- false
	}
	go dbInput()
	esInput := func() {
		_, er := r.Es.Index("user", strings.NewReader(`{"email":"`+input.Email+`"}`), r.Es.Index.WithDocumentID(uid))
		if er != nil {
			fmt.Println(er)
			esC <- true
		}
		esC <- false
	}
	go esInput()

	// recieve chan value from db and es

	err := <-dbC && <-esC
	if err {
		return nil, fmt.Errorf("failed to create user")
	}

	result := &model.PrivateUser{
		ID:    uid,
		Name:  input.Email,
		Email: input.Email,
	}
	end := time.Now()

	fmt.Println(end.Sub(now))
	return result, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.PrivateUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMe(ctx context.Context) (bool, error) {
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
