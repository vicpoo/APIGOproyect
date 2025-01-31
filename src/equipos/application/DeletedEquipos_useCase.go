package application

import "github.com/vicpoo/APIGOproyect/src/equipos/domain"

type DeleteEquipoUseCase struct {
	db domain.IEquipo
}

func NewDeleteEquipoUseCase(db domain.IEquipo) *DeleteEquipoUseCase {
	return &DeleteEquipoUseCase{
		db: db,
	}
}

func (uc *DeleteEquipoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
