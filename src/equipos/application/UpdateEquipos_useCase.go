package application

import (
	"github.com/vicpoo/APIGOproyect/src/equipos/domain"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type UpdateEquipo struct {
	db domain.IEquipo
}

func NewUpdateEquipo(db domain.IEquipo) *UpdateEquipo {
	return &UpdateEquipo{db: db}
}

func (ue *UpdateEquipo) Execute(id int, equipo entities.Equipo) error {
	return ue.db.Update(id, equipo)
}
