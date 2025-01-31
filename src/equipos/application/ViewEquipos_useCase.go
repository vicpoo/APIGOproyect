package application

import (
	"github.com/vicpoo/APIGOproyect/src/equipos/domain"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type ViewEquipo struct {
	db domain.IEquipo
}

func NewViewEquipo(db domain.IEquipo) *ViewEquipo {
	return &ViewEquipo{db: db}
}

func (ve *ViewEquipo) Execute(id int) (entities.Equipo, error) {
	return ve.db.FindByID(id)
}
