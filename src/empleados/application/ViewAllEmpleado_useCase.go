package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type ViewAllEmpleados struct {
	db domain.IEmpleado
}

func NewViewAllEmpleados(db domain.IEmpleado) *ViewAllEmpleados {
	return &ViewAllEmpleados{db: db}
}

func (ve *ViewAllEmpleados) Execute() ([]entities.Empleado, error) {
	return ve.db.GetAll()
}
