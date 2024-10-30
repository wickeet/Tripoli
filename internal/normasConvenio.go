package internal

import "time"

type NormasConvenio struct {
	maxHorasSemanales float64
	maxDiasSeguidos   int
	fechasFestivas    []time.Time
}
