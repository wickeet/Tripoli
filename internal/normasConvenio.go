package internal

import "time"

type NormasConvenio struct {
	MaxHorasSemanales       float64
	MinHorasDescansoSemanal float64
	FechasFestivas          []time.Time
}
