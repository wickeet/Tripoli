package internal

import "time"

type Enfermero struct {
	nombre    string
	apellidos string
	rol       string
}

type TurnosPreferidos map[*Enfermero][]Turno

type Disponibilidad map[*Enfermero][]time.Weekday

type TurnosAsignados map[*Enfermero][]Turno
