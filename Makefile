

.PHONY: proto
proto:
	sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=xxxxxx cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/userApi/userApi.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o userApi-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t userApi-api:latest
