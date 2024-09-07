package tdas_tp

import "fmt"

type candidato struct {
	nombre   string
	cantidad int
}

type partidoImplementacion struct {
	nombre     string
	presidente *candidato
	gobernador *candidato
	intendente *candidato
}

type partidoEnBlanco struct {
	votosEnBlanco [CANT_VOTACION]int
}

func (partido *partidoImplementacion) crearCandidatos() {
	partido.presidente = new(candidato)
	partido.gobernador = new(candidato)
	partido.intendente = new(candidato)
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.crearCandidatos()
	partido.presidente.nombre = candidatos[PRESIDENTE]
	partido.gobernador.nombre = candidatos[GOBERNADOR]
	partido.intendente.nombre = candidatos[INTENDENTE]
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	switch tipo {

	case PRESIDENTE:
		partido.presidente.cantidad++

	case GOBERNADOR:
		partido.gobernador.cantidad++

	case INTENDENTE:
		partido.intendente.cantidad++
	}
}

func obtenerVotos(partido string, tipo string, cantidad int) string {
	if cantidad == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido, tipo, cantidad)
	}
	return fmt.Sprintf("%s - %s: %d votos", partido, tipo, cantidad)
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var resultado string

	switch tipo {

	case PRESIDENTE:
		resultado = obtenerVotos(partido.nombre, partido.presidente.nombre, partido.presidente.cantidad)

	case GOBERNADOR:
		resultado = obtenerVotos(partido.nombre, partido.gobernador.nombre, partido.gobernador.cantidad)

	case INTENDENTE:
		resultado = obtenerVotos(partido.nombre, partido.intendente.nombre, partido.intendente.cantidad)
	}

	return resultado
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosEnBlanco[tipo] += 1
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votosEnBlanco[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blanco.votosEnBlanco[tipo])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blanco.votosEnBlanco[tipo])
}
