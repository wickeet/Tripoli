package internal

import "time"

type Turno struct {
	ID          int
	EnfermeroID int
	Inicio      time.Time
	Fin         time.Time
}
