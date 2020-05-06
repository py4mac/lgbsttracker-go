package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/py4mac/lgbsttracker-go/domain/entity"
	"github.com/py4mac/lgbsttracker-go/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperimentRepo struct {
	db *mongo.Database
}

func NewExperimentRepository(db *mongo.Database) *ExperimentRepo {
	return &ExperimentRepo{db}
}

//ExperimentRepo implements the repository.ExperimentRepository interface
var _ repository.ExperimentRepository = &ExperimentRepo{}

func (e *ExperimentRepo) CreateExperiment(experiment entity.Experiment) (entity.Experiment, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := e.db.Collection("experiments").InsertOne(ctx, experiment)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)
	return experiment, nil
}

func (e *ExperimentRepo) GetExperimentByUuid(uuid int) (entity.Experiment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var entity entity.Experiment
	err := e.db.Collection("experiments").FindOne(ctx, bson.D{{"experiment_uuid", uuid}}).Decode(&entity)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return entity, nil

}
