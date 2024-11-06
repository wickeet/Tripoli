package internal

import "time"

type HoraMinuto struct {
	hora   int
	minuto int
}

type Turno struct {
	diaSemana time.Weekday
	inicio    HoraMinuto
	duracion  time.Duration
}
