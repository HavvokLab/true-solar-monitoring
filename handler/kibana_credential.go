package handler

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

type KibanaCredentialHandler struct {
	serv service.KibanaCredentialService
}

func NewKibanaCredentialHandler(serv service.KibanaCredentialService) *KibanaCredentialHandler {
	return &KibanaCredentialHandler{serv: serv}
}

func (h *KibanaCredentialHandler) FindOne(utx *domain.UserContext, c *fiber.Ctx) error {
	result, err := h.serv.FindOne(utx)
	if err != nil {
		return util.ResponseError(c, err)
	}

	return util.ResponseOK(c, result)
}
