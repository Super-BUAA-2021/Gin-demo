BINARY_NAME=Gin-demo
# Github-action中指定go路径并使用更稳定
GO=/usr/local/go/bin/go

build:
	GOARCH=amd64 GOOS=linux ${GO} build -o ${BINARY_NAME}-linux main.go

# 根据文件名查找相应进程并kill
kill:
	ps -ef | grep ${BINARY_NAME} | grep -v grep | awk '{print $2}' | xargs kill -9

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

# 更新包
dep:
	${GO} mod tidy

vet:
	${GO} vet

lint:
	golangci-lint run --enable-all