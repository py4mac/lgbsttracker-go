package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/py4mac/lgbsttracker-go/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	Experiment repository.ExperimentRepository
	db         *mongo.Database
	client     *mongo.Client
}

func NewRepositories(ConnexionString string, DbName string) (*Repositories, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(ConnexionString))

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(DbName)

	return &Repositories{
		Experiment: NewExperimentRepository(db),
		db:         db,
		client:     client,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := s.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")
	return nil
}
