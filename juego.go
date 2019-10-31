package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Juego struct {
	IU                      *IU
	Pantalla                *Pantalla
	Fuente                  *Fuente
	Pixeles                 *Bidimensional
	PseudoPixeles           *Bidimensional
	PseudoTamaño            int
	TamañoDeSímbolo         int
	SímboloVacío            *Símbolo
	Colores                 []*Color
	Implementación          *Implementación
	Altura                  int
	Anchura                 int
	HayQueDibujar           bool
	TemporizadorDelCursor   *time.Timer
	CanalDeComandos         chan []*Símbolo
	TiempoDeSeleccion       int
	TemporizadorDeSeleccion *time.Timer
	TiempoTranscurrido      int
	Jugador1                *Jugador
	Jugador2                *Jugador
}

//CrearDatosDelJuego es una función
func CrearJuego() *Juego {
	juego := new(Juego)
	juego.Anchura = 1024
	juego.Altura = 512
	juego.PseudoTamaño = 1
	juego.TamañoDeSímbolo = 16
	juego.HayQueDibujar = true
	juego.CanalDeComandos = make(chan []*Símbolo, 64)

	/*juego.SímboloVacío = CrearSímbolo([][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	})*/
	juego.Fuente = CrearFuente("principal")
	pantalla := CrearPantalla(
		juego.Anchura/juego.TamañoDeSímbolo/juego.PseudoTamaño,
		juego.Altura/juego.TamañoDeSímbolo/juego.PseudoTamaño,
		juego.Fuente.Símbolos["nulo"],
	)
	juego.Pantalla = pantalla
	juego.Implementación = CrearImplementación(juego)
	juego.Colores = append(juego.Colores, &Color{Rojo: 255, Verde: 255, Azul: 255})
	juego.Colores = append(juego.Colores, &Color{Rojo: 0, Verde: 0, Azul: 0})

	//juego.Pixeles = CrearBidimensional(juego.Anchura, juego.Altura, Píxel{Color: Color{Rojo: 0, Verde: 0, Azul: 0}})
	/*juego.PseudoPixeles = CrearBidimensional(
		juego.Anchura/juego.PseudoTamaño,
		juego.Altura/juego.PseudoTamaño,
		Píxel{Color: Color{Rojo: 0, Verde: 0, Azul: 0}},
	)*/
	juego.Pixeles = juego.CrearPixeles(juego.Anchura, juego.Altura)
	juego.PseudoPixeles = juego.CrearPixeles(juego.Anchura/juego.PseudoTamaño, juego.Altura/juego.PseudoTamaño)
	juego.IU = juego.CrearIU()
	juego.Dibujar()
	return juego
}

func (juego *Juego) CrearWidgetDeEntrada() *WidgetDeEntrada {
	widgetDeEntrada := new(WidgetDeEntrada)
	widgetDeEntrada.Anchura = juego.IU.Anchura
	widgetDeEntrada.Altura = 3
	widgetDeEntrada.Búfer = make([]*Símbolo, 0)
	widgetDeEntrada.TamañoVisual = widgetDeEntrada.Anchura - 5
	widgetDeEntrada.Símbolos = CrearBidimensional(widgetDeEntrada.Anchura, widgetDeEntrada.Altura, juego.Fuente.Símbolos["nulo"])
	widgetDeEntrada.LlenarSímbolos(juego.Fuente)
	return widgetDeEntrada
}

func (juego *Juego) CrearWidgetDeSalida() *WidgetDeSalida {
	widgetDeSalida := new(WidgetDeSalida)
	widgetDeSalida.Anchura = juego.IU.Anchura
	widgetDeSalida.Altura = juego.IU.Altura - 3
	widgetDeSalida.Símbolos = CrearBidimensional(widgetDeSalida.Anchura, widgetDeSalida.Altura, juego.Fuente.Símbolos[" "])
	return widgetDeSalida
}

func (juego *Juego) CrearIU() *IU {
	iu := new(IU)
	juego.IU = iu
	iu.Altura = juego.Pantalla.LeerAltura()
	iu.Anchura = juego.Pantalla.LeerAnchura()
	iu.WidgetDeEntrada = juego.CrearWidgetDeEntrada()
	iu.WidgetDeSalida = juego.CrearWidgetDeSalida()
	iu.Símbolos = CrearBidimensional(iu.Anchura, iu.Altura, juego.Fuente.Símbolos["nulo"])
	return iu
}

func (juego *Juego) CrearPixeles(anchura, altura int) *Bidimensional {
	pixeles := CrearBidimensional(anchura, altura, nil)
	for fila := 0; fila < altura; fila++ {
		for columna := 0; columna < anchura; columna++ {
			píxel := new(Píxel)
			píxel.Color = juego.Colores[0]
			pixeles.Escribir(fila, columna, píxel)
		}
	}
	return pixeles
}

