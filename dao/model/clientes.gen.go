// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCliente = "Clientes"

// Cliente mapped from table <Clientes>
type Cliente struct {
	Codigo                int64     `gorm:"column:Codigo;primaryKey" json:"Codigo"`
	Nome                  string    `gorm:"column:Nome" json:"Nome"`
	DataCadastro          time.Time `gorm:"column:DataCadastro" json:"DataCadastro"`
	RazaoSocial           string    `gorm:"column:RazaoSocial" json:"RazaoSocial"`
	Limite                []uint8   `gorm:"column:Limite;not null;default:0" json:"Limite"`
	Endereco              string    `gorm:"column:Endereco" json:"Endereco"`
	CEP                   string    `gorm:"column:CEP" json:"CEP"`
	Bairro                string    `gorm:"column:Bairro" json:"Bairro"`
	UF                    string    `gorm:"column:UF" json:"UF"`
	Cidade                int64     `gorm:"column:Cidade;not null" json:"Cidade"`
	Fone                  string    `gorm:"column:Fone" json:"Fone"`
	Fax                   string    `gorm:"column:Fax" json:"Fax"`
	Contato               string    `gorm:"column:Contato" json:"Contato"`
	CNPJCPF               string    `gorm:"column:CNPJCPF" json:"CNPJCPF"`
	CGFRG                 string    `gorm:"column:CGFRG" json:"CGFRG"`
	Suframa               string    `gorm:"column:Suframa" json:"Suframa"`
	Email                 string    `gorm:"column:Email" json:"Email"`
	Classificacao         string    `gorm:"column:Classificacao;default:D" json:"Classificacao"`
	Vendedor              int64     `gorm:"column:Vendedor;not null" json:"Vendedor"`
	Corretor              int64     `gorm:"column:Corretor;not null" json:"Corretor"`
	NomeResponsavel       string    `gorm:"column:NomeResponsavel" json:"NomeResponsavel"`
	NascimentoResponsavel time.Time `gorm:"column:NascimentoResponsavel" json:"NascimentoResponsavel"`
	EnderecoResponsavel   string    `gorm:"column:EnderecoResponsavel" json:"EnderecoResponsavel"`
	CEPResponsavel        string    `gorm:"column:CEPResponsavel" json:"CEPResponsavel"`
	BairroResponsavel     string    `gorm:"column:BairroResponsavel" json:"BairroResponsavel"`
	UFResponsavel         string    `gorm:"column:UFResponsavel" json:"UFResponsavel"`
	CidadeResponsavel     int64     `gorm:"column:CidadeResponsavel;not null" json:"CidadeResponsavel"`
	FoneResponsavel       string    `gorm:"column:FoneResponsavel" json:"FoneResponsavel"`
	CPFResponsavel        string    `gorm:"column:CPFResponsavel" json:"CPFResponsavel"`
	RGResponsavel         string    `gorm:"column:RGResponsavel" json:"RGResponsavel"`
	MaeResponsavel        string    `gorm:"column:MaeResponsavel" json:"MaeResponsavel"`
	PontoReferencia       string    `gorm:"column:PontoReferencia" json:"PontoReferencia"`
	NomeTerceiro          string    `gorm:"column:NomeTerceiro" json:"NomeTerceiro"`
	NascimentoTerceiro    time.Time `gorm:"column:NascimentoTerceiro" json:"NascimentoTerceiro"`
	EnderecoTerceiro      string    `gorm:"column:EnderecoTerceiro" json:"EnderecoTerceiro"`
	CEPTerceiro           string    `gorm:"column:CEPTerceiro" json:"CEPTerceiro"`
	BairroTerceiro        string    `gorm:"column:BairroTerceiro" json:"BairroTerceiro"`
	UFTerceiro            string    `gorm:"column:UFTerceiro" json:"UFTerceiro"`
	CidadeTerceiro        int64     `gorm:"column:CidadeTerceiro;not null" json:"CidadeTerceiro"`
	FoneTerceiro          string    `gorm:"column:FoneTerceiro" json:"FoneTerceiro"`
	CPFTerceiro           string    `gorm:"column:CPFTerceiro" json:"CPFTerceiro"`
	RGTerceiro            string    `gorm:"column:RGTerceiro" json:"RGTerceiro"`
	MaeTerceiro           string    `gorm:"column:MaeTerceiro" json:"MaeTerceiro"`
	Obs                   string    `gorm:"column:Obs" json:"Obs"`
	RefBanco1             string    `gorm:"column:RefBanco1" json:"RefBanco1"`
	FoneBanco1            string    `gorm:"column:FoneBanco1" json:"FoneBanco1"`
	RefBanco2             string    `gorm:"column:RefBanco2" json:"RefBanco2"`
	FoneBanco2            string    `gorm:"column:FoneBanco2" json:"FoneBanco2"`
	Bloqueado             bool      `gorm:"column:Bloqueado;not null;default:0" json:"Bloqueado"`
	Inativo               bool      `gorm:"column:Inativo;not null;default:0" json:"Inativo"`
	Fisica                bool      `gorm:"column:Fisica;not null;default:0" json:"Fisica"`
	DataCompra            time.Time `gorm:"column:DataCompra" json:"DataCompra"`
	ValorCompra           []uint8   `gorm:"column:ValorCompra;not null;default:0" json:"ValorCompra"`
	Foto                  string    `gorm:"column:Foto" json:"Foto"`
	Atualizado            int64     `gorm:"column:Atualizado;not null" json:"Atualizado"`
	ObsRepresentante      string    `gorm:"column:ObsRepresentante" json:"ObsRepresentante"`
	OBSPedido             string    `gorm:"column:OBSPedido" json:"OBSPedido"`
	DescontoEspecial      float64   `gorm:"column:DescontoEspecial;not null" json:"DescontoEspecial"`
	IsentodeCNPJ          bool      `gorm:"column:IsentodeCNPJ;not null;default:0" json:"IsentodeCNPJ"`
	QtdeLojas             int64     `gorm:"column:QtdeLojas;not null" json:"QtdeLojas"`
	ValorCredito          []uint8   `gorm:"column:ValorCredito;not null;default:0" json:"ValorCredito"`
	CodigoLoja            int64     `gorm:"column:CodigoLoja;not null" json:"CodigoLoja"`
	EnderecoCobranca      string    `gorm:"column:EnderecoCobranca" json:"EnderecoCobranca"`
	CepCobranca           string    `gorm:"column:CepCobranca" json:"CepCobranca"`
	BairroCobranca        string    `gorm:"column:BairroCobranca" json:"BairroCobranca"`
	UFCobranca            string    `gorm:"column:UFCobranca" json:"UFCobranca"`
	CidadeCobranca        int64     `gorm:"column:CidadeCobranca;not null" json:"CidadeCobranca"`
	FoneCobranca          string    `gorm:"column:FoneCobranca" json:"FoneCobranca"`
	VendedorLoja          int64     `gorm:"column:VendedorLoja;not null" json:"VendedorLoja"`
	ChequeNaLoja          bool      `gorm:"column:ChequeNaLoja;default:0" json:"ChequeNaLoja"`
	CodigoPais            string    `gorm:"column:CodigoPais" json:"CodigoPais"`
	RGCliente             string    `gorm:"column:RGCliente" json:"RGCliente"`
	FormaPagtoPadrao      int64     `gorm:"column:FormaPagtoPadrao" json:"FormaPagtoPadrao"`
	CondPagtoPadrao       int64     `gorm:"column:CondPagtoPadrao" json:"CondPagtoPadrao"`
	MsgOrcamento          string    `gorm:"column:msgOrcamento" json:"msgOrcamento"`
	DiasTroca             int64     `gorm:"column:DiasTroca" json:"DiasTroca"`
	GrupoEmpresarial      int64     `gorm:"column:GrupoEmpresarial;not null" json:"GrupoEmpresarial"`
	Numero                string    `gorm:"column:Numero" json:"Numero"`
	TaxaEntrega           []uint8   `gorm:"column:TaxaEntrega;default:0" json:"TaxaEntrega"`
	Manequim              string    `gorm:"column:Manequim" json:"Manequim"`
}

// TableName Cliente's table name
func (*Cliente) TableName() string {
	return TableNameCliente
}
