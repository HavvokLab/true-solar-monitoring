package service

import (
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

type PlantService interface {
	BulkCreate([]*model.Plant) error
}

type plantService struct {
	repo   repo.PlantRepo
	logger logger.Logger
}

func NewPlantService(repo repo.PlantRepo, logger logger.Logger) PlantService {
	return &plantService{
		repo:   repo,
		logger: logger,
	}
}

func (s *plantService) BulkCreate(plants []*model.Plant) error {
	perBatch := 100
	batches := make([][]*model.Plant, 0)
	batch := make([]*model.Plant, 0)

	for i, plant := range plants {
		if (i+1)%perBatch == 0 {
			batches = append(batches, batch)
			batch = make([]*model.Plant, 0)
		}

		batch = append(batch, plant)
	}
	batches = append(batches, batch)

	for _, batch := range batches {
		if err := s.repo.BulkCreate(batch); err != nil {
			s.logger.Error(err)
			return err
		}
	}

	return nil
}
