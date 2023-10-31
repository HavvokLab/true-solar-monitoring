package service

import (
	"encoding/csv"
	"os"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

type PlantService interface {
	BulkCreate([]*model.Plant) error
	ExportToCsv() error
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
