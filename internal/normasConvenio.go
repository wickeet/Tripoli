package internal

import (
	"errors"
	"time"
)

type FechaFestiva struct {
	dia int
	mes time.Month
}

func NewFechaFestiva(dia int, mes time.Month) (*FechaFestiva, error) {
	if !esDiaValido(dia, mes) {
		return nil, errors.New("día inválido para el mes especificado")
	}
	return &FechaFestiva{dia: dia, mes: mes}, nil
}

func esDiaValido(dia int, mes time.Month) bool {
	if dia < 1 {
		return false
	}

	switch mes {
	case time.April, time.June, time.September, time.November:
		return dia <= 30
	case time.February:
		return dia <= 28
	default:
		return dia <= 31
	}
}

type NormasConvenio struct {
	maxHorasSemanales float64
	maxDiasSeguidos   int
	fechasFestivas    []FechaFestiva
}
