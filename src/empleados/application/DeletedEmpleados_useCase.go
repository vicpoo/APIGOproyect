package application

import "github.com/vicpoo/APIGOproyect/src/empleados/domain"

type DeleteEmpleadoUseCase struct {
	db domain.IEmpleado
}

func NewDeleteEmpleadoUseCase(db domain.IEmpleado) *DeleteEmpleadoUseCase {
	return &DeleteEmpleadoUseCase{
		db: db,
	}
}

func (uc *DeleteEmpleadoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
