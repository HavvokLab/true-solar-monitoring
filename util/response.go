package util

import (
	"github.com/HavvokLab/true-solar-monitoring/errors"
	"github.com/gofiber/fiber/v2"
)

type ServerResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewResponse(c *fiber.Ctx, code int, success bool, message string, result interface{}) error {
	res := new(ServerResponse)
	res.Message = message
	res.Result = result
	res.Success = success

	return c.Status(code).JSON(res)
}

func ResponseOK(c *fiber.Ctx, result interface{}) error {
	return NewResponse(c, fiber.StatusOK, true, "ok", result)
}

func ResponseCreated(c *fiber.Ctx, result interface{}) error {
	return NewResponse(c, fiber.StatusOK, true, "created", result)
}

func ResponseBadRequest(c *fiber.Ctx) error {
	return NewResponse(c, fiber.StatusBadRequest, false, "bad request", nil)
}

func ResponseError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errors.ServerError:
		return NewResponse(c, e.Code, false, e.Message, nil)
	default:
		return NewResponse(c, fiber.StatusInternalServerError, false, "system error", nil)
	}
}
