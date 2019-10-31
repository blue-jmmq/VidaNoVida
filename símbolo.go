package main

import (
	"image"
	"os"

	_ "image/png"
)

type Símbolo struct {
	Arreglo *Bidimensional
}

func CrearSímbolo(datos [][]byte) *Símbolo {
	altura := len(datos)
	anchura := len(datos[0])
	var símbolo Símbolo
	símbolo.Arreglo = CrearBidimensional(anchura, altura, 0)
	interfaz := make([][]interface{}, len(datos))
	for fila := 0; fila < len(datos); fila++ {
		interfaz[fila] = make([]interface{}, len(datos[fila]))
		for columna := 0; columna < len(datos[fila]); columna++ {
			interfaz[fila][columna] = datos[fila][columna]
		}
	}
	símbolo.Arreglo.LlenarDesdeDatos(interfaz)
	return &símbolo
}

func CrearSímboloDesdeImagen(fuente, nombre string) *Símbolo {
	var símbolo Símbolo
	archivo, err := os.Open("fuentes/" + fuente + "/" + nombre + ".png")
	if err != nil {
		panic(err)
	}
	defer archivo.Close()
	imagen, _, err := image.Decode(archivo)
	if err != nil {
		panic(err)
	}
	//ImprimirJSON(formato)
	anchura := imagen.Bounds().Max.X - imagen.Bounds().Min.Y
	altura := imagen.Bounds().Max.Y - imagen.Bounds().Min.Y
	if anchura != altura {
		panic("anchura != altura")
	}
	símbolo.Arreglo = CrearBidimensional(anchura, altura, byte(0))
	for fila := 0; fila < altura; fila++ {
		for columna := 0; columna < anchura; columna++ {
			clr := imagen.At(columna, fila)
			rojo, verde, azul, _ := clr.RGBA()
			código := byte(0)
			if (rojo+verde+azul)/3 < 128 {
				código = 1
			}
			símbolo.Arreglo.Escribir(fila, columna, código)
		}
	}
	return &símbolo
}
