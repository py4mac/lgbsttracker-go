package repository

import (
	"github.com/py4mac/lgbsttracker-go/domain/entity"
)

// ExperimentRepository Interface
type ExperimentRepository interface {
	CreateExperiment(entity.Experiment) (entity.Experiment, error)
	GetExperimentByUuid(int) (entity.Experiment, error)
}
