package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"runtime"

	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

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

type Toro struct {
	Bidimensional *Bidimensional
}

func CrearToro(anchura, altura int, valorPorDefecto interface{}) *Toro {
	var toro Toro
	toro.Bidimensional = CrearBidimensional(anchura, altura, valorPorDefecto)
	return &toro
}

func (toro *Toro) Limitar(valor, inferior, superior int) int {
	diferencia := superior - inferior
	for valor < inferior {
		valor += diferencia
	}
	for valor >= superior {
		valor -= diferencia
	}
	return valor
}

func (toro *Toro) ConvertirIndice(fila, columna int) (int, int) {
	fila = toro.Limitar(fila, 0, toro.LeerAltura())
	columna = toro.Limitar(columna, 0, toro.LeerAnchura())
	return fila, columna
}

func (toro *Toro) Leer(fila, columna int) interface{} {
	fila, columna = toro.ConvertirIndice(fila, columna)
	return toro.Bidimensional.Leer(fila, columna)
}

func (toro *Toro) Escribir(fila, columna int, valor interface{}) {
	fila, columna = toro.ConvertirIndice(fila, columna)
	toro.Bidimensional.Escribir(fila, columna, valor)
}

func (toro *Toro) LeerAltura() int {
	return toro.Bidimensional.LeerAltura()
}

func (toro *Toro) LeerAnchura() int {
	return toro.Bidimensional.LeerAnchura()
}

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

type Celda struct {
	Símbolo *Símbolo
}

type Píxel struct {
	Color Color
}

type Color struct {
	Rojo  byte
	Verde byte
	Azul  byte
}

type Implementación struct {
	Juego   *Juego
	Pixeles []byte
}

func CrearImplementación(juego *Juego) *Implementación {
	var implementación Implementación
	implementación.Juego = juego
	return &implementación
}

func (implementación *Implementación) Correr() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(implementación.Juego.Anchura, implementación.Juego.Altura, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	implementación.Pixeles = make([]byte, implementación.Juego.Anchura*implementación.Juego.Altura*4)
	implementación.Juego.Dibujar()
	implementación.Dibujar()
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	var frameBuffer uint32
	gl.GenFramebuffers(1, &frameBuffer)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, frameBuffer)
	gl.FramebufferTexture2D(gl.READ_FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, 0)

	//gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(implementación.Pixeles))
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, frameBuffer)
	gl.BlitFramebuffer(0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
		0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
		gl.COLOR_BUFFER_BIT, gl.LINEAR)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, 0)
	window.SwapBuffers()
	for !window.ShouldClose() {
		//var frameBuffer uint32
		//gl.GenFramebuffers(1, &frameBuffer)
		// Do OpenGL stuff.
		//

		glfw.PollEvents()
	}
}

func (implementación *Implementación) Dibujar() {
	for y := 0; y < implementación.Juego.Altura; y++ {
		for x := 0; x < implementación.Juego.Anchura; x++ {
			píxel := implementación.Juego.Pixeles.Leer(y, x).(Píxel)
			clr := píxel.Color
			rojo, verde, azul := clr.Rojo, clr.Verde, clr.Azul
			implementación.Pixeles[y*implementación.Juego.Anchura*4+x*4+0] = byte(rojo)
			implementación.Pixeles[y*implementación.Juego.Anchura*4+x*4+1] = byte(verde)
			implementación.Pixeles[y*implementación.Juego.Anchura*4+x*4+2] = byte(azul)
			implementación.Pixeles[y*implementación.Juego.Anchura*4+x*4+3] = 255
		}
	}
}

type Fuente struct {
	Nombre   string
	Símbolos map[rune]*Símbolo
	Runas    map[rune]string
}

func CrearFuente(nombre string) *Fuente {
	fuente := new(Fuente)
	fuente.Nombre = nombre
	fuente.LeerRunas()
	fuente.LeerSímbolos()
	//Imprimir("Runas:", fuente.Runas)
	//Imprimir("Símbolos:", fuente.Símbolos)

	return fuente
}

/*
func (fuente *Fuente) ContieneRuna(runa rune) bool {
	return false
}
*/

func (fuente *Fuente) LeerSímbolos() {
	fuente.Símbolos = make(map[rune]*Símbolo)
	for runa, imagen := range fuente.Runas {
		fuente.Símbolos[runa] = CrearSímboloDesdeImagen(fuente.Nombre, imagen)
	}
}

