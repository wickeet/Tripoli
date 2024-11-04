package internal

type Enfermero struct {
	disponibilidad []Turno
	preferencias   []Turno
}

type TurnosAsignados map[*Enfermero][]Turno
