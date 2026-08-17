package service

import "time"

func CalcularDiasRestantes(dataEntrega time.Time) int {
	hoje := time.Now()

	inicioHoje := time.Date(
		hoje.Year(),
		hoje.Month(),
		hoje.Day(),
		0,
		0,
		0,
		0,
		hoje.Location(),
	)

	inicioEntrega := time.Date(
		dataEntrega.Year(),
		dataEntrega.Month(),
		dataEntrega.Day(),
		0,
		0,
		0,
		0,
		hoje.Location(),
	)

	return int(inicioEntrega.Sub(inicioHoje).Hours() / 24)
}