func (fuente *Fuente) LeerRunas() {
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
	fuente.Runas = make(map[rune]string)
	for _, dupla := range runas {
		dupla := dupla.([]interface{})
		cadenaDeRuna := dupla[0].(string)
		runa := []rune(cadenaDeRuna)[0]
		imagen := dupla[1].(string)
		//Imprimir(reflect.TypeOf(dupla[0]))
		fuente.Runas[runa] = imagen
	}
}

type Juego struct {
	Pantalla        *Pantalla
	Fuente          *Fuente
	Pixeles         *Bidimensional
	PseudoPixeles   *Bidimensional
	PseudoTamaño    int
	TamañoDeSímbolo int
	SímboloVacío    *Símbolo
	Colores         []Color
	Implementación  *Implementación
	Altura          int
	Anchura         int
}

//CrearDatosDelJuego es una función
func CrearJuego() *Juego {
	var juego Juego
	juego.Anchura = 1024
	juego.Altura = 512
	juego.PseudoTamaño = 2
	juego.TamañoDeSímbolo = 16
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
	juego.SímboloVacío = CrearSímboloDesdeImagen(juego.Fuente.Nombre, "8")
	pantalla := CrearPantalla(
		juego.Anchura/juego.TamañoDeSímbolo/juego.PseudoTamaño,
		juego.Altura/juego.TamañoDeSímbolo/juego.PseudoTamaño,
		juego.SímboloVacío,
	)
	juego.Pantalla = pantalla
	juego.Implementación = CrearImplementación(&juego)
	juego.Colores = append(juego.Colores, Color{Rojo: 255, Verde: 255, Azul: 255})
	juego.Colores = append(juego.Colores, Color{Rojo: 0, Verde: 0, Azul: 0})
	juego.Pixeles = CrearBidimensional(juego.Anchura, juego.Altura, Píxel{Color: Color{Rojo: 0, Verde: 0, Azul: 0}})
	juego.PseudoPixeles = CrearBidimensional(
		juego.Anchura/juego.PseudoTamaño,
		juego.Altura/juego.PseudoTamaño,
		Píxel{Color: Color{Rojo: 0, Verde: 0, Azul: 0}},
	)
	return &juego
}

func (juego *Juego) DibujarPseudoPíxel(píxel Píxel, fila, columna int) {
	yInicial := fila * juego.PseudoTamaño
	xInicial := columna * juego.PseudoTamaño
	yFinal := yInicial + juego.PseudoTamaño
	xFinal := xInicial + juego.PseudoTamaño
	for y := yInicial; y < yFinal; y++ {
		for x := xInicial; x < xFinal; x++ {
			juego.Pixeles.Escribir(y, x, píxel)
		}
	}
}

func (juego *Juego) DibujarSímbolo(símbolo *Símbolo, fila, columna int) {
	yInicial := fila * juego.TamañoDeSímbolo
	xInicial := columna * juego.TamañoDeSímbolo
	yFinal := yInicial + juego.TamañoDeSímbolo
	xFinal := xInicial + juego.TamañoDeSímbolo
	for y := yInicial; y < yFinal; y++ {
		for x := xInicial; x < xFinal; x++ {
			colorIndex := símbolo.Arreglo.Leer(y-yInicial, x-xInicial).(byte)
			color := juego.Colores[colorIndex]
			píxel := Píxel{Color: color}
			juego.PseudoPixeles.Escribir(y, x, píxel)
			juego.DibujarPseudoPíxel(píxel, y, x)
		}
	}
}

func (juego *Juego) DibujarCelda(fila, columna int) {
	interfaz := juego.Pantalla.Celdas.Leer(fila, columna)
	celda := interfaz.(*Celda)
	juego.DibujarSímbolo(celda.Símbolo, fila, columna)
}

func (juego *Juego) Dibujar() {
	for fila := 0; fila < juego.Pantalla.LeerAltura(); fila++ {
		for columna := 0; columna < juego.Pantalla.LeerAnchura(); columna++ {
			juego.DibujarCelda(fila, columna)
		}
	}
}

func (juego *Juego) Jugar() {
	juego.Implementación.Correr()
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
