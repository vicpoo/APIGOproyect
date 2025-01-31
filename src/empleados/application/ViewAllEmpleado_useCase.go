package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

// ViewAllEmpleados es el caso de uso para obtener todos los empleados.
type ViewAllEmpleados struct {
	db domain.IEmpleado
}

// NewViewAllEmpleados crea una nueva instancia de ViewAllEmpleados.
func NewViewAllEmpleados(db domain.IEmpleado) *ViewAllEmpleados {
	return &ViewAllEmpleados{db: db}
}

// Execute ejecuta el caso de uso y devuelve la lista de empleados.
func (ve *ViewAllEmpleados) Execute() ([]entities.Empleado, error) {
	return ve.db.GetAll()
}
