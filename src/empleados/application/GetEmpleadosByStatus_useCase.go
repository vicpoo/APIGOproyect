package application

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type GetEmpleadosByStatusUseCase struct {
	db domain.IEmpleado
}

func NewGetEmpleadosByStatusUseCase(db domain.IEmpleado) *GetEmpleadosByStatusUseCase {
	return &GetEmpleadosByStatusUseCase{
		db: db,
	}
}

func (uc *GetEmpleadosByStatusUseCase) Run(status bool) ([]entities.Empleado, error) {
	return uc.db.GetByStatus(status)
}
