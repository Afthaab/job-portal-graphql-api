package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/afthaab/job-portal-graphql/graph/model"
)

// CreateSignup is the resolver for the createSignup field.
func (r *mutationResolver) CreateSignup(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Svc.UserSignup(input)
}

// CreateNewCompany is the resolver for the createNewCompany field.
func (r *mutationResolver) CreateNewCompany(ctx context.Context, input model.NewComapany) (*model.Company, error) {
	return r.Svc.CreateCompany(input)
}

// ViewAllCompanies is the resolver for the viewAllCompanies field.
func (r *queryResolver) ViewAllCompanies(ctx context.Context) ([]*model.Company, error) {
	return r.Svc.ViewAllCompanies()
}

// ViewCompanyByID is the resolver for the viewCompanyById field.
func (r *queryResolver) ViewCompanyByID(ctx context.Context, cid string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: ViewCompanyByID - viewCompanyById"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
