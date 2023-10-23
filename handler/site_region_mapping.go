package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type SiteRegionMappingHandler struct {
	siteRegionMappingService service.SiteRegionMappingService
}

func NewSiteRegionMappingHandler(siteRegionMappingService service.SiteRegionMappingService) *SiteRegionMappingHandler {
	return &SiteRegionMappingHandler{
		siteRegionMappingService: siteRegionMappingService,
	}
}

func (h *SiteRegionMappingHandler) GetCities(utx *domain.UserContext, c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 12)
	offset := c.QueryInt("offset", -1)
	result, err := h.siteRegionMappingService.FindAll(limit, offset)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *SiteRegionMappingHandler) GetRegions(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.siteRegionMappingService.FindRegion()
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *SiteRegionMappingHandler) CreateCity(utx *domain.UserContext, c *fiber.Ctx) error {
	req := domain.CreateCityRequest{}
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.siteRegionMappingService.CreateCity(&req); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseCreated(c, nil)
}

func (h *SiteRegionMappingHandler) UpdateRegion(utx *domain.UserContext, c *fiber.Ctx) error {
	var req domain.UpdateRegionRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.siteRegionMappingService.UpdateRegion(&req); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}

func (h *SiteRegionMappingHandler) UpdateCity(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	var req domain.UpdateCityRequest
	if err := c.BodyParser(&req); err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.siteRegionMappingService.UpdateCity(int64(id), &req); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}

func (h *SiteRegionMappingHandler) DeleteCity(utx *domain.UserContext, c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.siteRegionMappingService.DeleteCity(int64(id)); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}

func (h *SiteRegionMappingHandler) DeleteArea(utx *domain.UserContext, c *fiber.Ctx) error {
	area := c.Params("area", "")
	if util.EmptyString(area) {
		return util.ResponseBadRequest(c)
	}

	if err := h.siteRegionMappingService.DeleteArea(area); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}
