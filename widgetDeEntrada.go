package main

type WidgetDeEntrada struct {
	Anchura           int
	Altura            int
	Símbolos          *Bidimensional
	Búfer             []*Símbolo
	InicioVisible     int
	Índice            int
	TamañoVisual      int
	DeboDibujarCursor bool
}

func (widget *WidgetDeEntrada) Tick() {
	widget.DeboDibujarCursor = !widget.DeboDibujarCursor
}

func (widget *WidgetDeEntrada) RestablecerCursor() {
	widget.DeboDibujarCursor = true
}

func (widget *WidgetDeEntrada) EliminarSímbolo() {
	if widget.Índice > 0 {
		var provicional []*Símbolo
		provicional = append(provicional, widget.Búfer[0:widget.Índice-1]...)
		provicional = append(provicional, widget.Búfer[widget.Índice:]...)
		widget.Búfer = provicional
		widget.DesplazarCursor(-1)
	}
}

func (widget *WidgetDeEntrada) Escribir(símbolo *Símbolo) {
	var provicional []*Símbolo
	provicional = append(provicional, widget.Búfer[:widget.Índice]...)
	provicional = append(provicional, símbolo)
	provicional = append(provicional, widget.Búfer[widget.Índice:]...)
	widget.Búfer = provicional
	widget.DesplazarCursor(1)
}

func (widget *WidgetDeEntrada) DesplazarCursor(delta int) {
	widget.Índice += delta
	if widget.Índice < 0 {
		widget.Índice = 0
		widget.InicioVisible = 0
		return
	}
	if widget.Índice > len(widget.Búfer) {
		widget.Índice -= delta
		return
	}
	if widget.Índice >= widget.InicioVisible && widget.Índice < widget.InicioVisible+widget.TamañoVisual {
		return
	}
	if delta < 0 {
		widget.InicioVisible = widget.Índice
	} else {
		widget.InicioVisible = widget.Índice - widget.TamañoVisual
	}
}

func (widget *WidgetDeEntrada) LlenarSímbolos(fuente *Fuente) {
	for columna := 0; columna < widget.Anchura; columna++ {
		widget.Símbolos.Escribir(0, columna, fuente.Símbolos["bloque"])
	}
	for columna := 0; columna < widget.Anchura; columna++ {
		widget.Símbolos.Escribir(1, columna, fuente.Símbolos[" "])
	}
	widget.Símbolos.Escribir(1, 1, fuente.Símbolos["C"])

	for columna := 0; columna < widget.Anchura; columna++ {
		widget.Símbolos.Escribir(2, columna, fuente.Símbolos["bloque"])
	}
}

func (widget *WidgetDeEntrada) Leer(fila, columna int, fuente *Fuente) *Símbolo {
	switch fila {
	case 0:
		return fuente.Símbolos["línea horizontal"]
	case 1:
		if columna >= 0 && columna < 3 {
			if columna == 1 {
				return fuente.Símbolos["C"]
			}
			return fuente.Símbolos[" "]
		}
		if columna == widget.Anchura-1 {
			return fuente.Símbolos[" "]
		}
		columna -= 3
		columna += widget.InicioVisible
		if columna == widget.Índice && widget.DeboDibujarCursor {
			return fuente.Símbolos["bloque"]
		}
		if columna < len(widget.Búfer) {
			return widget.Búfer[columna]
		}
		return fuente.Símbolos[" "]
	case 2:
		return fuente.Símbolos["línea horizontal"]
	}
	return fuente.Símbolos["nulo"]
}
