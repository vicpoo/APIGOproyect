package domain

import "github.com/vicpoo/APIGOproyect/empleados/domain/entities"

type IEmpleado interface {
	Save(empleado entities.Empleado) error
	Update(id int, empleado entities.Empleado) error
	Delete(id int) error
	FindByID(id int) (entities.Empleado, error)
}
