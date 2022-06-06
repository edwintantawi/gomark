package id_gen

import "github.com/google/uuid"

type Uuid struct{}

func (_ *Uuid) Generate() string {
	return uuid.New().String()
}

func NewUuid() *Uuid {
	return &Uuid{}
}
