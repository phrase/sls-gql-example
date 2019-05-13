package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

func main() {
	schema, err := parseSchema()
	if err != nil {
		logrus.Fatal(err)
	}
	lambda.Start(gqlHandler(schema))
}
