package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Fuente struct {
	Nombre   string
	Símbolos map[string]*Símbolo
	Imágenes map[string]string
	Cadenas  map[*Símbolo]string
}

func CrearFuente(nombre string) *Fuente {
	fuente := new(Fuente)
	fuente.Nombre = nombre
	fuente.LeerImágenes()
	fuente.LeerSímbolos()
	fuente.CrearCadenas()
	//Imprimir("Runas:", fuente.Runas)
	//Imprimir("Símbolos:", fuente.Símbolos)

	return fuente
}
func (fuente *Fuente) CrearCadenas() {
	fuente.Cadenas = make(map[*Símbolo]string)
	for nombre, simbolo := range fuente.Símbolos {
		fuente.Cadenas[simbolo] = nombre
	}
}

/*
func (fuente *Fuente) ContieneRuna(runa rune) bool {
	return false
}
*/

func (fuente *Fuente) LeerSímbolos() {
	fuente.Símbolos = make(map[string]*Símbolo)
	for nombre, imagen := range fuente.Imágenes {
		fuente.Símbolos[nombre] = CrearSímboloDesdeImagen(fuente.Nombre, imagen)
	}
}

func (fuente *Fuente) LeerImágenes() {
	archivo, err := os.Open("fuentes/" + fuente.Nombre + "/fuente.json")
	if err != nil {
		panic(err)
	}
	defer archivo.Close()

	bytesJSON, err := ioutil.ReadFile("fuentes/" + fuente.Nombre + "/fuente.json")
	if err != nil {
		panic(err)
	}
	var runas []interface{}
	err = json.Unmarshal(bytesJSON, &runas)
	if err != nil {
		panic(err)
	}
	//Imprimir("Runas:", runas)
	fuente.Imágenes = make(map[string]string)
	for _, dupla := range runas {
		dupla := dupla.([]interface{})
		nombre := dupla[0].(string)
		imagen := dupla[1].(string)
		//Imprimir(reflect.TypeOf(dupla[0]))
		fuente.Imágenes[nombre] = imagen
	}
}