func (juego *Juego) DibujarPseudoPíxel(color *Color, fila, columna int) {
	yInicial := fila * juego.PseudoTamaño
	xInicial := columna * juego.PseudoTamaño
	yFinal := yInicial + juego.PseudoTamaño
	xFinal := xInicial + juego.PseudoTamaño
	for y := yInicial; y < yFinal; y++ {
		for x := xInicial; x < xFinal; x++ {
			píxel := juego.Pixeles.Leer(y, x).(*Píxel)
			píxel.Color = color
		}
	}
}

func (juego *Juego) DibujarSímbolo(símbolo *Símbolo, fila, columna int) {
	if símbolo == nil {
		//Imprimir("Fila", fila)
		//Imprimir("Columna", columna)
		símbolo = juego.Fuente.Símbolos["nulo"]
	}
	yInicial := fila * juego.TamañoDeSímbolo
	xInicial := columna * juego.TamañoDeSímbolo
	yFinal := yInicial + juego.TamañoDeSímbolo
	xFinal := xInicial + juego.TamañoDeSímbolo
	for y := yInicial; y < yFinal; y++ {
		for x := xInicial; x < xFinal; x++ {
			colorIndex := símbolo.Arreglo.Leer(y-yInicial, x-xInicial).(byte)
			color := juego.Colores[colorIndex]
			píxel := juego.PseudoPixeles.Leer(y, x).(*Píxel)
			píxel.Color = color
			juego.DibujarPseudoPíxel(color, y, x)
		}
	}
}

func (juego *Juego) DibujarCelda(fila, columna int) {
	interfaz := juego.Pantalla.Celdas.Leer(fila, columna)
	símbolo := interfaz.(*Símbolo)
	juego.DibujarSímbolo(símbolo, fila, columna)
}

func (juego *Juego) ActualizarIU() {
	for fila := 0; fila < juego.IU.WidgetDeSalida.Altura; fila++ {
		for columna := 0; columna < juego.IU.WidgetDeSalida.Anchura; columna++ {
			juego.IU.Símbolos.Escribir(fila, columna, juego.IU.WidgetDeSalida.Símbolos.Leer(fila, columna))
		}
	}
	for fila := juego.IU.WidgetDeSalida.Altura; fila < juego.IU.WidgetDeSalida.Altura+juego.IU.WidgetDeEntrada.Altura; fila++ {
		for columna := 0; columna < juego.IU.WidgetDeEntrada.Anchura; columna++ {
			juego.IU.Símbolos.Escribir(fila, columna, juego.IU.WidgetDeEntrada.Leer(fila-juego.IU.WidgetDeSalida.Altura, columna, juego.Fuente))
		}
	}
}

func (juego *Juego) ActualizarPantalla() {
	juego.ActualizarIU()
	for fila := 0; fila < juego.Pantalla.LeerAltura(); fila++ {
		for columna := 0; columna < juego.Pantalla.LeerAnchura(); columna++ {
			juego.Pantalla.Celdas.Escribir(fila, columna, juego.IU.Símbolos.Leer(fila, columna))
		}
	}
}

func (juego *Juego) Tick() {
	juego.TemporizadorDelCursor.Reset(512 * time.Millisecond)
	juego.Dibujar()

}

func (juego *Juego) BucleDelTemporizador() {
	for {
		<-juego.TemporizadorDelCursor.C
		juego.IU.WidgetDeEntrada.Tick()
		juego.Tick()
		//Imprimir("Timer 1 expired")
	}

}

func (juego *Juego) Dibujar() {
	juego.ActualizarPantalla()
	for fila := 0; fila < juego.Pantalla.LeerAltura(); fila++ {
		for columna := 0; columna < juego.Pantalla.LeerAnchura(); columna++ {
			juego.DibujarCelda(fila, columna)
		}
	}
	juego.HayQueDibujar = true
}

func (juego *Juego) RunaASímbolo(runa rune) *Símbolo {
	cadena := string(runa)
	cadena = strings.ToUpper(cadena)
	símbolo := juego.Fuente.Símbolos[cadena]
	if símbolo == nil {
		símbolo = juego.Fuente.Símbolos["nulo"]
	}
	return símbolo
}

func (juego *Juego) CadenaASímbolos(cadena string) []*Símbolo {
	var línea []*Símbolo
	cadena = strings.ToUpper(cadena)
	runas := []rune(cadena)
	var cadenas []string
	for índice := 0; índice < len(runas); índice++ {
		cadenas = append(cadenas, string(runas[índice]))
	}
	for índice := 0; índice < len(cadenas); índice++ {
		línea = append(línea, juego.Fuente.Símbolos[cadenas[índice]])
	}
	return línea
}

func (juego *Juego) Escribir(cadena string) {
	cadenas := strings.Split(cadena, "\n")
	//Imprimir(cadena)
	var líneas [][]*Símbolo
	for índice := 0; índice < len(cadenas); índice++ {
		línea := juego.CadenaASímbolos(cadenas[índice])
		líneas = append(líneas, línea)
	}
	var separador []*Símbolo
	separador = append(separador, juego.Fuente.Símbolos[" "])
	for índice := 0; índice < len(líneas); índice++ {
		juego.IU.WidgetDeSalida.Escribir(líneas[índice], juego.Fuente)
	}
	juego.IU.WidgetDeSalida.Escribir(separador, juego.Fuente)
}

