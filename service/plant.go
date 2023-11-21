package service

import (
	"encoding/csv"
	"os"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

type PlantService interface {
	BulkCreate([]*model.Plant) error
	ExportToCsv() error
	FindAllWithPagination(offset, limit int) (*domain.FindAllPlantResponse, error)
	Delete(id int64) error
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
	err := s.repo.BatchCreate(plants)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *plantService) ExportToCsv() error {
	plants, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error(err)
		return err
	}

	file, err := os.Create(constant.PLANT_REPORT_FILE)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(constant.PLANT_REPORT_HEADER)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	for _, plant := range plants {
		err = writer.Write(plant.CsvRow())
		if err != nil {
			s.logger.Error(err)
			return err
		}
	}

	return nil
}

func (s *plantService) FindAllWithPagination(offset, limit int) (*domain.FindAllPlantResponse, error) {
	if offset < 0 {
		offset = 0
	}

	if limit < 1 {
		limit = 12
	}

	plants, err := s.repo.FindAllWithPagination(offset, limit)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count()
	if err != nil {
		return nil, err
	}

	return &domain.FindAllPlantResponse{Plants: plants, Count: count}, nil
}

func (s *plantService) Delete(id int64) error {
	return s.repo.Delete(id)
}
