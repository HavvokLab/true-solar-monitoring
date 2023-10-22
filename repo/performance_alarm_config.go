package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"gorm.io/gorm"
)

type PerformanceAlarmConfigRepo interface {
	GetLowPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error)
	UpdateLowPerformanceAlarmConfig(*model.PerformanceAlarmConfig) error
	GetSumPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error)
	GetDailyPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error)
}

type performanceAlarmConfigRepo struct {
	db *gorm.DB
}

func NewPerformanceAlarmConfigRepo(db *gorm.DB) PerformanceAlarmConfigRepo {
	return &performanceAlarmConfigRepo{
		db: db,
	}
}

func (r *performanceAlarmConfigRepo) GetLowPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error) {
	tx := r.db.Session(&gorm.Session{})
	data := model.PerformanceAlarmConfig{}
	if err := tx.Find(&data, "name = ?", constant.LOW_PERFORMANCE_ALARM).Error; err != nil {
		return nil, util.TranslateSqliteError(err)
	}

	return &data, nil
}

func (r *performanceAlarmConfigRepo) UpdateLowPerformanceAlarmConfig(data *model.PerformanceAlarmConfig) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Model(&data).Where("name = ?", constant.LOW_PERFORMANCE_ALARM).Updates(data).Error; err != nil {
		return util.TranslateSqliteError(err)
	}

	return nil
}

func (r *performanceAlarmConfigRepo) GetSumPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error) {
	tx := r.db.Session(&gorm.Session{})
	data := model.PerformanceAlarmConfig{}
	if err := tx.Find(&data, "name = ?", constant.SUM_PERFORMANCE_ALARM).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *performanceAlarmConfigRepo) GetDailyPerformanceAlarmConfig() (*model.PerformanceAlarmConfig, error) {
	tx := r.db.Session(&gorm.Session{})
	data := model.PerformanceAlarmConfig{}
	if err := tx.Find(&data, "name = ?", constant.DAILY_PERFORMANCE_ALARM).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
