package domain

import "github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"

type IEquipo interface {
	Save(equipo entities.Equipo) error
	Update(id int, equipo entities.Equipo) error
	Delete(id int) error
	FindByID(id int) (entities.Equipo, error)
	GetAll() ([]entities.Equipo, error)
}
