package application

import (
	"github.com/vicpoo/APIGOproyect/empleados/domain"
	"github.com/vicpoo/APIGOproyect/empleados/domain/entities"
)

type CreateEmpleado struct {
	db domain.IEmpleado
}

func NewCreateEmpleado(db domain.IEmpleado) *CreateEmpleado {
	return &CreateEmpleado{db: db}
}

func (ce *CreateEmpleado) Execute(empleado entities.Empleado) error {
	return ce.db.Save(empleado)
}
