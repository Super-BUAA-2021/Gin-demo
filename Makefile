BINARY_NAME=Gin-demo
GO=/usr/local/go/bin/go

build:
	GOARCH=amd64 GOOS=linux ${GO} build -o ${BINARY_NAME}-linux main.go

run:
	./${BINARY_NAME}-linux

build_and_run: build run

clean:
	${GO} clean
	rm ${BINARY_NAME}-linux


test:
	${GO} test ./...

test_coverage:
	${GO} test ./... -coverprofile=coverage.out

dep:
	${GO} mod tidy

vet:
	${GO} vet

lint:
	golangci-lint run --enable-all