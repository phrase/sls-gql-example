package main

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
	Job       string                 `json:"job"`
}

func gqlHandler(schema *graphql.Schema) func(ctx context.Context, pl *payload) (interface{}, error) {
	return func(ctx context.Context, pl *payload) (interface{}, error) {
		start := time.Now()
		c := newContext(ctx, logrus.Fields{"job": pl.Job, "query": pl.Query})
		if pl.Query == "" {
			return nil, errors.Errorf("query must be present")
		}
		c.Printf("executing")
		ctx = context.WithValue(ctx, contextKey, c)
		res := schema.Exec(ctx, pl.Query, "", pl.Variables)
		if len(res.Errors) > 0 {
			c.WithField("errors", res.Errors).Error("failed")
		}
		c.WithField("total_time", time.Since(start).Seconds()).Print("finished")
		return res, nil
	}
}
