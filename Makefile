check: lint test

fmt:
	go fmt ./...

fmtfix:
	golangci-lint run --fix -E gofmt,gofumpt,goimports

lint:
	golangci-lint run

test:
	go test ./...

run:
	go run main.go

mock:
	mockery --all --keeptree

remock:
	rm -rf ./mocks
	mockery --all --keeptree

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	DEBUG=1 swagger generate spec -o ./swagger.json -m

swagger-validate:
	swagger validate

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.json