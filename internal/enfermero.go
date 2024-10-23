package internal

import "time"

type PeriodoTiempo struct {
	DiaSemana  time.Weekday
	HoraInicio time.Time
	HoraFin    time.Time
}

type PeriodoFecha struct {
	FechaInicio time.Time
	FechaFin    time.Time
}

type Disponibilidad struct {
	HorariosDisponibles    []PeriodoTiempo
	FechasIndisponibilidad []PeriodoFecha
}

type Enfermero struct {
	ID               int
	Nombre           string
	Apellidos        string
	Rol              string
	TurnosPreferidos []PeriodoTiempo
	Disponibilidad   Disponibilidad
}
