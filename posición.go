package main

type Posición struct {
	Fila    int
	Columna int
}

func CrearPosición(fila, columna int) *Posición {
	var posición Posición
	posición.Columna = columna
	posición.Fila = fila
	return &posición
}
