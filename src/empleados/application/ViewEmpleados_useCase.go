package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type ViewEmpleado struct {
	db domain.IEmpleado
}

func NewViewEmpleado(db domain.IEmpleado) *ViewEmpleado {
	return &ViewEmpleado{db: db}
}

func (ve *ViewEmpleado) Execute(id int) (entities.Empleado, error) {
	return ve.db.FindByID(id)
}
