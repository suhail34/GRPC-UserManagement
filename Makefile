client-run:
	go run client/cmd/user-mgmt/main.go
.PHONY: client-run

server-run:
	go run server/main.go
.PHONY: server-run

build-docker-img:
	docker build -t rest-client:latest -f Dockerfile-client .; docker build -t grpc-server:latest -f Dockerfile-server .
.PHONY: build-docker-img

run-docker-img:
	docker run grpc-server:latest; docker run rest-client:latest
.PHONY: run-docker-img
