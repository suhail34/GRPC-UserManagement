package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suhail34/user-mgmt/client/internal/server"
	rpc_server "github.com/suhail34/user-mgmt/server/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50501", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	client := rpc_server.NewUserManagementClient(conn)

	srv := server.NewServer(client)
	if err := srv.Listen(":8080"); err != nil {
		logrus.Error("Error starting the server : ", err)
		os.Exit(1)
	}
}
