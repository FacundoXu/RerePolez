package usuario

import (
	"fmt"
	"rerepolez/errores"
	"rerepolez/tdas_tp"
	"strconv"
	"tdas/cola"
)

func Votar(cmd []string, partidos []tdas_tp.Partido, fila cola.Cola[tdas_tp.Votante]) {

	// Error si la fila esta vacia
	if fila.EstaVacia() {
		fmt.Println(errores.FilaVacia{}.Error())
		return
	}

	// Error si el primero de la fila es un votante fraudulento, lo echamos de la fila
	if !fila.VerPrimero().EstaVotando() {
		fmt.Println(errores.ErrorVotanteFraudulento{Dni: fila.VerPrimero().LeerDNI()}.Error())
		fila.Desencolar()
		return
	}

	var tipo tdas_tp.TipoVoto

	// Convertimos el parametro ingresado en un TipoVoto
	switch cmd[1] {

	case "Presidente":
		tipo = tdas_tp.PRESIDENTE

	case "Gobernador":
		tipo = tdas_tp.GOBERNADOR

	case "Intendente":
		tipo = tdas_tp.INTENDENTE

	default:
		tipo = -1
	}

	// Error si el parametro ingresado no es un tipo valido
	if tipo == -1 {
		fmt.Println(errores.ErrorTipoVoto{}.Error())
		return
	}

	alternativa, err := strconv.Atoi(cmd[2])

	// Error si es la alternativa es invalida
	if alternativa < 0 || alternativa > len(partidos)-1 || err != nil {
		fmt.Println(errores.ErrorAlternativaInvalida{}.Error())
		return
	}

	// El primero de la fila realiza sus votaciones
	fila.VerPrimero().Votar(tipo, alternativa)
	fmt.Println(OK)
}
