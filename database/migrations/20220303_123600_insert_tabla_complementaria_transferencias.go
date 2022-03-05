package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaComplementariaTransferencias_20220303_123600 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaComplementariaTransferencias_20220303_123600{}
	m.Created = "20220303_123600"

	migration.Register("InsertTablaComplementariaTransferencias_20220303_123600", m)
}

// Run the migrations
func (m *InsertTablaComplementariaTransferencias_20220303_123600) Up() {
	file, err := ioutil.ReadFile("../scripts/20220303_123600_insert_tabla_complementaria_transferencias_up.sql")

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
func (m *InsertTablaComplementariaTransferencias_20220303_123600) Down() {
	file, err := ioutil.ReadFile("../scripts/20220303_123600_insert_tabla_complementaria_transferencias_down.sql")

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
