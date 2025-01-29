package entities

type Empleado struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Apellido          string `json:"apellido"`
	Area              string `json:"area"`
	CorreoElectronico string `json:"correo_electronico"`
	Deleted           bool   `json:"deleted"`
}