func (juego *Juego) Jugar() {
	juego.Implementación.Correr()
}

//strings.EqualFold
func (juego *Juego) EsperarComando() []*Símbolo {
	command := <-juego.CanalDeComandos
	return command
}

func (juego *Juego) LeerComando() []*Símbolo {
	select {
	case command := <-juego.CanalDeComandos:
		return command
    default:
        return nil
    }
}

//agregar carta Aleatoria
func (juego *Juego) AgregarCartaAleatoria(jugador *Jugador) {
	indiceAleatorio := rand.Intn(len(ArregloDeCartas))
	funcionCreadora := ArregloDeConstructoresDeCartas[indiceAleatorio]
	carta := funcionCreadora()
	jugador.Mano = append(jugador.Mano, carta)
}

//crear jugador
func (juego *Juego) Crearjugador() *Jugador {
	jugador := new(Jugador)
	jugador.Vida = 32
	jugador.ArmaduraRoja = 12
	jugador.Credito = 4
	jugador.Mano = nil
	//Tablero      [][]Cartajugador.
	return jugador
}
func (juego *Juego) JugarMultijugador() {

}
func (juego *Juego) JugarUnJugador() {
	juego.Jugador1 = juego.Crearjugador()
	juego.Jugador2 = juego.Crearjugador()
	juego.Escribir("Escoje cuatro cartas para tu mano inicial")
	juego.Escribir("Presiona la tecla ENTER para ver las cartas disponibles")
	juego.Dibujar()
	juego.EsperarComando()
	for i := 0; i < len(ArregloDeCartas); i++ {
		juego.Escribir(strconv.Itoa(i+1) + ")\n" + JSONIdentado(ArregloDeCartas[i].ObtenerInterfaz()))
	}
	juego.Escribir("Presiona Enter cuando estes listo para escojer una carta")
	juego.Dibujar()
	juego.EsperarComando()
	go juego.HiloDeTemporizadorDeSeleccion()
a:
	comando := juego.LeerComando()
	if juego.TiempoTranscurrido >= 10 {
		juego.Escribir("Lo sentimos ya acabo tu tiempo")
		juego.Dibujar()
		juego.AgregarCartaAleatoria(juego.Jugador1)
	} else if comando != nil{
		cadenaDeComando := juego.SimbolosACadena(comando)
		NCarta, err := juego.ConvertirStringAInt(cadenaDeComando)
		if err != nil || NCarta < 1 || NCarta > len(ArregloDeCartas) {
			juego.Escribir("Solo puedes escribir un número que corresponda a una carta existente")
			juego.Dibujar()
		}
		goto a
	} else {
		goto a
	}

	juego.Dibujar()
}
func (juego *Juego) HiloDeTemporizadorDeSeleccion() {
	juego.TiempoDeSeleccion = 10
	juego.TiempoTranscurrido = 0
	juego.TemporizadorDeSeleccion = time.NewTimer(time.Duration(1) * time.Second)
	for segundos := 0; segundos < juego.TiempoDeSeleccion; segundos++ {
		juego.Escribir("Tienes " + ConvetirIntAString(juego.TiempoDeSeleccion-juego.TiempoTranscurrido) + " Segundos para elegir tu primera carta")
		juego.Dibujar()
		<-juego.TemporizadorDeSeleccion.C
		juego.TemporizadorDeSeleccion.Reset(1 * time.Second)
		juego.TiempoTranscurrido += 1
	}
}

//Convertir de string a Int
func (juego *Juego) ConvertirStringAInt(cadena string) (int, error) {
	entero, err := strconv.Atoi(cadena)
	/*if Er != nil {
		juego.Escribir("Solo puedes ingresar Caracteres Numericos")
	}*/
	return entero, err
}

//Convertir arreglo de simbolos a una cadena
func (juego *Juego) SimbolosACadena(simbolos []*Símbolo) string {
	var comando string
	for indice := 0; indice < len(simbolos); indice++ {
		cadena := juego.Fuente.Cadenas[simbolos[indice]]
		//comando = comando + cadena
		comando += cadena
	}
	return comando
}

func (juego *Juego) PreguntarSoloOMultijugador() {
	juego.Escribir("Escoje el modo de juego que quieres jugar ¿Multi-Jugador o de 1 jugador?")
	juego.Escribir("A) Multi-jugador")
	juego.Escribir("B) 1 Jugador")
	juego.Dibujar()
a:
	for {
		comando := juego.EsperarComando()
		cadenaDeComando := juego.SimbolosACadena(comando)
		if strings.EqualFold(cadenaDeComando, "A") {
			juego.JugarMultijugador()
			break a
		} else if strings.EqualFold(cadenaDeComando, "B") {
			juego.JugarUnJugador()
			break a
		} else {
			juego.Escribir("Recuerda escribir una de las Opciones (A,B)")
			juego.Dibujar()
		}
	}

}

func (juego *Juego) HiloLógico() {
	juego.Escribir("Bienvenido a VidaNoVida")
	juego.PreguntarSoloOMultijugador()
}
