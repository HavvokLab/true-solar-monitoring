package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	serv service.AuthService
}

func NewAuthHandler(serv service.AuthService) *AuthHandler {
	return &AuthHandler{
		serv: serv,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := domain.LoginRequest{}
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	res, err := h.serv.Login(&req)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, res)
}
