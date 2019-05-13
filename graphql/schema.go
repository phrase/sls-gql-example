//go:generate go-bindata -pkg main -nometadata -ignore ".go" .
package main

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func parseSchema() (*graphql.Schema, error) {
	b, err := schemaGqlBytes()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return graphql.ParseSchema(string(b), &resolver{})
}

type resolver struct {
}

func (r *resolver) Hello(ctx context.Context, input struct {
	Msg *string
}) (string, error) {
	if input.Msg != nil {
		return *input.Msg, nil
	}
	return "world", nil
}

func (r *resolver) CreateUser(ctx context.Context, input struct {
	Name  string
	Age   int32
	Tags  *[]string
	Email *string
}) (graphql.ID, error) {
	id := uuid.NewV1().String()
	return graphql.ID(id), nil
}
