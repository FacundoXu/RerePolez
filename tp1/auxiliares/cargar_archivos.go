package auxiliares

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errores"
	"rerepolez/tdas_tp"
	"strconv"
	"strings"
)

// CargarPadrones recibe por parametro la ruta que contiene los datos de los padrones.
// Devuelve un arreglo de votantes con sus respectivos padrones
func CargarPadrones(ruta string) []tdas_tp.Votante {
	padrones := []int{}
	votantes := []tdas_tp.Votante{}

	// Abrimos el archivo
	archivo, err := os.Open(ruta)

	// Devuelve un error en caso de que no exista el archivo
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		os.Exit(0)
	}

	// Cierra el archivo al final de la ejecucion de la funcion
	defer archivo.Close()

	// Leemos cada linea del archivo y guardamos los padrones
	linea := bufio.NewScanner(archivo)

	for linea.Scan() {
		padron, _ := strconv.Atoi(linea.Text())
		padrones = append(padrones, padron)
	}

	// Utilizamos RadixSort en el arreglo de padrones
	padronesOrdenados := RadixSort(padrones)

	// Por cada padron del arreglo creamos un votante
	for _, padron := range padronesOrdenados {
		votantes = append(votantes, tdas_tp.CrearVotante(padron))
	}

	return votantes
}

// CargarPartidos recibe por parametro la ruta que contiene los datos de los partidos
// Devuelve un arreglo de partidos con sus respectivos candidatos
func CargarPartidos(ruta string) []tdas_tp.Partido {
	partidos := []tdas_tp.Partido{}

	// Abrimos el archivo
	archivo, err := os.Open(ruta)

	// Devuelve un error en caso de que no exista el archivo
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		os.Exit(0)
	}

	// Cierra el archivo al final de la ejecucion de la funcion
	defer archivo.Close()

	// Agregamos un partido en blanco en la primera posicion del arreglo
	partidoEnBLanco := tdas_tp.CrearVotosEnBlanco()
	partidos = append(partidos, partidoEnBLanco)

	// Leemos cada linea del archivo y guardamos los partidos
	linea := bufio.NewScanner(archivo)

	for linea.Scan() {
		linea := strings.Split(linea.Text(), ",")
		nombre, presidente, gobernador, intendente := linea[0], linea[1], linea[2], linea[3]
		partido := tdas_tp.CrearPartido(nombre, [tdas_tp.CANT_VOTACION]string{presidente, gobernador, intendente})
		partidos = append(partidos, partido)
	}

	return partidos
}
