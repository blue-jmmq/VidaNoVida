package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
)

//convertir int a string
func ConvetirIntAString(entero int) string {
	return strconv.Itoa(entero)
}

//JSON es una función
func JSON(interfaz interface{}) string {
	bytesJSON, _ := json.Marshal(interfaz)
	return string(bytesJSON)
}

//JSONIdentado es una función
func JSONIdentado(interfaz interface{}) string {
	bytesJSON, _ := json.MarshalIndent(interfaz, "", "    ")
	return string(bytesJSON)
}

func Imprimir(interfaz ...interface{}) {
	fmt.Println(interfaz...)
}

//ImprimirJSON es una función
func ImprimirJSON(interfaz interface{}) {
	fmt.Println(JSON(interfaz))
}

//ImprimirIdentado es una función
func ImprimirIdentado(interfaz interface{}) {
	fmt.Println(JSONIdentado(interfaz))
}

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	juego := CrearJuego()
	juego.Jugar()
	//arreglo := CrearBidimensional(4, 2, nil)
	//ImprimirIdentado(arreglo)
}
