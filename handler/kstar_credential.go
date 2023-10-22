package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type KStarCredentialHandler struct {
	serv service.KStarCredentialService
}

func NewKStarCredentialHandler(serv service.KStarCredentialService) *KStarCredentialHandler {
	return &KStarCredentialHandler{serv: serv}
}

func (h *KStarCredentialHandler) FindAll(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.serv.FindAll(utx)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *KStarCredentialHandler) Create(utx *domain.UserContext, c *fiber.Ctx) error {
	var req domain.CreateKStarCredentialRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	err := h.serv.Create(utx, &req)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseCreated(c, nil)
}

func (h *KStarCredentialHandler) Update(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	var req domain.UpdateKStarCredentialRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	err = h.serv.Update(utx, int64(id), &req)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseCreated(c, nil)
}

func (h *KStarCredentialHandler) Delete(utx *domain.UserContext, c *fiber.Ctx) error {
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
