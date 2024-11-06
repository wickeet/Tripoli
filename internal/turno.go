package internal

import (
	"errors"
	"time"
)

type HoraMinuto struct {
	hora   int
	minuto int
}

func NewHoraMinuto(hora, minuto int) (*HoraMinuto, error) {
	if !esHoraValida(hora, minuto) {
		return nil, errors.New("hora inválida")
	}
	return &HoraMinuto{hora: hora, minuto: minuto}, nil
}

func esHoraValida(hora, minuto int) bool {
	if hora < 0 || hora > 23 {
		return false
	}

	minutosValidos := []int{0, 15, 30, 45}
	for _, m := range minutosValidos {
		if minuto == m {
			return true
		}
	}
	return false
}

type Turno struct {
	diaSemana time.Weekday
	inicio    HoraMinuto
	duracion  time.Duration
}
