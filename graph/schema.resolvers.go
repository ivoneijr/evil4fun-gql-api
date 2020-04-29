package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-api/graph/generated"
	"graphql-api/models"
)

func (r *mutationResolver) NewPost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Tags(ctx context.Context, obj *models.Post) ([]*models.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return []*models.Post{
		{
			ID:          "1",
			Content:     "some",
			Description: "DEC",
			Name:        "Name",
			CreatedAt:   "121",
			UpdatedAt:   "121",
		},
	}, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UsersRepo.GetUsers()
}

func (r *tagResolver) Posts(ctx context.Context, obj *models.Tag) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
