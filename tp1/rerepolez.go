package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/auxiliares"
	"rerepolez/errores"
	"rerepolez/usuario"
	"rerepolez/tdas_tp"
	"strings"
	"tdas/cola"
)

const (
	LISTA                     = 1
	PADRON                    = 2
	MAX_PARAMETROS_REQUERIDOS = 3
)

func main() {

	// Variables a utilizar
	args := os.Args

	// Error si no hay suficiente parametros ingresados
	if len(args) < MAX_PARAMETROS_REQUERIDOS {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}

	// Cargamos la informacion de los partidos y votantes
	partidos := auxiliares.CargarPartidos(args[LISTA])
	votantes := auxiliares.CargarPadrones(args[PADRON])

	// Error si no existen los archivos
	if partidos == nil || votantes == nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		return
	}

	// Fila de votantes y cantidad de votos impugnados
	fila := cola.CrearColaEnlazada[tdas_tp.Votante]()
	impugnados := 0

	// Pedimos entradas al usuario
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		// Entrada del usuario
		entrada := s.Text()

		// Separamos la entrada
		cmd := strings.Split(entrada, " ")

		// Acciones dependiendo del comando ingresado por el usuario
		usuario.Comandos(cmd, partidos, votantes, fila, &impugnados)
	}

	// Error si todavia hay ciudadanos sin votar
	if !fila.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{}.Error())
	}

	// Mostramos los resultados de las votaciones
	usuario.MostrarResultados(partidos, &impugnados)
}
