package services

import (
	"runners/models"
	"runners/repositories"
)

type RunnersService struct {
	runnersRepository *repositories.RunnersRepository
	resultsRepository *repositories.ResultsRepository
}

func NewRunnersService(
	runnersRepository *repositories.RunnersRepository,
	resultsRepository *repositories.ResultsRepository,
) *RunnersService {
	return &RunnersService{
		runnersRepository: runnersRepository,
		resultsRepository: resultsRepository,
	}
}

func (rs RunnersService) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) {
	return nil, nil
}

func (rs RunnersService) UpdateRunner(runner *models.Runner) *models.ResponseError {
	return nil
}

func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {
	return nil
}

func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, *models.ResponseError) {
	return nil, nil
}

func (rs RunnersService) GetRunnersBatch(country string, year string) ([]*models.Runner, *models.ResponseError) {
	return nil, nil
}
