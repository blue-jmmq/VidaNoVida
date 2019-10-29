package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type StatisticsPerLevel struct {
	CostPerLevel       int
	HealthPerLevel     int
	RedArmorPerLever   int
	BlueArmorPerLevel  int
	RedDamagePerLever  int
	BlueDamagePerLever int
	HealingPerLever    int
}

//Card structure
type Card struct {
	Name               string
	Descripción        string
	Cost               int
	Health             int
	RedArmor           int
	BlueArmor          int
	RedDamage          int
	BlueDamage         int
	Healing            int
	AntiAttackSpeed    int
	Range              int
	Level              int
	Experience         int
	StatisticsPerLevel *StatisticsPerLevel
}

//NewWarriorCard creates a Warrior Card
func NewWarriorCard() *Card {
	return &Card{
		Name:            "Guerrero",
		Descripción:     "Noble caballero de la Edad Media con sus armaduras y armas que se caracteriza por su resistencia y su ataque fisico",
		Cost:            1,
		RedDamage:       4,
		BlueDamage:      1,
		Healing:         0,
		RedArmor:        4,
		BlueArmor:       2,
		AntiAttackSpeed: 4,
		Range:           1,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   2,
			BlueArmorPerLevel:  1,
			HealingPerLever:    0,
			RedDamagePerLever:  2,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Ninja Card
func NewNinjaCard() *Card {
	return &Card{
		Name:            "Ninja",
		Descripción:     "Ninja es considerado un mercenario tipo de guerrero japonés contratado para ejercer asesinatos caracterizado por su gran rapidez, su daño fisico y pobre defensa",
		Cost:            1,
		RedDamage:       2,
		BlueDamage:      1,
		Healing:         0,
		RedArmor:        1,
		BlueArmor:       1,
		AntiAttackSpeed: 1,
		Range:           1,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   2,
			BlueArmorPerLevel:  1,
			HealingPerLever:    0,
			RedDamagePerLever:  1,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Mage Card
func NewMageCard() *Card {
	return &Card{
		Name:            "Mago",
		Descripción:     "Considerados por muchos como un hechiceros especializados en la magia y el conosimiento mistico Caracterizados por su daño magico y defensa magica",
		Cost:            1,
		RedDamage:       1,
		BlueDamage:      4,
		Healing:         0,
		RedArmor:        1,
		BlueArmor:       4,
		AntiAttackSpeed: 4,
		Range:           3,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   2,
			BlueArmorPerLevel:  1,
			HealingPerLever:    0,
			RedDamagePerLever:  0,
			BlueDamagePerLever: 3,
		},
	}
}

//NewWarriorCard creates a Ogre Card
func NewOgreCard() *Card {
	return &Card{
		Name:            "Ogro",
		Descripción:     "Un ogro es el miembro de una raza de humanoides grandes, fieros y crueles que comen carne humana",
		Cost:            1,
		RedDamage:       2,
		BlueDamage:      0,
		Healing:         0,
		RedArmor:        8,
		BlueArmor:       2,
		AntiAttackSpeed: 4,
		Range:           1,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   4,
			BlueArmorPerLevel:  2,
			HealingPerLever:    0,
			RedDamagePerLever:  1,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Wizard elf Card
func NewWizardElfCard() *Card {
	return &Card{
		Name:            "Elfo Mago",
		Descripción:     "Misteriosos hasta para los otros miembros de clan elfico, usando su magia para llegar hasta donde los otros elfos no han llegado",
		Cost:            1,
		RedDamage:       1,
		BlueDamage:      8,
		Healing:         0,
		RedArmor:        2,
		BlueArmor:       3,
		AntiAttackSpeed: 3,
		Range:           8,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   1,
			BlueArmorPerLevel:  2,
			HealingPerLever:    0,
			RedDamagePerLever:  0,
			BlueDamagePerLever: 2,
		},
	}
}

//NewWarriorCard creates a Archer Elf Card
func CrearCartaDeElfoArquero() *Card {
	return &Card{
		Name:            "Elfo Arquero",
		Descripción:     "Guerros del clan elfico que aprovecha las magia para sus ataques de larga distancia y portar un arco elfico con encantamientos de daño",
		Cost:            1,
		RedDamage:       6,
		BlueDamage:      1,
		Healing:         0,
		RedArmor:        1,
		BlueArmor:       1,
		AntiAttackSpeed: 2,
		Range:           7,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   1,
			BlueArmorPerLevel:  1,
			HealingPerLever:    0,
			RedDamagePerLever:  3,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Human archer Card
func NewHumanArcherCard() *Card {
	return &Card{
		Name:            "Arquero Aumano",
		Descripción:     "Humanos del antiguo clan de los Exiliados que disfrutanban llevar a la locura a sus victimas disparandoles flechzas hasta asesinarlos",
		Cost:            1,
		RedDamage:       8,
		BlueDamage:      1,
		Healing:         0,
		RedArmor:        2,
		BlueArmor:       1,
		AntiAttackSpeed: 3,
		Range:           6,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   2,
			BlueArmorPerLevel:  1,
			HealingPerLever:    0,
			RedDamagePerLever:  2,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Priest Card
func NewPriestCard() *Card {
	return &Card{
		Name:            "Sacerdote",
		Descripción:     "Los sacerdotes están entregados a lo espiritual sirviendo a la gente con su inquebrantable fe y sus dones místicos dedicados a sanar a sus compañeros en la guerra",
		Cost:            1,
		RedDamage:       0,
		BlueDamage:      1,
		Healing:         3,
		RedArmor:        2,
		BlueArmor:       1,
		AntiAttackSpeed: 4,
		Range:           4,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   1,
			BlueArmorPerLevel:  1,
			HealingPerLever:    2,
			RedDamagePerLever:  0,
			BlueDamagePerLever: 0,
		},
	}
}

//NewWarriorCard creates a Warlock Card
func NewWarlockCard() *Card {
	return &Card{
		Name:            "Brujo",
		Descripción:     "Los brujos son entrenados en las artes oscuras estos letales hechiceros usan su magia para ejercer dominacion sobre sus enemigos",
		Cost:            1,
		RedDamage:       1,
		BlueDamage:      4,
		Healing:         1,
		RedArmor:        1,
		BlueArmor:       1,
		AntiAttackSpeed: 4,
		Range:           3,
		Level:           1,
		Experience:      0,
		StatisticsPerLevel: &StatisticsPerLevel{
			CostPerLevel:       1,
			HealthPerLevel:     1,
			RedArmorPerLever:   1,
			BlueArmorPerLevel:  1,
			HealingPerLever:    1,
			RedDamagePerLever:  0,
			BlueDamagePerLever: 2,
		},
	}
}

//Interfaz de Cartas
type InterfazDeCarta struct {
	Nombre      string
	Costo       int
	Descripción string
}

//strings.EqualFold
func (juego *Juego) LeerComando() []*Símbolo {
	command := <-juego.CanalDeComandos
	return command
}
func (juego *Juego) JugarMultijugador() {

}
func (juego *Juego) JugarUnJugador() {
	juego.Escribir("Escoje cuatro cartas para tu mano inicial")
	juego.Escribir("Presiona la tecla ENTER para ver las cartas disponibles")
	juego.Dibujar()
	juego.LeerComando()
	for i := 0; i < len(ArregloDeCartas); i++ {
		juego.Escribir(strconv.Itoa(i+1) + ")\n" + JSONIdentado(ArregloDeCartas[i].ObtenerInterfaz()))
	}
	juego.Dibujar()
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
		comando := juego.LeerComando()
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
	//juego.Escribir("Bienvenido a VidaEsVida")
	//juego.Escribir("Bienvenido a VidaEsVida")
	//juego.Escribir("aaaaaaaaaaaaaaaaaaaaaaa")
	//juego.Escribir("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	juego.Escribir("Bienvenido a VidaNoVida")
	//juego.Dibujar()
	//Imprimir("juego.Escribir(\"Bienvenido a VidaEsVida\"")
	juego.PreguntarSoloOMultijugador()
}

func (carta *Card) ObtenerInterfaz() *InterfazDeCarta {
	var interfaz InterfazDeCarta
	interfaz.Nombre = carta.Name
	interfaz.Costo = carta.Cost
	interfaz.Descripción = carta.Descripción
	return &interfaz
}

//Arreglo de Cartas
var ArregloDeCartas []*Card = []*Card{
	NewWarriorCard(),
	NewNinjaCard(),
	NewMageCard(),
	NewOgreCard(),
	NewWizardElfCard(),
	CrearCartaDeElfoArquero(),
	NewHumanArcherCard(),
	NewPriestCard(),
	NewWarlockCard(),
}

//Player structure
type Player struct {
	Health   int
	RedArmor int
	Energy   int
	Credit   int
	Deck     []Card
	Hand     []Card
	Board    [][]Card
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

func (widget *WidgetDeEntrada) EliminarSímbolo() {
	widget.DeboDibujarCursor = true
	if widget.Índice > 0 {
		var provicional []*Símbolo
		provicional = append(provicional, widget.Búfer[0:widget.Índice-1]...)
		provicional = append(provicional, widget.Búfer[widget.Índice:]...)
		widget.Búfer = provicional
		widget.DesplazarCursor(-1)
	}
}

func (widget *WidgetDeEntrada) Escribir(símbolo *Símbolo) {
	widget.DeboDibujarCursor = true
	var provicional []*Símbolo
	provicional = append(provicional, widget.Búfer[:widget.Índice]...)
	provicional = append(provicional, símbolo)
	provicional = append(provicional, widget.Búfer[widget.Índice:]...)
	widget.Búfer = provicional
	widget.DesplazarCursor(1)
}

func (widget *WidgetDeEntrada) DesplazarCursor(delta int) {
	widget.DeboDibujarCursor = true
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

/*
func (widget *WidgetDeEntrada) ActualizarSímbolos(fuente *Fuente) {
	tamañoDelBúfer := len(widget.Búfer)
	columnaInicial := 3
	var índice int
	Imprimir("widget.Índice", widget.Índice)
	if tamañoDelBúfer < widget.TamañoVisual {
		for columna := columnaInicial; columna < tamañoDelBúfer+columnaInicial; columna++ {
			índice = columna - columnaInicial
			if índice == widget.Índice {
				widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
			} else {
				widget.Símbolos.Escribir(1, columna, widget.Búfer[índice])
			}
		}
		for columna := tamañoDelBúfer + columnaInicial; columna < widget.TamañoVisual+columnaInicial; columna++ {
			índice = columna - columnaInicial
			if índice == widget.Índice {
				widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
			} else {
				widget.Símbolos.Escribir(1, columna, fuente.Símbolos[" "])
			}
		}
	} else if tamañoDelBúfer == widget.TamañoVisual {
		for columna := columnaInicial; columna < widget.TamañoVisual+columnaInicial; columna++ {
			índice = columna - columnaInicial
			if índice == widget.Índice {
				widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
			} else {
				widget.Símbolos.Escribir(1, columna, widget.Búfer[índice])
			}
		}
		índice = widget.TamañoVisual
		if índice == widget.Índice {
			widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
		}
	} else {
		Imprimir("tamañoDelBúfer", tamañoDelBúfer)
		Imprimir("widget.TamañoVisual", widget.TamañoVisual)
		Imprimir("tamañoDelBúfer - widget.TamañoVisual", tamañoDelBúfer-widget.TamañoVisual)
		for columna := columnaInicial; columna < widget.TamañoVisual+columnaInicial; columna++ {
			índice = tamañoDelBúfer - widget.TamañoVisual + columna - columnaInicial
			if índice == widget.Índice {
				widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
			} else {
				widget.Símbolos.Escribir(1, columna, widget.Búfer[índice])
			}
		}
		índice = widget.TamañoVisual
		if índice == widget.Índice {
			widget.Símbolos.Escribir(1, columna, fuente.Símbolos["bloque"])
		}
	}

}*/

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

type IU struct {
	WidgetDeSalida  *WidgetDeSalida
	WidgetDeEntrada *WidgetDeEntrada
	Anchura         int
	Altura          int
	Símbolos        *Bidimensional
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
	Color *Color
}

type Color struct {
	Rojo  byte
	Verde byte
	Azul  byte
}

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
	juego.CanalDeComandos <- juego.IU.WidgetDeEntrada.Búfer
	if len(juego.IU.WidgetDeEntrada.Búfer) > 0 {
		juego.IU.WidgetDeEntrada.Búfer = make([]*Símbolo, 0)
		juego.IU.WidgetDeEntrada.Índice = 0
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
		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyRight && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeEntrada.DesplazarCursor(1)
		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyBackspace && ((acción == glfw.Press) || (acción == glfw.Repeat)) {
		implementación.Juego.IU.WidgetDeEntrada.EliminarSímbolo()
		implementación.Juego.Tick()
	}
	if tecla == glfw.KeyEnter && acción == glfw.Press {
		implementación.Juego.EnviarComando()
	}
	//Imprimir("tecla", tecla, glfw.KeyUp, glfw.GetKeyName(tecla, scancode), scancode)
}

func (implementación *Implementación) LlamadaDeTexto(ventana *glfw.Window, runa rune) {
	símbolo := implementación.Juego.RunaASímbolo(runa)
	implementación.Juego.IU.WidgetDeEntrada.Escribir(símbolo)
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
	fmt.Println("OpenGL version", version)

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

type Fuente struct {
	Nombre   string
	Símbolos map[string]*Símbolo
	Imágenes map[string]string
	Cadenas  map[*Símbolo]string
}

func CrearFuente(nombre string) *Fuente {
	fuente := new(Fuente)
	fuente.Nombre = nombre
	fuente.LeerImágenes()
	fuente.LeerSímbolos()
	fuente.CrearCadenas()
	//Imprimir("Runas:", fuente.Runas)
	//Imprimir("Símbolos:", fuente.Símbolos)

	return fuente
}
func (fuente *Fuente) CrearCadenas() {
	fuente.Cadenas = make(map[*Símbolo]string)
	for nombre, simbolo := range fuente.Símbolos {
		fuente.Cadenas[simbolo] = nombre
	}
}

/*
func (fuente *Fuente) ContieneRuna(runa rune) bool {
	return false
}
*/

func (fuente *Fuente) LeerSímbolos() {
	fuente.Símbolos = make(map[string]*Símbolo)
	for nombre, imagen := range fuente.Imágenes {
		fuente.Símbolos[nombre] = CrearSímboloDesdeImagen(fuente.Nombre, imagen)
	}
}

func (fuente *Fuente) LeerImágenes() {
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
	fuente.Imágenes = make(map[string]string)
	for _, dupla := range runas {
		dupla := dupla.([]interface{})
		nombre := dupla[0].(string)
		imagen := dupla[1].(string)
		//Imprimir(reflect.TypeOf(dupla[0]))
		fuente.Imágenes[nombre] = imagen
	}
}

type Juego struct {
	IU                    *IU
	Pantalla              *Pantalla
	Fuente                *Fuente
	Pixeles               *Bidimensional
	PseudoPixeles         *Bidimensional
	PseudoTamaño          int
	TamañoDeSímbolo       int
	SímboloVacío          *Símbolo
	Colores               []*Color
	Implementación        *Implementación
	Altura                int
	Anchura               int
	HayQueDibujar         bool
	TemporizadorDelCursor *time.Timer
	CanalDeComandos       chan []*Símbolo
}

//CrearDatosDelJuego es una función
func CrearJuego() *Juego {
	juego := new(Juego)
	juego.Anchura = 1024
	juego.Altura = 512
	juego.PseudoTamaño = 1
	juego.TamañoDeSímbolo = 16
	juego.HayQueDibujar = true
	juego.CanalDeComandos = make(chan []*Símbolo)

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
	juego.Colores = append(juego.Colores, &Color{Rojo: 0, Verde: 0, Azul: 0})
	juego.Colores = append(juego.Colores, &Color{Rojo: 255, Verde: 255, Azul: 0})
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
