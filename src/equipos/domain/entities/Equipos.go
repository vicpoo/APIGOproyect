package entities

type Equipo struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Marca       string `json:"marca"`
	NumeroSerie string `json:"numero_serie"`
	Modelo      string `json:"modelo"`
}

func NewEquipo(id int, nombre, marca, numeroSerie, modelo string) *Equipo {
	return &Equipo{
		ID:          id,
		Nombre:      nombre,
		Marca:       marca,
		NumeroSerie: numeroSerie,
		Modelo:      modelo,
	}
}

func (e *Equipo) GetID() int {
	return e.ID
}

func (e *Equipo) SetID(id int) {
	e.ID = id
}

func (e *Equipo) GetNombre() string {
	return e.Nombre
}

func (e *Equipo) SetNombre(nombre string) {
	e.Nombre = nombre
}

func (e *Equipo) GetMarca() string {
	return e.Marca
}

func (e *Equipo) SetMarca(marca string) {
	e.Marca = marca
}

func (e *Equipo) GetNumeroSerie() string {
	return e.NumeroSerie
}

func (e *Equipo) SetNumeroSerie(numeroSerie string) {
	e.NumeroSerie = numeroSerie
}

func (e *Equipo) GetModelo() string {
	return e.Modelo
}

func (e *Equipo) SetModelo(modelo string) {
	e.Modelo = modelo
}
