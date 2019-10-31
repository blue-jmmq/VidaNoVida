package main

type Jugador struct {
	Vida         int
	ArmaduraRoja int
	Credito      int
	Mano         []*Carta
	Tablero      [][]*Carta
}
