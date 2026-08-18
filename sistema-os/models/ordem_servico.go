package models

import "time"

type OrdemServico struct {
	ID            int       `json:"id"`
	Cliente       string    `json:"cliente"`
	Descricao     string    `json:"descricao"`
	ValorEstimado float64   `json:"valor_estimado"`
	DataDeEntrega time.Time `json:"data_de_entrega"`
	Status        string    `json:"status"`
	ValorFinal    float64   `json:"valor_final"`
}
