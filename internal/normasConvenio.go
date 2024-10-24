package internal

import "time"

type NormasConvenio struct {
	MaxHorasSemanales float64
	MaxDiasSeguidos   int
	FechasFestivas    []time.Time
}
