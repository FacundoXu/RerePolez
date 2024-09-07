package usuario

import (
	"fmt"
	"rerepolez/errores"
	"rerepolez/tdas_tp"
	"tdas/cola"
)

func FinVotar(fila cola.Cola[tdas_tp.Votante], partidos []tdas_tp.Partido, impugnados *int) {

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

	// Obtenemos el ultimo voto realizado por el votante
	voto := fila.VerPrimero().FinVoto()

	// Cargamos los datos del voto realizado a los candidatos de sus respectivos partidos
	if !voto.Impugnado {

		for i := 0; i < int(tdas_tp.CANT_VOTACION); i++ {
			votado := voto.VotoPorTipo[tdas_tp.TipoVoto(i)]

			if votado != 0 { // Si es un voto valido
				partidos[votado].VotadoPara(tdas_tp.TipoVoto(i))

			} else { // Si es un voto en blanco
				partidos[0].VotadoPara(tdas_tp.TipoVoto(i))
			}
		}

	} else if voto.Impugnado { // Si es un voto impugnado
		*impugnados++
	}

	// El votante termina de votar, lo echamos de la fila
	// A partir de ahora si el mismo votante vuelve a ingresar a la fila, lo echamos por fraude
	fila.Desencolar()
	fmt.Println(OK)
}
