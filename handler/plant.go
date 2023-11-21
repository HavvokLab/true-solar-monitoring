package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type plantHandler struct {
	serv service.PlantService
}

func NewPlantHandler(serv service.PlantService) *plantHandler {
	return &plantHandler{serv: serv}
}

func (h *plantHandler) FindAll(utx *domain.UserContext, c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 12)
	offset := c.QueryInt("offset", -1)
	result, err := h.serv.FindAllWithPagination(offset, limit)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *plantHandler) Delete(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.serv.Delete(int64(id)); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}
