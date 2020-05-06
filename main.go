package main

import (
	"os"
	"time"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/py4mac/lgbsttracker-go/cmd"
	"github.com/py4mac/lgbsttracker-go/domain/entity"
	"github.com/py4mac/lgbsttracker-go/infrastructure/persistence"
	_ "github.com/py4mac/lgbsttracker-go/interfaces"
)

func main() {
	var connectionString = os.Getenv("EXPERIMENT_STORAGE_URI")
	var dbName = "lgbsttracker"
	services, err := persistence.NewRepositories(connectionString, dbName)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	experiment1 := entity.Experiment{
		ExperimentUuid: 1,
		Ts:             time.Now(),
		Action:         "Right",
		VisionSensor:   0.0,
		Speed:          0.0,
		State:          "RUNNING",
	}

	
	newExperiment, err := services.Experiment.CreateExperiment(experiment1)
	if err != nil {
		return
	}
	fmt.Println(newExperiment)


	experiment, err := services.Experiment.GetExperimentByUuid(experiment1.ExperimentUuid)
	if err != nil {
		return
	}
	fmt.Println(experiment)


	// experiments := interfaces.NewExperiments(services.Experiment)
	// experiments.CreateExperiment(experiment1)
	// experiments.GetExperimentByUuid(1)


}