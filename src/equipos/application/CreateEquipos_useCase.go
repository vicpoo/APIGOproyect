package application

import (
	"github.com/vicpoo/APIGOproyect/src/equipos/domain"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type CreateEquipoUseCase struct {
	db domain.IEquipo
}

func NewCreateEquipoUseCase(db domain.IEquipo) *CreateEquipoUseCase {
	return &CreateEquipoUseCase{
		db: db,
	}
}

func (uc *CreateEquipoUseCase) Run(equipo *entities.Equipo) (*entities.Equipo, error) {
	err := uc.db.Save(*equipo)
	return equipo, err
}
