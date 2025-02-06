package infrastructure

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

func InitEmpleadoDependencies() (
	*CreateEmpleadoController,
	*ViewEmpleadoController,
	*UpdateEmpleadoController,
	*DeleteEmpleadoController,
	*ViewAllEmpleadosController,
	*GetEmpleadosByStatusController,
) {
	// Inicializar el repositorio
	repo := NewMysqlEmpleadoRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateEmpleadoUseCase(repo)
	viewUseCase := application.NewViewEmpleado(repo)
	updateUseCase := application.NewUpdateEmpleado(repo)
	deleteUseCase := application.NewDeleteEmpleadoUseCase(repo)
	viewAllUseCase := application.NewViewAllEmpleados(repo)
	getEmpleadosBystatus := application.NewGetEmpleadosByStatusUseCase(repo)

	// Crear controladores
	createController := NewCreateEmpleadoController(createUseCase)
	viewController := NewViewEmpleadoController(viewUseCase)
	updateController := NewUpdateEmpleadoController(updateUseCase)
	deleteController := NewDeleteEmpleadoController(deleteUseCase)
	viewAllController := NewViewAllEmpleadosController(viewAllUseCase)
	getEmpleadosBystatuscontroller := NewGetEmpleadosByStatusController(getEmpleadosBystatus)

	return createController, viewController, updateController, deleteController, viewAllController, getEmpleadosBystatuscontroller
}
