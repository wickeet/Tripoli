package internal

type Enfermeria struct {
	turnos map[*Enfermero][]Turno
}
