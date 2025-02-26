package registry

import (
	"errors"
)

type IRegistry interface {
	Create(user Registry) error
	Read(id int) (*Registry, error)
	Update(id int, user Registry) error
	Delete(id int) error
}

type Registry struct {
	Id   int
	Name string
	URI  string
}

func (r *Registry) TableName() string {
	return "registry"
}

func (r *Registry) IsValid() error {
	if r.Name == "" || r.URI == "" {
		return errors.New("registry data invalid")
	}
	return nil
}
