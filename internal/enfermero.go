package internal

import "time"

type Enfermero struct {
	nombre    string
	apellidos string
	rol       string
}

type Disponibilidad struct {
	diasTrabajo map[*Enfermero][]time.Weekday
}

type Preferencias struct {
	turnosPreferidos map[*Enfermero][]Turno
}
