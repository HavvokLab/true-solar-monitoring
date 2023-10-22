package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type InstalledCapacityHandler struct {
	serv service.InstalledCapacityService
}

func NewInstalledCapacityHandler(serv service.InstalledCapacityService) *InstalledCapacityHandler {
	return &InstalledCapacityHandler{serv: serv}
}

func (h *InstalledCapacityHandler) FindOne(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.serv.FindOne(utx)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *InstalledCapacityHandler) Update(utx *domain.UserContext, c *fiber.Ctx) error {
	var request domain.UpdateInstalledCapacityRequest
	if err := c.BodyParser(&request); err != nil {
		return util.ResponseBadRequest(c)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.serv.Update(utx, int64(id), &request); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}
