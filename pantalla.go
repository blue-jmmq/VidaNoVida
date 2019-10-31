package main

type Pantalla struct {
	Celdas *Toro
}

func CrearPantalla(anchura, altura int, símboloVacío *Símbolo) *Pantalla {
	var pantalla Pantalla
	pantalla.Celdas = CrearToro(anchura, altura, nil)
	for fila := 0; fila < altura; fila++ {
		for columna := 0; columna < anchura; columna++ {
			var celda Celda
			celda.Símbolo = símboloVacío
			pantalla.Celdas.Escribir(fila, columna, &celda)
		}
	}
	return &pantalla
}

func (pantalla *Pantalla) LeerAltura() int {
	return pantalla.Celdas.LeerAltura()
}

func (pantalla *Pantalla) LeerAnchura() int {
	return pantalla.Celdas.LeerAnchura()
}
