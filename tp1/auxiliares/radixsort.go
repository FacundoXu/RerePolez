package auxiliares

// Busca el numero mas grande de un arreglo
func buscarNumMasGrande(padrones []int) int {
	numMasGrande := 0

	for i := 0; i < len(padrones); i++ {
		if padrones[i] > numMasGrande {
			numMasGrande = padrones[i]
		}
	}
	return numMasGrande
}

// Ordenamiento no comparativo
func RadixSort(padrones []int) []int {
	numMasGrande := buscarNumMasGrande(padrones)
	largo := len(padrones)
	digitoSignificativo := 1
	semiOrdenado := make([]int, largo)

	// Hacemos un ciclo hasta encontrar el numero mas significativo
	for numMasGrande/digitoSignificativo > 0 {
		bucket := [10]int{0}

		// Cuenta el número de "claves" o dígitos que entrarán en cada bucket
		for i := 0; i < largo; i++ {
			bucket[(padrones[i]/digitoSignificativo)%10]++
		}

		// Agrega el conteo de los buckets anteriores
		// Consigue los índices luego de cada posicion de bucket ubicado en el arreglo
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// Usamos el bucket para completar un arreglo semiordenado
		for i := largo - 1; i >= 0; i-- {
			bucket[(padrones[i]/digitoSignificativo)%10]--
			semiOrdenado[bucket[(padrones[i]/digitoSignificativo)%10]] = padrones[i]
		}

		// Reemplaza el arreglo actual con el arreglo semiordenado
		for i := 0; i < largo; i++ {
			padrones[i] = semiOrdenado[i]
		}

		// Cambia al proximo digito significativo
		digitoSignificativo *= 10
	}
	return padrones
}
