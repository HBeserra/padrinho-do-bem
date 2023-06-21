package entity

import "time"

type Children struct {
	id               uint
	name             string
	birthDate        time.Time
	gender           string
	history          string
	addressCity      string
	addressState     string
	currentSituation string
	//legalGuardian
}

type MemoryChildren interface {
	Get(id uint) (Children, error)
	GetAll(offset int, perPage int) ([]Children, error)
	Save(c Children) (id uint, err error)
	Delete(c Children) error
	DeleteByID(id uint) error
}

func (c Children) GetName() string {
	return c.name
}
