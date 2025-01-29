package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type UpdateEmpleado struct {
	db domain.IEmpleado
}

func NewUpdateEmpleado(db domain.IEmpleado) *UpdateEmpleado {
	return &UpdateEmpleado{db: db}
}

func (ue *UpdateEmpleado) Execute(id int, empleado entities.Empleado) error {
	return ue.db.Update(id, empleado)
}
