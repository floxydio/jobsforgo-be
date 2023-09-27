package handler

import (
	"jobsforgobe/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryJobs struct {
	Repo repository.JobsRepository
}

func JobsController(db *gorm.DB) *RepositoryJobs {
	repo := repository.JobsUserInterface(db)
	return &RepositoryJobs{Repo: repo}
}

func (sql *RepositoryJobs) FindAllJobHandler(c echo.Context) error {
	jobs, err := sql.Repo.FindAllJob()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, jobs)
}

func (sql *RepositoryJobs) FindAllJobByCategoryHandler(c echo.Context) error {
	jobs, err := sql.Repo.FindAllJobByCategory(c.Param("category"))

	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, jobs)
}
