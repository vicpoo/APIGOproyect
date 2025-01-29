package application

import "github.com/vicpoo/APIGOproyect/src/empleados/domain"

type ViewEmpleado struct {
	db domain.IEmpleado
}

func NewViewEmpleado(db domain.IEmpleado) *ViewEmpleado {
	return &ViewEmpleado{db: db}
}

func (ve *ViewEmpleado) Execute(id int) (domain.Empleado, error) {
	return ve.db.FindByID(id)
}
