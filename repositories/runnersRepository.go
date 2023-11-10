package repositories

import (
	"database/sql"
	"runners/models"
)

type RunnersRepository struct {
}

func NewRunnersRepository(db *sql.DB) {

}

func (rp RunnersRepository) GetRunner(runnerID string) (*models.Runner, *models.ResponseError) {
	return nil, nil
}

func (rp RunnersRepository) UpdateRunnerResults(runner *models.Runner) *models.ResponseError {
	return nil
}
