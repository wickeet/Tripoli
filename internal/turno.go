package internal

import "time"

type Turno struct {
	diaSemana time.Weekday
	inicio    time.Time
	fin       time.Time
}
