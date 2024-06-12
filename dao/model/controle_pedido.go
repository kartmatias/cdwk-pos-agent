package model

import "time"

const TableNameIntegracaoPedido = "Firebase_Pedido"

type IntegracaoPedido struct {
	ID           string    `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Comanda      int64     `gorm:"column:Comanda" json:"Comanda"`
	AtualizadoEm time.Time `gorm:"column:AtualizadoEm" json:"AtualizadoEm"`
}

// TableName Controle Produtos
func (*IntegracaoPedido) TableName() string {
	return TableNameIntegracaoPedido
}
