package routes

import (

	"github.com/gofiber/fiber/v2"
	"github.com/suhail34/user-mgmt/client/internal/handlers"
	rpc_server "github.com/suhail34/user-mgmt/server/proto"
)

func SetupRoutes(app *fiber.App, client rpc_server.UserManagementClient) {
  app.Get("/user/:id", handlers.GetUserById(client))
  app.Post("/users", handlers.GetUsersByIds(client))
}
