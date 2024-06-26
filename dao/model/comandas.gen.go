// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComanda = "Comandas"

// Comanda mapped from table <Comandas>
type Comanda struct {
	Comanda                   int64     `gorm:"column:Comanda;primaryKey" json:"Comanda"`
	Data                      time.Time `gorm:"column:Data" json:"Data"`
	Cliente                   int64     `gorm:"column:Cliente;not null" json:"Cliente"`
	Consumidor                string    `gorm:"column:Consumidor" json:"Consumidor"`
	Vendedor                  int64     `gorm:"column:Vendedor;not null" json:"Vendedor"`
	ValorBruto                []uint8   `gorm:"column:ValorBruto;not null;default:0" json:"ValorBruto"`
	ValorDesconto             []uint8   `gorm:"column:ValorDesconto;not null;default:0" json:"ValorDesconto"`
	Abono                     []uint8   `gorm:"column:Abono;not null;default:0" json:"Abono"`
	CodCondPagto              int64     `gorm:"column:CodCondPagto;not null" json:"CodCondPagto"`
	CondPagto                 string    `gorm:"column:CondPagto" json:"CondPagto"`
	FormaPagto                string    `gorm:"column:FormaPagto" json:"FormaPagto"`
	CodFormaPagto             int64     `gorm:"column:CodFormaPagto;not null" json:"CodFormaPagto"`
	Corretor                  int64     `gorm:"column:Corretor;not null" json:"Corretor"`
	Loja                      int64     `gorm:"column:Loja;not null;default:1" json:"Loja"`
	TaxaEntrega               []uint8   `gorm:"column:TaxaEntrega;not null;default:0" json:"TaxaEntrega"`
	AVista                    []uint8   `gorm:"column:AVista;not null;default:0" json:"AVista"`
	APrazo                    []uint8   `gorm:"column:APrazo;not null;default:0" json:"APrazo"`
	Tabela                    int64     `gorm:"column:Tabela;not null;default:1" json:"Tabela"`
	TipoVenda                 string    `gorm:"column:TipoVenda" json:"TipoVenda"`
	Pedido                    int64     `gorm:"column:Pedido;not null" json:"Pedido"`
	Funcionario               bool      `gorm:"column:Funcionario;not null;default:0" json:"Funcionario"`
	Varejo                    bool      `gorm:"column:Varejo;not null;default:0" json:"Varejo"`
	Comercial                 bool      `gorm:"column:Comercial;not null;default:0" json:"Comercial"`
	PercentualBonificacao     float64   `gorm:"column:PercentualBonificacao;not null" json:"PercentualBonificacao"`
	ValorBonificacao          []uint8   `gorm:"column:ValorBonificacao;not null;default:0" json:"ValorBonificacao"`
	EnviadoparaoRepresentante bool      `gorm:"column:EnviadoparaoRepresentante;not null;default:0" json:"EnviadoparaoRepresentante"`
	NotaFiscal                int64     `gorm:"column:NotaFiscal;not null" json:"NotaFiscal"`
	Justificativa             string    `gorm:"column:Justificativa" json:"Justificativa"`
	CupomFiscal               int64     `gorm:"column:CupomFiscal" json:"CupomFiscal"`
	Suframa                   bool      `gorm:"column:Suframa;not null;default:0" json:"Suframa"`
	NumerodoCaixa             int64     `gorm:"column:NumerodoCaixa;not null" json:"NumerodoCaixa"`
	MotivodaDoacao            string    `gorm:"column:MotivodaDoacao" json:"MotivodaDoacao"`
	ControledeSeparacao       int64     `gorm:"column:ControledeSeparacao;not null" json:"ControledeSeparacao"`
	ComissaodoVendedor        float64   `gorm:"column:ComissaodoVendedor;not null" json:"ComissaodoVendedor"`
	NotaFiscalSerie           string    `gorm:"column:NotaFiscal_Serie" json:"NotaFiscal_Serie"`
	NotaFiscalLoja            int64     `gorm:"column:NotaFiscal_Loja" json:"NotaFiscal_Loja"`
	NotaFiscalModelo          int64     `gorm:"column:NotaFiscal_Modelo" json:"NotaFiscal_Modelo"`
	CodBalconista             int64     `gorm:"column:CodBalconista" json:"CodBalconista"`
	Acrescimo                 []uint8   `gorm:"column:acrescimo" json:"acrescimo"`
	Arredondamento            []uint8   `gorm:"column:arredondamento" json:"arredondamento"`
	Hora                      int64     `gorm:"column:hora" json:"hora"`
	CNPJCPFConsumidor         string    `gorm:"column:CNPJCPFConsumidor" json:"CNPJCPFConsumidor"`
	PremiouCorretor           bool      `gorm:"column:PremiouCorretor;default:0" json:"PremiouCorretor"`
	LancouComissao            bool      `gorm:"column:LancouComissao;default:0" json:"LancouComissao"`
	CodigoTransportadora      int64     `gorm:"column:CodigoTransportadora;not null" json:"CodigoTransportadora"`
	ValorTotalDinheiro        []uint8   `gorm:"column:ValorTotalDinheiro;not null;default:0" json:"ValorTotalDinheiro"`
	ValorDescontoDinheiro     []uint8   `gorm:"column:ValorDescontoDinheiro;not null;default:0" json:"ValorDescontoDinheiro"`
	ValorTotalCartao          []uint8   `gorm:"column:ValorTotalCartao;not null;default:0" json:"ValorTotalCartao"`
	ValorDescontoCartao       []uint8   `gorm:"column:ValorDescontoCartao;not null;default:0" json:"ValorDescontoCartao"`
	PrecodoKG                 []uint8   `gorm:"column:PrecodoKG;not null;default:0" json:"PrecodoKG"`
	ConciliacaoBancaria       int64     `gorm:"column:ConciliacaoBancaria;not null" json:"ConciliacaoBancaria"`
	CanalVenda                int64     `gorm:"column:CanalVenda;not null" json:"CanalVenda"`
	Entrega                   string    `gorm:"column:Entrega" json:"Entrega"`
	Observacao                string    `gorm:"column:Observacao" json:"Observacao"`
}

// TableName Comanda's table name
func (*Comanda) TableName() string {
	return TableNameComanda
}
