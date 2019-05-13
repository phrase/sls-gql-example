# sls-gql

Example to use GraphQL as an API for AWS Lambda.

## Setup

* Install go-bindata from https://github.com/go-bindata/go-bindata
* Install serverless framework: `yarn global add yarn` or `npm install -g serverless`
* make sure you have AWS credentials (e.g. `AWS_ACCESS_KEY_ID`,
	`AWS_SECRET_ACCESS_KEY`) in your ENV.
* `make deploy`

## Schema Updates

The schema can be found at `graphql/schema.gql` but in the code the inlined
version is used (via go-bindata). When you update the schema you also need to
either run `make schemagen` or `go generate ./...`.

## Execute query via invoke

### Without variables
	sls invoke -f graphql  --data '{"query":"query { hello } "}'

### With variables

```
sls invoke -f graphql <<EOF
{
  "query": "query q(\$msg:String!) { hello(msg:\$msg) }",
  "variables": {
    "msg": "World"
  }
}
EOF
```

## Teardown

* `sls remove`
