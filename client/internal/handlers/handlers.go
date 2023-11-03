package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	rpc_server "github.com/suhail34/user-mgmt/server/proto"
)

func GetUserById(client rpc_server.UserManagementClient) func(c *fiber.Ctx) error {

  return func(c *fiber.Ctx) error {
    param := c.Params("id")
    id, err := strconv.ParseInt(param, 10, 64)
    if err != nil {
      logrus.Errorf("Invalid Id : %v", id)
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error":"Invalid Id",
      })
    }
    req := &rpc_server.UserByIdRequest{UserId: id}
    res, err := client.GetUserById(c.Context(), req)
    if err != nil {
      logrus.Errorf("User with id %v does not exist", id)
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error" : fmt.Sprintf("Error while getting user by that Id %v", id),
      })
    }
    logrus.Infof("Successfully retrieved user with id : %v", id)
    return c.JSON(res)
  }
}

func GetUsersByIds(client rpc_server.UserManagementClient) func(c *fiber.Ctx) error {
  return func(c *fiber.Ctx) error {
    var body *rpc_server.UsersByIdsRequest
    err := c.BodyParser(&body)
    if err != nil {
      logrus.Error("Invalid JSON body : ", string(c.Body()), err)
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error" : fmt.Sprintf("Error parsing the json body : %v %v", string(c.Body()), err),
      })
    }
    req := &rpc_server.UsersByIdsRequest{UsersByIdsRequest: body.UsersByIdsRequest}
    res, err := client.GetUsersByIds(c.Context(), req)
    if err != nil {
      logrus.Error("Users with this ids does not exits : ", body)
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error" : fmt.Sprintf("Error while getting users by ids %v", body),
      })
    }
    logrus.Infof("Successfully retrived users with ids : %v", body)
    return c.JSON(res)
  }
}
