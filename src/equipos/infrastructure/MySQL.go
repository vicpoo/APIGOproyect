package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/APIGOproyect/src/core"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type MysqlEquipo struct {
	conn *sql.DB
}

func NewMysqlEquipoRepository() domain.IEquipo {
	conn := core.GetDB()
	return &MysqlEquipo{conn: conn}
}

func (mysql *MysqlEquipo) Save(equipo entities.Equipo) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO equipos (nombre, marca, numero_serie, modelo) VALUES (?, ?, ?, ?)",
		equipo.Nombre,
		equipo.Marca,
		equipo.NumeroSerie,
		equipo.Modelo,
	)
	if err != nil {
		log.Println("Error al guardar el equipo:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	equipo.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlEquipo) Update(id int, equipo entities.Equipo) error {
	result, err := mysql.conn.Exec(
		"UPDATE equipos SET nombre = ?, marca = ?, numero_serie = ?, modelo = ? WHERE id = ?",
		equipo.Nombre,
		equipo.Marca,
		equipo.NumeroSerie,
		equipo.Modelo,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el equipo:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el equipo con ID:", id)
		return fmt.Errorf("equipo con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlEquipo) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM equipos WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el equipo:", err)
		return err
	}
	return nil
}

func (mysql *MysqlEquipo) FindByID(id int) (entities.Equipo, error) {
	var equipo entities.Equipo
	row := mysql.conn.QueryRow("SELECT id, nombre, marca, numero_serie, modelo FROM equipos WHERE id = ?", id)

	err := row.Scan(
		&equipo.ID,
		&equipo.Nombre,
		&equipo.Marca,
		&equipo.NumeroSerie,
		&equipo.Modelo,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Equipo no encontrado:", err)
			return entities.Equipo{}, fmt.Errorf("equipo con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el equipo por ID:", err)
		return entities.Equipo{}, err
	}

	return equipo, nil
}

func (mysql *MysqlEquipo) GetAll() ([]entities.Equipo, error) {
	var equipos []entities.Equipo

	rows, err := mysql.conn.Query("SELECT id, nombre, marca, numero_serie, modelo FROM equipos")
	if err != nil {
		log.Println("Error al obtener todos los equipos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var equipo entities.Equipo
		err := rows.Scan(
			&equipo.ID,
			&equipo.Nombre,
			&equipo.Marca,
			&equipo.NumeroSerie,
			&equipo.Modelo,
		)
		if err != nil {
			log.Println("Error al escanear equipo:", err)
			return nil, err
		}
		equipos = append(equipos, equipo)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return equipos, nil
}
