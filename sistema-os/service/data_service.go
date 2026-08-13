package service

import "time"

func CalcularDiasRestantes(dataEntrega time.Time) int {
	diferenca := dataEntrega.Sub(time.Now())
	dias := int(diferenca.Hours() / 24)

	return dias
}
