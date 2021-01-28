package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaVinculacion_20210127_160659 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaVinculacion_20210127_160659{}
	m.Created = "20210127_160659"

	migration.Register("CrearTablaVinculacion_20210127_160659", m)
}

// Run the migrations
func (m *CrearTablaVinculacion_20210127_160659) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210127_160659_crear_tabla_vinculacion_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearTablaVinculacion_20210127_160659) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210127_160659_crear_tabla_vinculacion_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
