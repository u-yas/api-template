package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-template/src/db"
	"api-template/src/graphql/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.PrivateUser, error) {
	now := time.Now() // ①
	uid := uuid.NewString()
	user := db.User{
		UUID:    uid,
		Email:   input.Email,
		Name:    input.Email,
		Suspend: false,
	}
	err := user.Insert(ctx, r.DbCon, boil.Infer())
	if err != nil {
		return nil, err
	}

	for i, title := range []string{"Test One", "Test Two"} {
		r.Wg.Add(1)

		go func(i int, title string) {
			defer r.Wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), r.Es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	r.Wg.Wait()
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
