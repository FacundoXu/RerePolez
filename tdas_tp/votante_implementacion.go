package tdas_tp

import (
	"rerepolez/errores"
	"tdas/pila"
)

type votanteImplementacion struct {
	dni             int
	voto            Voto
	votosRealizados pila.Pila[Voto]
	estaVotando     bool
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.voto = *new(Voto)
	votante.votosRealizados = pila.CrearPilaDinamica[Voto]()
	votante.estaVotando = true
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if !votante.estaVotando {
		return errores.ErrorVotanteFraudulento{}
	}

	if alternativa == 0 {
		votante.voto.Impugnado = true
	}

	votante.voto.VotoPorTipo[tipo] = alternativa
	votante.votosRealizados.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if !votante.estaVotando {
		return errores.ErrorVotanteFraudulento{}
	}

	if votante.votosRealizados.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}

	votante.votosRealizados.Desapilar()

	if !votante.votosRealizados.EstaVacia() {
		votante.voto = votante.votosRealizados.VerTope()

	} else {
		votante.voto = Voto{}
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() Voto {
	votante.estaVotando = false
	return votante.voto
}

func (votante votanteImplementacion) EstaVotando() bool {
	return votante.estaVotando
}
