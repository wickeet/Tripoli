package internal

import "time"

type FechaFestiva struct {
	dia int
	mes time.Month
}

type NormasConvenio struct {
	maxHorasSemanales float64
	maxDiasSeguidos   int
	fechasFestivas    []FechaFestiva
}
