package pokectx

import (
	"sync"
	"time"
)


type Service struct {
	Name string
	StartedAt time.Time
	mu sync.RWMutex
}

type Services struct {
	db DB
}

type Servicable interface {
	Run()
}