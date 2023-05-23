package object

import (
	"reflect"
)

type Animal struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Species     string `json:"species"`
	IsCarnivore bool   `json:"is_carnivore"`
	Kind        string `json:"kind"`
}

func (a *Animal) GetKind() string {
	return reflect.TypeOf(a).String()
}

func (a *Animal) GetID() string {
	return a.ID
}

func (a *Animal) GetName() string {
	return a.Name
}

func (a *Animal) SetID(s string) {
	a.ID = s
}

func (a *Animal) SetName(s string) {
	a.Name = s
}
