package application

import (
	"github.com/py4mac/lgbsttracker-go/domain/entity"
	"github.com/py4mac/lgbsttracker-go/domain/repository"
)

type experimentApp struct {
	r repository.ExperimentRepository
}

//ExperimentApp implements the ExperimentAppInterface
var _ ExperimentAppInterface = &experimentApp{}

type ExperimentAppInterface interface {
	CreateExperiment(entity.Experiment) (entity.Experiment, error)
	GetExperimentByUuid(int) (entity.Experiment, error)
}

func (e *experimentApp) CreateExperiment(experiment entity.Experiment) (entity.Experiment, error) {
	return e.r.CreateExperiment(experiment)
}

func (e *experimentApp) GetExperimentByUuid(uuid int) (entity.Experiment, error) {
	return e.r.GetExperimentByUuid(uuid)
}
