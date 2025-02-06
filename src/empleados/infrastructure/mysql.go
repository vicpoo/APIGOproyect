package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/APIGOproyect/src/core"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type MysqlEmpleado struct {
	conn *sql.DB
}

func NewMysqlEmpleadoRepository() domain.IEmpleado {
	conn := core.GetDB()
	return &MysqlEmpleado{conn: conn}
}

func (mysql *MysqlEmpleado) Save(empleado entities.Empleado) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO empleados (nombre, apellido, area, correo_electronico, deleted) VALUES (?, ?, ?, ?, ?)",
		empleado.Nombre,
		empleado.Apellido,
		empleado.Area,
		empleado.CorreoElectronico,
		empleado.Deleted,
	)
	if err != nil {
		log.Println("Error al guardar el empleado:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	empleado.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlEmpleado) Update(id int, empleado entities.Empleado) error {
	result, err := mysql.conn.Exec(
		"UPDATE empleados SET nombre = ?, apellido = ?, area = ?, correo_electronico = ?, deleted = ? WHERE id = ?",
		empleado.Nombre,
		empleado.Apellido,
		empleado.Area,
		empleado.CorreoElectronico,
		empleado.Deleted,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el empleado:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el empleado con ID:", id)
		return fmt.Errorf("empleado con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlEmpleado) Delete(id int) error {
	_, err := mysql.conn.Exec("UPDATE empleados SET deleted = true WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar (soft delete) el empleado:", err)
		return err
	}
	return nil
}

func (mysql *MysqlEmpleado) FindByID(id int) (entities.Empleado, error) {
	var empleado entities.Empleado
	row := mysql.conn.QueryRow("SELECT id, nombre, apellido, area, correo_electronico, deleted FROM empleados WHERE id = ?", id)

	err := row.Scan(
		&empleado.ID,
		&empleado.Nombre,
		&empleado.Apellido,
		&empleado.Area,
		&empleado.CorreoElectronico,
		&empleado.Deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Empleado no encontrado:", err)
			return entities.Empleado{}, fmt.Errorf("empleado con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el empleado por ID:", err)
		return entities.Empleado{}, err
	}

	return empleado, nil
}

func (mysql *MysqlEmpleado) GetAll() ([]entities.Empleado, error) {
	var empleados []entities.Empleado

	rows, err := mysql.conn.Query("SELECT id, nombre, apellido, area, correo_electronico, deleted FROM empleados")
	if err != nil {
		log.Println("Error al obtener todos los empleados:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var empleado entities.Empleado
		err := rows.Scan(
			&empleado.ID,
			&empleado.Nombre,
			&empleado.Apellido,
			&empleado.Area,
			&empleado.CorreoElectronico,
			&empleado.Deleted,
		)
		if err != nil {
			log.Println("Error al escanear empleado:", err)
			return nil, err
		}
		empleados = append(empleados, empleado)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return empleados, nil
}

func (mysql *MysqlEmpleado) GetByStatus(deleted bool) ([]entities.Empleado, error) {
	var empleados []entities.Empleado

	rows, err := mysql.conn.Query("SELECT id, nombre, apellido, area, correo_electronico, deleted FROM empleados WHERE deleted = ?", deleted)
	if err != nil {
		log.Println("Error al obtener empleados por estado:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var empleado entities.Empleado
		err := rows.Scan(
			&empleado.ID,
			&empleado.Nombre,
			&empleado.Apellido,
			&empleado.Area,
			&empleado.CorreoElectronico,
			&empleado.Deleted,
		)
		if err != nil {
			log.Println("Error al filtrar a los empleado:", err)
			return nil, err
		}
		empleados = append(empleados, empleado)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al filtar los empleados:", err)
		return nil, err
	}

	return empleados, nil
}
