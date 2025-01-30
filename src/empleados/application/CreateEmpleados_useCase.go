package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type CreateEmpleadoUseCase struct {
	db domain.IEmpleado
}

func NewCreateEmpleadoUseCase(db domain.IEmpleado) *CreateEmpleadoUseCase {
	return &CreateEmpleadoUseCase{
		db: db,
	}
}

func (uc *CreateEmpleadoUseCase) Run(empleado *entities.Empleado) (*entities.Empleado, error) {
	err := uc.db.Save(*empleado)
	return empleado, err
}
