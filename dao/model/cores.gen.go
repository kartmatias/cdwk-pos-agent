// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCore = "Cores"

// Core mapped from table <Cores>
type Core struct {
	Codigo    int64  `gorm:"column:Codigo;primaryKey" json:"Codigo"`
	Descricao string `gorm:"column:Descricao" json:"Descricao"`
	Almox     bool   `gorm:"column:Almox;not null;default:1" json:"Almox"`
	DPA       bool   `gorm:"column:DPA;not null;default:1" json:"DPA"`
	Sigla     string `gorm:"column:Sigla" json:"Sigla"`
}

// TableName Core's table name
func (*Core) TableName() string {
	return TableNameCore
}
