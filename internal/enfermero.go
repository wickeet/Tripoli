package internal

import "time"

type Enfermero struct {
	nombre    string
	apellidos string
	rol       string
}

type TurnosPreferidos map[*Enfermero][]Turno

type DiasTrabajo map[*Enfermero][]time.Weekday

type TurnosAsignados map[*Enfermero][]Turno
