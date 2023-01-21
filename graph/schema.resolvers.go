package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"practice-server/ent"
	"practice-server/graph/generated"
	"practice-server/graph/model"
	"time"

	"entgo.io/ent/dialect"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// TODO: userService的なものをを作る
	// TODO: 既存ユーザーのチェックをする
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	passwordHash := string(hashed)
	
	client, err := ent.Open(dialect.MySQL, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	user, err := client.User.Create().SetUserName(input.Name).SetUserEmail(input.Email).SetPassHash(passwordHash).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID: user.ID.String(),
		Name: user.UserName,
		Email: user.UserEmail,
	}, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%s", time.Now()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
