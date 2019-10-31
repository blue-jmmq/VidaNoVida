package main

type Carta struct {
	Nombre               string
	Descripción          string
	Costo                int
	Vida                 int
	ArmaduraRoja         int
	ArmaduraAzul         int
	DañoRojo             int
	DañoAzul             int
	Curación             int
	AntiVelocidad        int
	Rango                int
	Nivel                int
	Experiencia          int
	EstadísticasPorNivel *EstadísticasPorNivel
}

func (carta *Carta) ObtenerInterfaz() *InterfazDeCarta {
	var interfaz InterfazDeCarta
	interfaz.Nombre = carta.Nombre
	interfaz.Costo = carta.Costo
	interfaz.Descripción = carta.Descripción
	return &interfaz
}

var ArregloDeConstructoresDeCartas []func() *Carta = []func() *Carta{
	NuevaCartaDeGuerrero,
	NuevaCartaDeNinja,
	NuevaCartaDeMago,
	NuevaCartaDeOgro,
	NuevaCartaDeElfoMago,
	NuevaCartaDeElfoArquero,
	NuevaCartaDeArqueroHumano,
	NuevaCartaDeSacerdote,
	NuevaCartaDeBrujo,
}

//Arreglo de Cartas
var ArregloDeCartas []*Carta = []*Carta{
	NuevaCartaDeGuerrero(),
	NuevaCartaDeNinja(),
	NuevaCartaDeMago(),
	NuevaCartaDeOgro(),
	NuevaCartaDeElfoMago(),
	NuevaCartaDeElfoArquero(),
	NuevaCartaDeArqueroHumano(),
	NuevaCartaDeSacerdote(),
	NuevaCartaDeBrujo(),
}

//NuevaCartaDeGuerrero creates a Warrior Carta
func NuevaCartaDeGuerrero() *Carta {
	return &Carta{
		Nombre:        "Guerrero",
		Descripción:   "Noble caballero de la Edad Media con sus armaduras y armas que se caracteriza por su resistencia y su ataque fisico",
		Costo:         1,
		DañoRojo:      4,
		DañoAzul:      1,
		Curación:      0,
		ArmaduraRoja:  4,
		ArmaduraAzul:  2,
		AntiVelocidad: 4,
		Rango:         1,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 2,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     2,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Ninja Carta
func NuevaCartaDeNinja() *Carta {
	return &Carta{
		Nombre:        "Ninja",
		Descripción:   "Ninja es considerado un mercenario tipo de guerrero japonés contratado para ejercer asesinatos caracterizado por su gran rapidez, su daño fisico y pobre defensa",
		Costo:         1,
		DañoRojo:      2,
		DañoAzul:      1,
		Curación:      0,
		ArmaduraRoja:  1,
		ArmaduraAzul:  1,
		AntiVelocidad: 1,
		Rango:         1,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 2,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     1,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Mage Carta
func NuevaCartaDeMago() *Carta {
	return &Carta{
		Nombre:        "Mago",
		Descripción:   "Considerados por muchos como un hechiceros especializados en la magia y el conosimiento mistico Caracterizados por su daño magico y defensa magica",
		Costo:         1,
		DañoRojo:      1,
		DañoAzul:      4,
		Curación:      0,
		ArmaduraRoja:  1,
		ArmaduraAzul:  4,
		AntiVelocidad: 4,
		Rango:         3,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 2,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     0,
			DañoAzulPorNivel:     3,
		},
	}
}

//NuevaCartaDeGuerrero creates a Ogre Carta
func NuevaCartaDeOgro() *Carta {
	return &Carta{
		Nombre:        "Ogro",
		Descripción:   "Un ogro es el miembro de una raza de humanoides grandes, fieros y crueles que comen carne humana",
		Costo:         1,
		DañoRojo:      2,
		DañoAzul:      0,
		Curación:      0,
		ArmaduraRoja:  8,
		ArmaduraAzul:  2,
		AntiVelocidad: 4,
		Rango:         1,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 4,
			ArmaduraAzulPorNivel: 2,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     1,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Wizard elf Carta
func NuevaCartaDeElfoMago() *Carta {
	return &Carta{
		Nombre:        "Elfo Mago",
		Descripción:   "Misteriosos hasta para los otros miembros de clan elfico, usando su magia para llegar hasta donde los otros elfos no han llegado",
		Costo:         1,
		DañoRojo:      1,
		DañoAzul:      8,
		Curación:      0,
		ArmaduraRoja:  2,
		ArmaduraAzul:  3,
		AntiVelocidad: 3,
		Rango:         8,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 1,
			ArmaduraAzulPorNivel: 2,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     0,
			DañoAzulPorNivel:     2,
		},
	}
}

//NuevaCartaDeGuerrero creates a Archer Elf Carta
func NuevaCartaDeElfoArquero() *Carta {
	return &Carta{
		Nombre:        "Elfo Arquero",
		Descripción:   "Guerros del clan elfico que aprovecha las magia para sus ataques de larga distancia y portar un arco elfico con encantamientos de daño",
		Costo:         1,
		DañoRojo:      6,
		DañoAzul:      1,
		Curación:      0,
		ArmaduraRoja:  1,
		ArmaduraAzul:  1,
		AntiVelocidad: 2,
		Rango:         7,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 1,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     3,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Human archer Carta
func NuevaCartaDeArqueroHumano() *Carta {
	return &Carta{
		Nombre:        "Arquero Humano",
		Descripción:   "Humanos del antiguo clan de los Exiliados que disfrutanban llevar a la locura a sus victimas disparandoles flechzas hasta asesinarlos",
		Costo:         1,
		DañoRojo:      8,
		DañoAzul:      1,
		Curación:      0,
		ArmaduraRoja:  2,
		ArmaduraAzul:  1,
		AntiVelocidad: 3,
		Rango:         6,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 2,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     0,
			DañoRojoPorNivel:     2,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Priest Carta
func NuevaCartaDeSacerdote() *Carta {
	return &Carta{
		Nombre:        "Sacerdote",
		Descripción:   "Los sacerdotes están entregados a lo espiritual sirviendo a la gente con su inquebrantable fe y sus dones místicos dedicados a sanar a sus compañeros en la guerra",
		Costo:         1,
		DañoRojo:      0,
		DañoAzul:      1,
		Curación:      3,
		ArmaduraRoja:  2,
		ArmaduraAzul:  1,
		AntiVelocidad: 4,
		Rango:         4,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 1,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     2,
			DañoRojoPorNivel:     0,
			DañoAzulPorNivel:     0,
		},
	}
}

//NuevaCartaDeGuerrero creates a Warlock Carta
func NuevaCartaDeBrujo() *Carta {
	return &Carta{
		Nombre:        "Brujo",
		Descripción:   "Los brujos son entrenados en las artes oscuras estos letales hechiceros usan su magia para ejercer dominacion sobre sus enemigos",
		Costo:         1,
		DañoRojo:      1,
		DañoAzul:      4,
		Curación:      1,
		ArmaduraRoja:  1,
		ArmaduraAzul:  1,
		AntiVelocidad: 4,
		Rango:         3,
		Nivel:         1,
		Experiencia:   0,
		EstadísticasPorNivel: &EstadísticasPorNivel{
			CostoPorNivel:        1,
			VidaPorNivel:         1,
			ArmaduraRojaPorNivel: 1,
			ArmaduraAzulPorNivel: 1,
			CuraciónPorNivel:     1,
			DañoRojoPorNivel:     0,
			DañoAzulPorNivel:     2,
		},
	}
}
