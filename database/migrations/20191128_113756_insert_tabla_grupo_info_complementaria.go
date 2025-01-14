package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaGrupoInfoComplementaria_20191128_113756 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaGrupoInfoComplementaria_20191128_113756{}
	m.Created = "20191128_113756"

	migration.Register("InsertTablaGrupoInfoComplementaria_20191128_113756", m)
}

// Run the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191128_113756) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (11, 'Información Adicional', 'Información Adicional.', 'Grupo_11', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('DESCRIPCIÓN', 'DESCRIPCIÓN', TRUE, 11, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
}

// Reverse the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191128_113756) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 11;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 11;")
}
