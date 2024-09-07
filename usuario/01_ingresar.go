package usuario

import (
	"fmt"
	"rerepolez/auxiliares"
	"rerepolez/errores"
	"rerepolez/tdas_tp"
	"strconv"
	"tdas/cola"
)

func Ingresar(cmd []string, votantes []tdas_tp.Votante, fila cola.Cola[tdas_tp.Votante]) {

	// Convertimos el padron ingresado a un entero
	padron, err := strconv.Atoi(cmd[1])

	// Error si el padron es incorrecto
	if err != nil || padron <= 0 {
		fmt.Println(errores.DNIError{}.Error())
		return
	}

	// Chequeamos que el padron esta dentro de los padrones validos
	esValido, pos := auxiliares.PadronValido(padron, votantes)

	// Error si el padron ingresado no está en el padron de votantes
	if !esValido {
		fmt.Println(errores.DNIFueraPadron{}.Error())
		return
	}

	// El padron es correcto, entonces entra en la fila de votantes
	fila.Encolar(votantes[pos])
	fmt.Println(OK)
}
