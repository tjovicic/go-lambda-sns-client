build: format
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go sns.go

test:
	go test ./...

format:
	go fmt ./...

clean:
	rm -rf ./bin

deploy: clean format build
	sls deploy --verbose

deploy-prod: build
	sls deploy -s prod --verbose

deploy-dev: build
	sls deploy -s dev --verbose

remove:
	sls remove
