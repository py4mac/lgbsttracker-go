package interfaces

import (
	"fmt"

	"github.com/py4mac/lgbsttracker-go/application"
	"github.com/py4mac/lgbsttracker-go/domain/entity"
)

//Users struct defines the dependencies that will be used
type Experiments struct {
	r application.ExperimentAppInterface
}

//Experiments constructor
func NewExperiments(r application.ExperimentAppInterface) *Experiments {
	return &Experiments{
		r: r,
	}
}

//func (e *Experiments) CreateExperiment(c *gin.Context) {
func (e *Experiments) CreateExperiment(experiment entity.Experiment)  {

	newExperiment, err := e.r.CreateExperiment(experiment)
	if err != nil {
		return
	}
	fmt.Println(newExperiment)
	// var user entity.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"invalid_json": "invalid json",
	// 	})
	// 	return
	// }
	// //validate the request:
	// validateErr := user.Validate("")
	// if len(validateErr) > 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, validateErr)
	// 	return
	// }
	// newUser, err := s.us.SaveUser(&user)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	// c.JSON(http.StatusCreated, newUser.PublicUser())
}

//func (e *Experiments) GetExperimentByUuid(c *gin.Context) {
func (e *Experiments) GetExperimentByUuid(uuid int) {
	experiment, err := e.r.GetExperimentByUuid(uuid)
	if err != nil {
		return
	}
	fmt.Println(experiment)
	// users := entity.Users{} //customize user
	// var err error
	// //us, err = application.UserApp.GetUsers()
	// users, err = s.us.GetUsers()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, users.PublicUsers())
}
