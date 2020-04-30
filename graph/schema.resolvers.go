package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-api/db"
	"graphql-api/graph/generated"
	"graphql-api/models"
)

func (r *mutationResolver) NewPost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented 1"))
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic(fmt.Errorf("not implemented 2"))
}

func (r *postResolver) Tags(ctx context.Context, obj *models.Post) ([]*models.Tag, error) {
	panic(fmt.Errorf("not implemented 3"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	//fmt.Println("@@@@@@@@@@@@@@@@@@")
	//fmt.Println(ctx)
	//return []*models.Post{
	//	{
	//		ID:          1,
	//		Content:     "some",
	//		Description: "DEC",
	//		Name:        "Name",
	//	},
	//}, nil
	panic(fmt.Errorf("not implemented yet"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	db, err := db.Connection()

	if err != nil {
		return nil, err
	}

	var users []*models.User
	db.Find(&users)

	return users, nil
}

func (r *tagResolver) Posts(ctx context.Context, obj *models.Tag) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented 4 "))
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	fmt.Println(obj)
	db, err := db.Connection()

	if err != nil {
		return nil, err
	}

	var postsByUser []*models.Post

	db.Where("user_id = ?", obj.ID).Find(&postsByUser)

	return postsByUser, nil
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
