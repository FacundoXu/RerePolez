package auxiliares

import "rerepolez/tdas_tp"

// PadronValido recibe por parametro un padron y un arreglo de votantes ordenado.
// Devuelve un booleano y un entero, (true, posicion) en caso de que encuentre
// el padron pasado por parametro en el arreglo ordenado, (false, -1) en caso contrario.
func padronValido(padron int, padrones []tdas_tp.Votante, inicio int, fin int) (bool, int) {
	if inicio > fin {
		return false, -1
	}

	medio := (inicio + fin) / 2

	if padrones[medio].LeerDNI() == padron {
		return true, medio

	} else if padron > padrones[medio].LeerDNI() {
		return padronValido(padron, padrones, medio+1, fin)

	} else {
		return padronValido(padron, padrones, inicio, medio-1)
	}
}

func PadronValido(padron int, padrones []tdas_tp.Votante) (bool, int) {
	return padronValido(padron, padrones, 0, len(padrones)-1)
}
