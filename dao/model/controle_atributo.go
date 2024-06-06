package model

const TableNameIntegracaoAtributo = "Integracao_Atributo"

type IntegracaoAtributo struct {
	ID       int64  `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Atributo string `gorm:"column:Atributo" json:"Atributo"`
}

// TableName Controle Produtos
func (*IntegracaoAtributo) TableName() string {
	return TableNameIntegracaoAtributo
}
