package main

//Bidimensional es una estructura
type Bidimensional struct {
	Interno [][]interface{}
	Anchura int
	Altura  int
}

//CrearBidimensional es una función que crea un nuevo Bidimensional
func CrearBidimensional(anchura, altura int, valorPorDefecto interface{}) *Bidimensional {
	var arreglo Bidimensional
	arreglo.Anchura = anchura
	arreglo.Altura = altura
	arreglo.ConstruirInterno()
	arreglo.Llenar(valorPorDefecto)
	return &arreglo
}

//ConstruirInterno es una función
func (arreglo *Bidimensional) ConstruirInterno() {
	arreglo.Interno = make([][]interface{}, arreglo.Altura)
	for fila := 0; fila < arreglo.Altura; fila++ {
		arreglo.Interno[fila] = make([]interface{}, arreglo.Anchura)
	}
}

//Llenar es una función
func (arreglo *Bidimensional) Llenar(valor interface{}) {
	for fila := 0; fila < arreglo.Altura; fila++ {
		for columna := 0; columna < arreglo.Anchura; columna++ {
			arreglo.Interno[fila][columna] = valor
		}
	}
}

//Leer es una función
func (arreglo *Bidimensional) Leer(fila, columna int) interface{} {
	return arreglo.Interno[fila][columna]
}

//Escribir es una función
func (arreglo *Bidimensional) Escribir(fila, columna int, valor interface{}) {
	arreglo.Interno[fila][columna] = valor
}

func (arreglo *Bidimensional) LeerAltura() int {
	return arreglo.Altura
}

func (arreglo *Bidimensional) LeerAnchura() int {
	return arreglo.Anchura
}

func (arreglo *Bidimensional) LlenarDesdeDatos(datos [][]interface{}) {
	for fila := 0; fila < arreglo.Altura; fila++ {
		for columna := 0; columna < arreglo.Anchura; columna++ {
			arreglo.Interno[fila][columna] = datos[fila][columna]
		}
	}
}
