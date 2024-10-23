package internal

import "time"

type Descanso struct {
	ID          int
	EnfermeroID int
	Inicio      time.Time
	Fin         time.Time
}
