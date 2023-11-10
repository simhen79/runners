package repositories

import (
	"database/sql"
	"runners/models"
)

type ResultsRepository struct{}

func NewResultsRepository(db *sql.DB) {

}

func (rp ResultsRepository) DeleteResult(resultID string) (*models.Result, *models.ResponseError) {
	return nil, nil
}

func (rp ResultsRepository) GetPersonalBestResults(runnerID string) (*models.Result, *models.ResponseError) {
	return nil, nil
}

func (rp ResultsRepository) GetSeasonBestResults(runnerID string, year int) (*models.Result, *models.ResponseError) {
	return nil, nil
}
