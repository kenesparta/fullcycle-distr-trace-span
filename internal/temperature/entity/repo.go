package entity

import (
	"context"
)

type LocationRepo interface {
	Get(ctx context.Context, cep string) (Location, error)
}

type TemperatureRepo interface {
	Get(ctx context.Context, location string) (Temperature, error)
}
