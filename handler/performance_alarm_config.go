package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type PerformanceAlarmConfigHandler struct {
	performanceAlarmConfigService service.PerformanceAlarmConfigService
}

func NewPerformanceAlarmConfigHandler(performanceAlarmConfigService service.PerformanceAlarmConfigService) *PerformanceAlarmConfigHandler {
	return &PerformanceAlarmConfigHandler{
		performanceAlarmConfigService: performanceAlarmConfigService,
	}
}

func (h *PerformanceAlarmConfigHandler) GetLowPerformanceAlarmConfig(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.performanceAlarmConfigService.GetLowPerformanceAlarmConfig(utx)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}

func (h *PerformanceAlarmConfigHandler) UpdateLowPerformanceAlarmConfig(utx *domain.UserContext, c *fiber.Ctx) error {
	request := domain.UpdatePerformanceAlarmConfigRequest{}
	if err := c.BodyParser(&request); err != nil {
		return util.ResponseBadRequest(c)
	}

	if err := h.performanceAlarmConfigService.UpdateLowPerformanceAlarmConfig(utx, &request); err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, nil)
}
