package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suhail34/user-mgmt/client/internal/routes"
  rpc_server "github.com/suhail34/user-mgmt/server/proto"
)

func NewServer(client rpc_server.UserManagementClient) *fiber.App {
  app := fiber.New()
  routes.SetupRoutes(app, client)
  return app
}
