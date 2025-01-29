package application

import "github.com/vicpoo/APIGOproyect/src/empleados/domain"

type DeleteEmpleado struct {
	db domain.IEmpleado
}

func NewDeleteEmpleado(db domain.IEmpleado) *DeleteEmpleado {
	return &DeleteEmpleado{db: db}
}

func (de *DeleteEmpleado) Execute(id int) error {
	return de.db.Delete(id)
}
