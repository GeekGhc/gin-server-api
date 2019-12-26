package tag_service

import "gin-server-api/models"

type Tag struct {
	ID        int
	Name      string
	CreatedBy string
	UpdatedBy string
	Status    int

	PageNo   int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.Status, t.CreatedBy)
}
