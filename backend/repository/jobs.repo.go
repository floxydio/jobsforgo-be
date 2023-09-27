package repository

import (
	"jobsforgobe/database"
	"jobsforgobe/entities"

	"gorm.io/gorm"
)

type JobsSqlStruct struct {
	Db *gorm.DB
}

type JobsRepository interface {
	FindAllJob() ([]entities.Jobs, error)
	FindAllJobByCategory(string) ([]entities.Jobs, error)
	FindAllJobByTitle(string) ([]entities.Jobs, error)
}

func JobsUserInterface(db *gorm.DB) JobsRepository {
	return &JobsSqlStruct{Db: database.DBAccess}
}

func (r *JobsSqlStruct) FindAllJob() ([]entities.Jobs, error) {
	var jobs []entities.Jobs
	err := r.Db.Find(&jobs).Error
	if err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (r *JobsSqlStruct) FindAllJobByCategory(category string) ([]entities.Jobs, error) {
	var jobs []entities.Jobs
	err := r.Db.Where("job_category = ?", category).Find(&jobs).Error
	if err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (r *JobsSqlStruct) FindAllJobByTitle(title string) ([]entities.Jobs, error) {
	var jobs []entities.Jobs
	err := r.Db.Where("position_name LIKE ?", "%"+title+"%").Find(&jobs).Error
	if err != nil {
		return jobs, err
	}
	return jobs, nil
}