package domain

import "github.com/HavvokLab/true-solar-monitoring/model"

type FindAllPlantResponse struct {
	Plants []*model.Plant `json:"plants"`
	Count  int64          `json:"count"`
}
