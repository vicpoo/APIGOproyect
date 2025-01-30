package entities

type Empleado struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Apellido          string `json:"apellido"`
	Area              string `json:"area"`
	CorreoElectronico string `json:"correo_electronico"`
	Deleted           bool   `json:"deleted"`
}

func NewEmpleado(id int, nombre, apellido, area, correoElectronico string, deleted bool) *Empleado {
	return &Empleado{
		ID:                id,
		Nombre:            nombre,
		Apellido:          apellido,
		Area:              area,
		CorreoElectronico: correoElectronico,
		Deleted:           deleted,
	}
}

func (e *Empleado) GetID() int {
	return e.ID
}

func (e *Empleado) SetID(id int) {
	e.ID = id
}

func (e *Empleado) GetNombre() string {
	return e.Nombre
}

func (e *Empleado) SetNombre(nombre string) {
	e.Nombre = nombre
}

func (e *Empleado) GetApellido() string {
	return e.Apellido
}

func (e *Empleado) SetApellido(apellido string) {
	e.Apellido = apellido
}

func (e *Empleado) GetArea() string {
	return e.Area
}

func (e *Empleado) SetArea(area string) {
	e.Area = area
}

func (e *Empleado) GetCorreoElectronico() string {
	return e.CorreoElectronico
}

func (e *Empleado) SetCorreoElectronico(correoElectronico string) {
	e.CorreoElectronico = correoElectronico
}

func (e *Empleado) IsDeleted() bool {
	return e.Deleted
}

func (e *Empleado) SetDeleted(deleted bool) {
	e.Deleted = deleted
}
