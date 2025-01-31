package application

import (
	"github.com/vicpoo/APIGOproyect/src/equipos/domain"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type ViewAllEquipos struct {
	db domain.IEquipo
}

func NewViewAllEquipos(db domain.IEquipo) *ViewAllEquipos {
	return &ViewAllEquipos{db: db}
}

func (ve *ViewAllEquipos) Execute() ([]entities.Equipo, error) {
	return ve.db.GetAll()
}
