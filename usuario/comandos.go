package usuario

import (
	"fmt"
	"rerepolez/tdas_tp"
	"tdas/cola"
)

const (
	OK        = "OK"
	INGRESAR  = "ingresar"
	VOTAR     = "votar"
	DESHACER  = "deshacer"
	FIN_VOTAR = "fin-votar"

	ERROR_ENTRADA = "ERROR: Comando incorrecto"
)

func Comandos(cmd []string, partidos []tdas_tp.Partido, votantes []tdas_tp.Votante, fila cola.Cola[tdas_tp.Votante], impugnados *int) {

	// Condiciones de programa dependiendo de la entrada ingresada
	switch cmd[0] {

	case INGRESAR:
		Ingresar(cmd, votantes, fila)

	case VOTAR:
		Votar(cmd, partidos, fila)

	case DESHACER:
		Deshacer(fila)

	case FIN_VOTAR:
		FinVotar(fila, partidos, impugnados)

	default:
		fmt.Println(ERROR_ENTRADA) // Error cuando ingresa un comando invalido
	}
}

func MostrarResultados(partidos []tdas_tp.Partido, impugnados *int) {
	fmt.Println("Presidente:")
	for _, partido := range partidos {
		fmt.Println(partido.ObtenerResultado(tdas_tp.PRESIDENTE))
	}
	fmt.Print("\n")

	fmt.Println("Gobernador:")
	for _, partido := range partidos {
		fmt.Println(partido.ObtenerResultado(tdas_tp.GOBERNADOR))
	}
	fmt.Print("\n")

	fmt.Println("Intendente:")
	for _, partido := range partidos {
		fmt.Println(partido.ObtenerResultado(tdas_tp.INTENDENTE))
	}
	fmt.Print("\n")

	if *impugnados == 1 {
		fmt.Println("Votos Impugnados:", *impugnados, "voto")

	} else {
		fmt.Println("Votos Impugnados:", *impugnados, "votos")
	}
}
