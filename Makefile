.PHONY: build clean deploy gomodgen schemagen

schemagen:
	go generate ./...

build: schemagen gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/graphql ./graphql

clean:
	rm -rf ./bin ./vendor Gopkg.lock

test:
	go test ./...

deploy: clean test build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
