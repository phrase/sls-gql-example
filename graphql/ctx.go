package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/sirupsen/logrus"
)

var contextKey = struct{}{}

func newContext(ctx context.Context, fields logrus.Fields) *Context {
	raw := logrus.New()
	raw.Formatter = &logrus.JSONFormatter{}
	c := &Context{FieldLogger: raw}
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		c.FieldLogger = c.FieldLogger.WithField("request_id", lc.AwsRequestID)
		c.RequestID = lc.AwsRequestID
	}
	if fields != nil {
		for k, v := range fields {
			if v != "" {
				c.FieldLogger = c.FieldLogger.WithField(k, v)
			}
		}
	}
	return c
}

func fromContext(ctx context.Context) *Context {
	c, ok := ctx.Value(contextKey).(*Context)
	if !ok {
		return newContext(ctx, nil)
	}
	return c
}

type Context struct {
	logrus.FieldLogger
	RequestID string
}
