package main

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	server "github.com/suhail34/user-mgmt/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
  port = ":50501"
)

type UserManagementServer struct {
  server.UnimplementedUserManagementServer
}

func GetAllUsers() ([]*server.UserByIdResponse, error) {
  dir, err := os.Getwd()
  if err != nil {
    return nil, err
  }

  filePath := filepath.Join(dir, "server/users.json")
  file, err := os.Open(filePath)
  if err!=nil {
    return nil, err
  }
  defer file.Close()

  var users []*server.UserByIdResponse
  err = json.NewDecoder(file).Decode(&users)
  if err!=nil {
    return nil, err
  }

  return users, nil
}

func (s *UserManagementServer) GetUserById(ctx context.Context, in *server.UserByIdRequest) (*server.UserByIdResponse, error) {
  users, err := GetAllUsers()
  if err != nil {
    logrus.Error(err)
    return nil, err
  }
  var user server.UserByIdResponse
  for _, u := range users {
    if u.UserId==in.UserId {
      user = *u
    }
  }

  if user.FirstName=="" {
    err = errors.New("user not found with the specified Id")
    logrus.Error(err)
    return nil, err
  }
  logrus.Infof("Successfully Retrived user with id : %v", in.UserId)

  return &user, nil
}

func (s *UserManagementServer) GetUsersByIds(ctx context.Context, in *server.UsersByIdsRequest) (*server.UsersByIdsResponse, error) {
  users, err := GetAllUsers()
  if err != nil {
    logrus.Error(err)
    return nil, err
  }

  usersMap := make(map[int]server.UserByIdResponse)
  for _, user := range users {
    usersMap[int(user.UserId)] = *user
  }

  var foundUsers server.UsersByIdsResponse
  ids := in.UsersByIdsRequest

  for _, id := range ids {
    if user, ok := usersMap[int(id)]; ok {
      foundUsers.Users = append(foundUsers.Users, &user)
    }
  }
  logrus.Infof("Retrived all users with ids : %v", ids)

  return &foundUsers, nil
}

func main() {
  listener, err := net.Listen("tcp", port)
  if err!=nil {
    logrus.Error(err)
  }

  srv := grpc.NewServer()
  server.RegisterUserManagementServer(srv, &UserManagementServer{})
  reflection.Register(srv)

  logrus.Infof("Grpc Server listening on port %v", port)
  if err := srv.Serve(listener); err != nil {
    logrus.Error(err)
  }
}
