package entity

import (
	"time"
)

type Experiment struct {
	ExperimentUuid int
	Ts             time.Time
	Action         string
	VisionSensor   float32
	Speed          float32
	State          string
}
