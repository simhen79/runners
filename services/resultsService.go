package services

import (
	"runners/models"
	"runners/repositories"
)

type ResultsService struct {
	runnersRepository *repositories.RunnersRepository
	resultsRepository *repositories.ResultsRepository
}

func NewResultsService(
	runnersRepository *repositories.RunnersRepository,
	resultsRepository *repositories.ResultsRepository,
) *ResultsService {
	return &ResultsService{
		runnersRepository: runnersRepository,
		resultsRepository: resultsRepository,
	}
}

func (rs ResultsService) CreateResult(result *models.Result) (*models.Result, *models.ResponseError) {
	return nil, nil
}

func (rs ResultsService) DeleteResult(resultId string) *models.ResponseError {
	return nil
}
