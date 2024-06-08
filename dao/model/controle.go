package model

import "time"

// Mudando nomes de tabelas para evitar conflito com api woocommerce
const TableNameIntegracaoProduto = "Firebase_Produto"
const TableNameIntegracaoGrupo = "Firebase_Grupo"
const TableNameIntegracaoVariacao = "Firebase_Variacao"

// IntegracaoProduto Produtos
type IntegracaoProduto struct {
	ID           string    `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Referencia   string    `gorm:"column:Referencia" json:"Referencia"`
	Contador     int64     `gorm:"column:Contador" json:"Contador"`
	AtualizadoEm time.Time `gorm:"column:AtualizadoEm" json:"AtualizadoEm"`
}

// TableName Controle Produtos
func (*IntegracaoProduto) TableName() string {
	return TableNameIntegracaoProduto
}

// IntegracaoGrupo com Woo
type IntegracaoGrupo struct {
	ID     string `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Codigo int64  `gorm:"column:Codigo" json:"Codigo"`
}

// TableName Controle Grupo
func (*IntegracaoGrupo) TableName() string {
	return TableNameIntegracaoGrupo
}

// IntegracaoVariacao com Woo
type IntegracaoVariacao struct {
	ID           int64     `gorm:"column:Id;primaryKey;autoIncrement:false" json:"ID"`
	Referencia   string    `gorm:"column:Referencia" json:"Referencia"`
	Cor          int64     `gorm:"column:Cor" json:"Cor"`
	Tam          string    `gorm:"column:Tam" json:"Tam"`
	Coluna       int64     `gorm:"column:Coluna" json:"Coluna"`
	Contador     int64     `gorm:"column:Contador" json:"Contador"`
	AtualizadoEm time.Time `gorm:"column:AtualizadoEm" json:"AtualizadoEm"`
}

// TableName Controle Variacao Cor Tamanho
func (*IntegracaoVariacao) TableName() string {
	return TableNameIntegracaoVariacao
}

type QueryVariation struct {
	Referencia  string  `gorm:"column:Referencia" json:"Referencia"`
	Descricao   string  `gorm:"column:Descricao" json:"Descricao"`
	Unidade     string  `gorm:"column:UND" json:"Unidade"`
	Preco       []uint8 `gorm:"column:Preco;default:0" json:"Preco"`
	TamanhoDesc string  `gorm:"column:TamanhoDesc" json:"TamanhoDesc"`
	Tamanho     string  `gorm:"column:Tamanho" json:"Tamanho"`
	Coluna      string  `gorm:"column:Coluna" json:"Coluna"`
	Cor         int64   `gorm:"column:Cor" json:"Cor"`
	NomeCor     string  `gorm:"column:NomeCor" json:"NomeCor"`
	Saldo       int64   `gorm:"column:Saldo" json:"Saldo"`
	Rank        int64   `gorm:"column:Rank" json:"Rank"`
}

type QuerySequencias struct {
	Ultimo int64 `gorm:"column:Ultimo" json:"Ultimo"`
}
