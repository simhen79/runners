package services

import (
	"net/http"
	"runners/models"
	"runners/repositories"
	"time"
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

	err := validateResult(result)

	if err != nil {
		return nil, err
	}

	response, responseErr := rs.resultsRepository.CreateResult(result)
	if responseErr != nil {
		return nil, responseErr
	}

	runner, reresponseErr := rs.runnersRepository.GetRunner(result.RunnerID)
	if reresponseErr != nil {
		return nil, responseErr
	}

	if runner == nil {
		return nil, &models.ResponseError{
			Message: "Runner not found",
			Status:  http.StatusNotFound,
		}
	}

	err = updatePersonalBest(result, runner)
	if err != nil {
		return nil, err
	}

	err = updateSeasonBest(result, runner)
	if err != nil {
		return nil, err
	}

	responseErr = rs.runnersRepository.UpdateRunnerResults(runner)
	if reresponseErr != nil {
		return nil, responseErr
	}

	return response, nil
}

func (rs ResultsService) DeleteResult(resultId string) *models.ResponseError {
	if resultId == "" {
		return &models.ResponseError{
			Message: "Invalid result ID",
			Status:  http.StatusBadRequest,
		}
	}

	err := repositories.BeginTransaction(rs.runnersRepository, rs.resultsRepository)
	if err != nil {
		return &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	result, responseErr := rs.resultsRepository.DeleteResult(resultId)
	if responseErr != nil {
		return responseErr
	}

	runner, responseErr := rs.runnersRepository.GetRunner(result.RunnerID)
	if responseErr != nil {
		return responseErr
	}

	// Checking if deleted result is personal best
	if runner.PersonalBest == result.RaceResult {
		personalBest, responseErr := rs.resultsRepository.GetPersonalBestResults(result.RunnerID)
		if responseErr != nil {
			repositories.RollbackTransaction(rs.runnersRepository, rs.resultsRepository)
			return responseErr
		}
		runner.PersonalBest = personalBest
	}

	// Checking if deleted result is season best
	currentYear := time.Now().Year()
	if runner.SeasonBest == result.RaceResult && result.Year == currentYear {
		seasonBest, responseErr := rs.resultsRepository.GetSeasonBestResults(result.RunnerID, result.Year)
		if responseErr != nil {
			repositories.RollbackTransaction(rs.runnersRepository, rs.resultsRepository)
			return responseErr
		}
		runner.SeasonBest = seasonBest
	}

	responseErr = rs.runnersRepository.UpdateRunnerResults(runner)
	if responseErr != nil {
		repositories.RollbackTransaction(rs.runnersRepository, rs.resultsRepository)
		return responseErr
	}

	repositories.CommitTransaction(rs.runnersRepository, rs.resultsRepository)
	return nil
}

func updateSeasonBest(result *models.Result, runner *models.Runner) *models.ResponseError {

	raceResult, err := parseRaceResult(result.RaceResult)

	if err != nil {
		return &models.ResponseError{
			Message: "Invalid Race Result",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.SeasonBest == "" {
		runner.SeasonBest = result.RaceResult
	} else {
		seasonBest, err := parseRaceResult(runner.SeasonBest)
		if err != nil {
			return &models.ResponseError{
				Message: "Failed to parse season best",
				Status:  http.StatusInternalServerError,
			}
		}

		if raceResult < seasonBest {
			runner.SeasonBest = result.RaceResult
		}
	}

	return nil
}

func updatePersonalBest(result *models.Result, runner *models.Runner) *models.ResponseError {

	raceResult, err := parseRaceResult(result.RaceResult)

	if err != nil {
		return &models.ResponseError{
			Message: "Invalid Race Result",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.PersonalBest == "" {
		runner.PersonalBest = result.RaceResult
	} else {
		personalBest, err := parseRaceResult(runner.PersonalBest)
		if err != nil {
			return &models.ResponseError{
				Message: "Failed to parse personal best",
				Status:  http.StatusInternalServerError,
			}
		}

		if raceResult < personalBest {
			runner.PersonalBest = result.RaceResult
		}
	}

	return nil
}

func validateResult(result *models.Result) *models.ResponseError {
	if result.RunnerID == "" {
		return &models.ResponseError{
			Message: "Invalid Runner ID",
			Status:  http.StatusBadRequest,
		}
	}

	if result.RaceResult == "" {
		return &models.ResponseError{
			Message: "Invalid Race Result",
			Status:  http.StatusBadRequest,
		}
	}

	if result.Position < 0 {
		return &models.ResponseError{
			Message: "Invalid Position",
			Status:  http.StatusBadRequest,
		}
	}

	currentYear := time.Now().Year()
	if result.Year < 0 || result.Year > currentYear {
		return &models.ResponseError{
			Message: "Invalid Year",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func parseRaceResult(timeString string) (time.Duration, error) {
	return time.ParseDuration(timeString[0:2] + "h" + timeString[3:5] + "m" + timeString[6:8] + "s")
}
