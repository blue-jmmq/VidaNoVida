package main

type WidgetDeSalida struct {
	Anchura  int
	Altura   int
	Símbolos *Bidimensional
	Búfer    [][]*Símbolo
	Entradas []*Bidimensional
	Índice   int
}

func (widget *WidgetDeSalida) DesplazarseHaciaAbajo(fuente *Fuente) {
	altura := len(widget.Búfer)
	if widget.Índice < altura {
		widget.Índice++
		widget.ActualizarSímbolos(fuente)
	}
}
func (widget *WidgetDeSalida) DesplazarseHaciaArriba(fuente *Fuente) {
	if widget.Índice > 0 {
		widget.Índice--
		widget.ActualizarSímbolos(fuente)
	}
}
func (widget *WidgetDeSalida) AñadirEntrada(entrada *Bidimensional) {
	for fila := 0; fila < entrada.LeerAltura(); fila++ {
		var línea []*Símbolo
		for columna := 0; columna < entrada.LeerAnchura(); columna++ {
			línea = append(línea, entrada.Leer(fila, columna).(*Símbolo))
		}
		widget.Búfer = append(widget.Búfer, línea)
	}
}

func (widget *WidgetDeSalida) ActualizarÍndice() {
	alturaDelBúfer := len(widget.Búfer)
	if alturaDelBúfer >= widget.Altura {
		widget.Índice = alturaDelBúfer - widget.Altura
	}
}

func (widget *WidgetDeSalida) ActualizarSímbolos(fuente *Fuente) {
	alturaDelBúfer := len(widget.Búfer)
	if alturaDelBúfer < widget.Altura {
		for fila := 0; fila < alturaDelBúfer; fila++ {
			for columna := 0; columna < widget.Anchura; columna++ {
				widget.Símbolos.Escribir(fila, columna, widget.Búfer[fila][columna])
			}
		}
		for fila := len(widget.Búfer); fila < len(widget.Búfer); fila++ {
			for columna := 0; columna < widget.Símbolos.LeerAnchura(); columna++ {
				widget.Símbolos.Escribir(fila, columna, fuente.Símbolos[" "])
			}
		}
	} else if alturaDelBúfer < widget.Índice+widget.Altura {
		//Imprimir("widget.Altura", widget.Altura)
		//Imprimir("alturaDelBúfer", alturaDelBúfer)
		//Imprimir("widget.Índice", widget.Índice)
		var fila int
		for fila = widget.Índice; fila < alturaDelBúfer; fila++ {
			for columna := 0; columna < widget.Anchura; columna++ {
				widget.Símbolos.Escribir(fila-widget.Índice, columna, widget.Búfer[fila][columna])
			}
		}
		for ; fila < widget.Índice+widget.Altura; fila++ {
			for columna := 0; columna < widget.Anchura; columna++ {
				widget.Símbolos.Escribir(fila-widget.Índice, columna, fuente.Símbolos[" "])
			}
		}

	} else {
		//Imprimir("widget.Altura", widget.Altura)
		//Imprimir("alturaDelBúfer", alturaDelBúfer)
		//Imprimir("widget.Índice", widget.Índice)
		for fila := widget.Índice; fila < widget.Índice+widget.Altura; fila++ {
			for columna := 0; columna < widget.Anchura; columna++ {
				widget.Símbolos.Escribir(fila-widget.Índice, columna, widget.Búfer[fila][columna])
			}
		}
	}
}

func (widget *WidgetDeSalida) Escribir(línea []*Símbolo, fuente *Fuente) {
	entrada := widget.LíneaAEntrada(línea, fuente)
	widget.Entradas = append(widget.Entradas, entrada)
	widget.AñadirEntrada(entrada)
	widget.ActualizarÍndice()
	widget.ActualizarSímbolos(fuente)
}

//GetDrawableLines function
func (widget *WidgetDeSalida) LíneaAEntrada(línea []*Símbolo, fuente *Fuente) *Bidimensional {
	var entrada *Bidimensional
	anchura := widget.Anchura
	var altura int
	cantidadaDeSímbolos := len(línea)
	if cantidadaDeSímbolos <= widget.Anchura {
		altura = 1
		entrada = CrearBidimensional(anchura, altura, fuente.Símbolos[" "])
		for columna := 0; columna < cantidadaDeSímbolos; columna++ {
			entrada.Escribir(0, columna, línea[columna])
		}
	} else {
		cantidadDeLíneasCompletas := cantidadaDeSímbolos / anchura
		cantidadDeSímbolosSobrantes := cantidadaDeSímbolos % anchura
		sobranSímbolos := cantidadDeSímbolosSobrantes > 0
		if sobranSímbolos {
			altura = cantidadDeLíneasCompletas + 1
		} else {
			altura = cantidadDeLíneasCompletas
		}
		entrada = CrearBidimensional(anchura, altura, fuente.Símbolos[" "])

		for fila := 0; fila < cantidadDeLíneasCompletas; fila++ {
			for columna := 0; columna < anchura; columna++ {
				entrada.Escribir(fila, columna, línea[fila*anchura+columna])
			}
		}
		if sobranSímbolos {
			fila := cantidadDeLíneasCompletas
			for columna := 0; columna < cantidadDeSímbolosSobrantes; columna++ {
				entrada.Escribir(fila, columna, línea[fila*anchura+columna])
			}
		}
	}
	return entrada
}
