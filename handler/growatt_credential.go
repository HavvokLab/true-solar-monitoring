package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type GrowattCredentialHandler struct {
	serv service.GrowattCredentialService
}

func NewGrowattCredentialHandler(serv service.GrowattCredentialService) *GrowattCredentialHandler {
	return &GrowattCredentialHandler{serv: serv}
}

func (h *GrowattCredentialHandler) FindAll(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.serv.FindAll(utx)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *GrowattCredentialHandler) Create(utx *domain.UserContext, c *fiber.Ctx) error {
	var req domain.CreateGrowattCredentialRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	err := h.serv.Create(utx, &req)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseCreated(c, nil)
}

func (h *GrowattCredentialHandler) Update(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	var req domain.UpdateGrowattCredentialRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	err = h.serv.Update(utx, int64(id), &req)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseCreated(c, nil)
}

func (h *GrowattCredentialHandler) Delete(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	err = h.serv.Delete(utx, int64(id))
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}
