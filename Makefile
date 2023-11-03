client-run:
	go run client/cmd/user-mgmt/main.go
.PHONY: client-run

server-run:
	go run server/main.go
.PHONY: server-run

