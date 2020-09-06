#!make

test: test-config test-log test-errors test-mysql

test-config:
	cd ./config && \
	go test -timeout 300s -cover -a -v

test-log:
	cd ./log && \
	go test -timeout 300s -cover -a -v

test-errors:
	cd ./errors && \
	go test -timeout 300s -cover -a -v

test-mysql:
	cd ./libs/mysql && \
	go test -timeout 300s -cover -a -v

start:
	go mod vendor && \
	go run main.go
