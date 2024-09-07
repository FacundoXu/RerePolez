package usuario

import (
	"fmt"
	"rerepolez/errores"
	"rerepolez/tdas_tp"
	"tdas/cola"
)

func Deshacer(fila cola.Cola[tdas_tp.Votante]) {

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

	// El primero de la fila deshace la ultima votacion realizada
	err := fila.VerPrimero().Deshacer()

	// Error si el votante no tiene votos anteriores
	if err != nil {
		fmt.Println(errores.ErrorNoHayVotosAnteriores{}.Error())
		return
	}

	fmt.Println(OK)
}
