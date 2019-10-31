package main

import (
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Implementación struct {
	Juego   *Juego
	Pixeles []byte
	Ventana *glfw.Window
	búfer   uint32
}

func CrearImplementación(juego *Juego) *Implementación {
	var implementación Implementación
	implementación.Juego = juego
	return &implementación
}
func (juego *Juego) EnviarComando() {
	select {
	case juego.CanalDeComandos <- juego.IU.WidgetDeEntrada.Búfer:
		if len(juego.IU.WidgetDeEntrada.Búfer) > 0 {
			juego.IU.WidgetDeEntrada.Búfer = make([]*Símbolo, 0)
			juego.IU.WidgetDeEntrada.Índice = 0
		}
	default:
	}
}

func (implementación *Implementación) LlamadaDeTecla(ventana *glfw.Window, tecla glfw.Key, scancode int, acción glfw.Action, mods glfw.ModifierKey) {
	if tecla == glfw.KeyUp && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeSalida.DesplazarseHaciaArriba(implementación.Juego.Fuente)
		implementación.Juego.Dibujar()
	}
	if tecla == glfw.KeyDown && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeSalida.DesplazarseHaciaAbajo(implementación.Juego.Fuente)
		implementación.Juego.Dibujar()
	}
	if tecla == glfw.KeyLeft && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeEntrada.DesplazarCursor(-1)
		implementación.Juego.IU.WidgetDeEntrada.RestablecerCursor()

		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyRight && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeEntrada.DesplazarCursor(1)
		implementación.Juego.IU.WidgetDeEntrada.RestablecerCursor()

		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyBackspace && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeEntrada.EliminarSímbolo()
		implementación.Juego.IU.WidgetDeEntrada.RestablecerCursor()
		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyEnter && acción == glfw.Press {
		implementación.Juego.EnviarComando()
		implementación.Juego.IU.WidgetDeEntrada.RestablecerCursor()
		implementación.Juego.Tick()
	}
	//Imprimir("tecla", tecla, glfw.KeyUp, glfw.GetKeyName(tecla, scancode), scancode)
}

func (implementación *Implementación) LlamadaDeTexto(ventana *glfw.Window, runa rune) {
	símbolo := implementación.Juego.RunaASímbolo(runa)
	implementación.Juego.IU.WidgetDeEntrada.Escribir(símbolo)
	implementación.Juego.IU.WidgetDeEntrada.RestablecerCursor()
	implementación.Juego.Tick()
	//Imprimir("Runa:", string(runa))
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

	ventana, err := glfw.CreateWindow(implementación.Juego.Anchura, implementación.Juego.Altura, "VidaNoVida", nil, nil)
	if err != nil {
		panic(err)
	}
	implementación.Ventana = ventana
	ventana.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	Imprimir("OpenGL version", version)

	implementación.Pixeles = make([]byte, implementación.Juego.Anchura*implementación.Juego.Altura*4)
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	gl.GenFramebuffers(1, &implementación.búfer)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, implementación.búfer)
	gl.FramebufferTexture2D(gl.READ_FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, 0)

	//gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	/*
		gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(implementación.Pixeles))
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, implementación.búfer)
		gl.BlitFramebuffer(0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
			0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
			gl.COLOR_BUFFER_BIT, gl.LINEAR)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, 0)
		window.SwapBuffers()
	*/
	//implementación.Juego.Dibujar()
	ventana.SetKeyCallback(implementación.LlamadaDeTecla)
	ventana.SetCharCallback(implementación.LlamadaDeTexto)
	implementación.Juego.TemporizadorDelCursor = time.NewTimer(512 * time.Millisecond)
	go implementación.Juego.BucleDelTemporizador()
	go implementación.Juego.HiloLógico()
	for !ventana.ShouldClose() {
		//var frameBuffer uint32
		//gl.GenFramebuffers(1, &frameBuffer)
		// Do OpenGL stuff.
		//
		implementación.Dibujar()

		glfw.PollEvents()
	}
}

func (implementación *Implementación) Dibujar() {
	if implementación.Juego.HayQueDibujar {
		for y := 0; y < implementación.Juego.Altura; y++ {
			for x := 0; x < implementación.Juego.Anchura; x++ {
				píxel := implementación.Juego.Pixeles.Leer(y, x).(*Píxel)
				color := píxel.Color
				rojo, verde, azul := color.Rojo, color.Verde, color.Azul
				fila := implementación.Juego.Altura - y - 1
				implementación.Pixeles[fila*implementación.Juego.Anchura*4+x*4+0] = byte(rojo)
				implementación.Pixeles[fila*implementación.Juego.Anchura*4+x*4+1] = byte(verde)
				implementación.Pixeles[fila*implementación.Juego.Anchura*4+x*4+2] = byte(azul)
				implementación.Pixeles[fila*implementación.Juego.Anchura*4+x*4+3] = 255
			}
		}
		gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(implementación.Pixeles))
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, implementación.búfer)
		gl.BlitFramebuffer(0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
			0, 0, int32(implementación.Juego.Anchura), int32(implementación.Juego.Altura),
			gl.COLOR_BUFFER_BIT, gl.LINEAR)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, 0)
		implementación.Ventana.SwapBuffers()
		implementación.Juego.HayQueDibujar = false
	}
}
