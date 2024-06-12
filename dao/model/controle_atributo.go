package model

const TableNameIntegracaoAtributo = "Firebase_Atributo"

type IntegracaoAtributo struct {
	ID       string `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Atributo string `gorm:"column:Atributo" json:"Atributo"`
}

// TableName Controle Produtos
func (*IntegracaoAtributo) TableName() string {
	return TableNameIntegracaoAtributo
}
