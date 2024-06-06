// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/kartmatias/cdwk-pos-agent/dao/model"
)

func newCliente(db *gorm.DB, opts ...gen.DOOption) cliente {
	_cliente := cliente{}

	_cliente.clienteDo.UseDB(db, opts...)
	_cliente.clienteDo.UseModel(&model.Cliente{})

	tableName := _cliente.clienteDo.TableName()
	_cliente.ALL = field.NewAsterisk(tableName)
	_cliente.Codigo = field.NewInt64(tableName, "Codigo")
	_cliente.Nome = field.NewString(tableName, "Nome")
	_cliente.DataCadastro = field.NewTime(tableName, "DataCadastro")
	_cliente.RazaoSocial = field.NewString(tableName, "RazaoSocial")
	_cliente.Limite = field.NewField(tableName, "Limite")
	_cliente.Endereco = field.NewString(tableName, "Endereco")
	_cliente.CEP = field.NewString(tableName, "CEP")
	_cliente.Bairro = field.NewString(tableName, "Bairro")
	_cliente.UF = field.NewString(tableName, "UF")
	_cliente.Cidade = field.NewInt64(tableName, "Cidade")
	_cliente.Fone = field.NewString(tableName, "Fone")
	_cliente.Fax = field.NewString(tableName, "Fax")
	_cliente.Contato = field.NewString(tableName, "Contato")
	_cliente.CNPJCPF = field.NewString(tableName, "CNPJCPF")
	_cliente.CGFRG = field.NewString(tableName, "CGFRG")
	_cliente.Suframa = field.NewString(tableName, "Suframa")
	_cliente.Email = field.NewString(tableName, "Email")
	_cliente.Classificacao = field.NewString(tableName, "Classificacao")
	_cliente.Vendedor = field.NewInt64(tableName, "Vendedor")
	_cliente.Corretor = field.NewInt64(tableName, "Corretor")
	_cliente.NomeResponsavel = field.NewString(tableName, "NomeResponsavel")
	_cliente.NascimentoResponsavel = field.NewTime(tableName, "NascimentoResponsavel")
	_cliente.EnderecoResponsavel = field.NewString(tableName, "EnderecoResponsavel")
	_cliente.CEPResponsavel = field.NewString(tableName, "CEPResponsavel")
	_cliente.BairroResponsavel = field.NewString(tableName, "BairroResponsavel")
	_cliente.UFResponsavel = field.NewString(tableName, "UFResponsavel")
	_cliente.CidadeResponsavel = field.NewInt64(tableName, "CidadeResponsavel")
	_cliente.FoneResponsavel = field.NewString(tableName, "FoneResponsavel")
	_cliente.CPFResponsavel = field.NewString(tableName, "CPFResponsavel")
	_cliente.RGResponsavel = field.NewString(tableName, "RGResponsavel")
	_cliente.MaeResponsavel = field.NewString(tableName, "MaeResponsavel")
	_cliente.PontoReferencia = field.NewString(tableName, "PontoReferencia")
	_cliente.NomeTerceiro = field.NewString(tableName, "NomeTerceiro")
	_cliente.NascimentoTerceiro = field.NewTime(tableName, "NascimentoTerceiro")
	_cliente.EnderecoTerceiro = field.NewString(tableName, "EnderecoTerceiro")
	_cliente.CEPTerceiro = field.NewString(tableName, "CEPTerceiro")
	_cliente.BairroTerceiro = field.NewString(tableName, "BairroTerceiro")
	_cliente.UFTerceiro = field.NewString(tableName, "UFTerceiro")
	_cliente.CidadeTerceiro = field.NewInt64(tableName, "CidadeTerceiro")
	_cliente.FoneTerceiro = field.NewString(tableName, "FoneTerceiro")
	_cliente.CPFTerceiro = field.NewString(tableName, "CPFTerceiro")
	_cliente.RGTerceiro = field.NewString(tableName, "RGTerceiro")
	_cliente.MaeTerceiro = field.NewString(tableName, "MaeTerceiro")
	_cliente.Obs = field.NewString(tableName, "Obs")
	_cliente.RefBanco1 = field.NewString(tableName, "RefBanco1")
	_cliente.FoneBanco1 = field.NewString(tableName, "FoneBanco1")
	_cliente.RefBanco2 = field.NewString(tableName, "RefBanco2")
	_cliente.FoneBanco2 = field.NewString(tableName, "FoneBanco2")
	_cliente.Bloqueado = field.NewBool(tableName, "Bloqueado")
	_cliente.Inativo = field.NewBool(tableName, "Inativo")
	_cliente.Fisica = field.NewBool(tableName, "Fisica")
	_cliente.DataCompra = field.NewTime(tableName, "DataCompra")
	_cliente.ValorCompra = field.NewField(tableName, "ValorCompra")
	_cliente.Foto = field.NewString(tableName, "Foto")
	_cliente.Atualizado = field.NewInt64(tableName, "Atualizado")
	_cliente.ObsRepresentante = field.NewString(tableName, "ObsRepresentante")
	_cliente.OBSPedido = field.NewString(tableName, "OBSPedido")
	_cliente.DescontoEspecial = field.NewFloat64(tableName, "DescontoEspecial")
	_cliente.IsentodeCNPJ = field.NewBool(tableName, "IsentodeCNPJ")
	_cliente.QtdeLojas = field.NewInt64(tableName, "QtdeLojas")
	_cliente.ValorCredito = field.NewField(tableName, "ValorCredito")
	_cliente.CodigoLoja = field.NewInt64(tableName, "CodigoLoja")
	_cliente.EnderecoCobranca = field.NewString(tableName, "EnderecoCobranca")
	_cliente.CepCobranca = field.NewString(tableName, "CepCobranca")
	_cliente.BairroCobranca = field.NewString(tableName, "BairroCobranca")
	_cliente.UFCobranca = field.NewString(tableName, "UFCobranca")
	_cliente.CidadeCobranca = field.NewInt64(tableName, "CidadeCobranca")
	_cliente.FoneCobranca = field.NewString(tableName, "FoneCobranca")
	_cliente.VendedorLoja = field.NewInt64(tableName, "VendedorLoja")
	_cliente.ChequeNaLoja = field.NewBool(tableName, "ChequeNaLoja")
	_cliente.CodigoPais = field.NewString(tableName, "CodigoPais")
	_cliente.RGCliente = field.NewString(tableName, "RGCliente")
	_cliente.FormaPagtoPadrao = field.NewInt64(tableName, "FormaPagtoPadrao")
	_cliente.CondPagtoPadrao = field.NewInt64(tableName, "CondPagtoPadrao")
	_cliente.MsgOrcamento = field.NewString(tableName, "msgOrcamento")
	_cliente.DiasTroca = field.NewInt64(tableName, "DiasTroca")
	_cliente.GrupoEmpresarial = field.NewInt64(tableName, "GrupoEmpresarial")
	_cliente.Numero = field.NewString(tableName, "Numero")
	_cliente.TaxaEntrega = field.NewField(tableName, "TaxaEntrega")
	_cliente.Manequim = field.NewString(tableName, "Manequim")

	_cliente.fillFieldMap()

	return _cliente
}

type cliente struct {
	clienteDo clienteDo

	ALL                   field.Asterisk
	Codigo                field.Int64
	Nome                  field.String
	DataCadastro          field.Time
	RazaoSocial           field.String
	Limite                field.Field
	Endereco              field.String
	CEP                   field.String
	Bairro                field.String
	UF                    field.String
	Cidade                field.Int64
	Fone                  field.String
	Fax                   field.String
	Contato               field.String
	CNPJCPF               field.String
	CGFRG                 field.String
	Suframa               field.String
	Email                 field.String
	Classificacao         field.String
	Vendedor              field.Int64
	Corretor              field.Int64
	NomeResponsavel       field.String
	NascimentoResponsavel field.Time
	EnderecoResponsavel   field.String
	CEPResponsavel        field.String
	BairroResponsavel     field.String
	UFResponsavel         field.String
	CidadeResponsavel     field.Int64
	FoneResponsavel       field.String
	CPFResponsavel        field.String
	RGResponsavel         field.String
	MaeResponsavel        field.String
	PontoReferencia       field.String
	NomeTerceiro          field.String
	NascimentoTerceiro    field.Time
	EnderecoTerceiro      field.String
	CEPTerceiro           field.String
	BairroTerceiro        field.String
	UFTerceiro            field.String
	CidadeTerceiro        field.Int64
	FoneTerceiro          field.String
	CPFTerceiro           field.String
	RGTerceiro            field.String
	MaeTerceiro           field.String
	Obs                   field.String
	RefBanco1             field.String
	FoneBanco1            field.String
	RefBanco2             field.String
	FoneBanco2            field.String
	Bloqueado             field.Bool
	Inativo               field.Bool
	Fisica                field.Bool
	DataCompra            field.Time
	ValorCompra           field.Field
	Foto                  field.String
	Atualizado            field.Int64
	ObsRepresentante      field.String
	OBSPedido             field.String
	DescontoEspecial      field.Float64
	IsentodeCNPJ          field.Bool
	QtdeLojas             field.Int64
	ValorCredito          field.Field
	CodigoLoja            field.Int64
	EnderecoCobranca      field.String
	CepCobranca           field.String
	BairroCobranca        field.String
	UFCobranca            field.String
	CidadeCobranca        field.Int64
	FoneCobranca          field.String
	VendedorLoja          field.Int64
	ChequeNaLoja          field.Bool
	CodigoPais            field.String
	RGCliente             field.String
	FormaPagtoPadrao      field.Int64
	CondPagtoPadrao       field.Int64
	MsgOrcamento          field.String
	DiasTroca             field.Int64
	GrupoEmpresarial      field.Int64
	Numero                field.String
	TaxaEntrega           field.Field
	Manequim              field.String

	fieldMap map[string]field.Expr
}

func (c cliente) Table(newTableName string) *cliente {
	c.clienteDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cliente) As(alias string) *cliente {
	c.clienteDo.DO = *(c.clienteDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cliente) updateTableName(table string) *cliente {
	c.ALL = field.NewAsterisk(table)
	c.Codigo = field.NewInt64(table, "Codigo")
	c.Nome = field.NewString(table, "Nome")
	c.DataCadastro = field.NewTime(table, "DataCadastro")
	c.RazaoSocial = field.NewString(table, "RazaoSocial")
	c.Limite = field.NewField(table, "Limite")
	c.Endereco = field.NewString(table, "Endereco")
	c.CEP = field.NewString(table, "CEP")
	c.Bairro = field.NewString(table, "Bairro")
	c.UF = field.NewString(table, "UF")
	c.Cidade = field.NewInt64(table, "Cidade")
	c.Fone = field.NewString(table, "Fone")
	c.Fax = field.NewString(table, "Fax")
	c.Contato = field.NewString(table, "Contato")
	c.CNPJCPF = field.NewString(table, "CNPJCPF")
	c.CGFRG = field.NewString(table, "CGFRG")
	c.Suframa = field.NewString(table, "Suframa")
	c.Email = field.NewString(table, "Email")
	c.Classificacao = field.NewString(table, "Classificacao")
	c.Vendedor = field.NewInt64(table, "Vendedor")
	c.Corretor = field.NewInt64(table, "Corretor")
	c.NomeResponsavel = field.NewString(table, "NomeResponsavel")
	c.NascimentoResponsavel = field.NewTime(table, "NascimentoResponsavel")
	c.EnderecoResponsavel = field.NewString(table, "EnderecoResponsavel")
	c.CEPResponsavel = field.NewString(table, "CEPResponsavel")
	c.BairroResponsavel = field.NewString(table, "BairroResponsavel")
	c.UFResponsavel = field.NewString(table, "UFResponsavel")
	c.CidadeResponsavel = field.NewInt64(table, "CidadeResponsavel")
	c.FoneResponsavel = field.NewString(table, "FoneResponsavel")
	c.CPFResponsavel = field.NewString(table, "CPFResponsavel")
	c.RGResponsavel = field.NewString(table, "RGResponsavel")
	c.MaeResponsavel = field.NewString(table, "MaeResponsavel")
	c.PontoReferencia = field.NewString(table, "PontoReferencia")
	c.NomeTerceiro = field.NewString(table, "NomeTerceiro")
	c.NascimentoTerceiro = field.NewTime(table, "NascimentoTerceiro")
	c.EnderecoTerceiro = field.NewString(table, "EnderecoTerceiro")
	c.CEPTerceiro = field.NewString(table, "CEPTerceiro")
	c.BairroTerceiro = field.NewString(table, "BairroTerceiro")
	c.UFTerceiro = field.NewString(table, "UFTerceiro")
	c.CidadeTerceiro = field.NewInt64(table, "CidadeTerceiro")
	c.FoneTerceiro = field.NewString(table, "FoneTerceiro")
	c.CPFTerceiro = field.NewString(table, "CPFTerceiro")
	c.RGTerceiro = field.NewString(table, "RGTerceiro")
	c.MaeTerceiro = field.NewString(table, "MaeTerceiro")
	c.Obs = field.NewString(table, "Obs")
	c.RefBanco1 = field.NewString(table, "RefBanco1")
	c.FoneBanco1 = field.NewString(table, "FoneBanco1")
	c.RefBanco2 = field.NewString(table, "RefBanco2")
	c.FoneBanco2 = field.NewString(table, "FoneBanco2")
	c.Bloqueado = field.NewBool(table, "Bloqueado")
	c.Inativo = field.NewBool(table, "Inativo")
	c.Fisica = field.NewBool(table, "Fisica")
	c.DataCompra = field.NewTime(table, "DataCompra")
	c.ValorCompra = field.NewField(table, "ValorCompra")
	c.Foto = field.NewString(table, "Foto")
	c.Atualizado = field.NewInt64(table, "Atualizado")
	c.ObsRepresentante = field.NewString(table, "ObsRepresentante")
	c.OBSPedido = field.NewString(table, "OBSPedido")
	c.DescontoEspecial = field.NewFloat64(table, "DescontoEspecial")
	c.IsentodeCNPJ = field.NewBool(table, "IsentodeCNPJ")
	c.QtdeLojas = field.NewInt64(table, "QtdeLojas")
	c.ValorCredito = field.NewField(table, "ValorCredito")
	c.CodigoLoja = field.NewInt64(table, "CodigoLoja")
	c.EnderecoCobranca = field.NewString(table, "EnderecoCobranca")
	c.CepCobranca = field.NewString(table, "CepCobranca")
	c.BairroCobranca = field.NewString(table, "BairroCobranca")
	c.UFCobranca = field.NewString(table, "UFCobranca")
	c.CidadeCobranca = field.NewInt64(table, "CidadeCobranca")
	c.FoneCobranca = field.NewString(table, "FoneCobranca")
	c.VendedorLoja = field.NewInt64(table, "VendedorLoja")
	c.ChequeNaLoja = field.NewBool(table, "ChequeNaLoja")
	c.CodigoPais = field.NewString(table, "CodigoPais")
	c.RGCliente = field.NewString(table, "RGCliente")
	c.FormaPagtoPadrao = field.NewInt64(table, "FormaPagtoPadrao")
	c.CondPagtoPadrao = field.NewInt64(table, "CondPagtoPadrao")
	c.MsgOrcamento = field.NewString(table, "msgOrcamento")
	c.DiasTroca = field.NewInt64(table, "DiasTroca")
	c.GrupoEmpresarial = field.NewInt64(table, "GrupoEmpresarial")
	c.Numero = field.NewString(table, "Numero")
	c.TaxaEntrega = field.NewField(table, "TaxaEntrega")
	c.Manequim = field.NewString(table, "Manequim")

	c.fillFieldMap()

	return c
}

func (c *cliente) WithContext(ctx context.Context) *clienteDo { return c.clienteDo.WithContext(ctx) }

func (c cliente) TableName() string { return c.clienteDo.TableName() }

func (c cliente) Alias() string { return c.clienteDo.Alias() }

func (c cliente) Columns(cols ...field.Expr) gen.Columns { return c.clienteDo.Columns(cols...) }

func (c *cliente) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cliente) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 80)
	c.fieldMap["Codigo"] = c.Codigo
	c.fieldMap["Nome"] = c.Nome
	c.fieldMap["DataCadastro"] = c.DataCadastro
	c.fieldMap["RazaoSocial"] = c.RazaoSocial
	c.fieldMap["Limite"] = c.Limite
	c.fieldMap["Endereco"] = c.Endereco
	c.fieldMap["CEP"] = c.CEP
	c.fieldMap["Bairro"] = c.Bairro
	c.fieldMap["UF"] = c.UF
	c.fieldMap["Cidade"] = c.Cidade
	c.fieldMap["Fone"] = c.Fone
	c.fieldMap["Fax"] = c.Fax
	c.fieldMap["Contato"] = c.Contato
	c.fieldMap["CNPJCPF"] = c.CNPJCPF
	c.fieldMap["CGFRG"] = c.CGFRG
	c.fieldMap["Suframa"] = c.Suframa
	c.fieldMap["Email"] = c.Email
	c.fieldMap["Classificacao"] = c.Classificacao
	c.fieldMap["Vendedor"] = c.Vendedor
	c.fieldMap["Corretor"] = c.Corretor
	c.fieldMap["NomeResponsavel"] = c.NomeResponsavel
	c.fieldMap["NascimentoResponsavel"] = c.NascimentoResponsavel
	c.fieldMap["EnderecoResponsavel"] = c.EnderecoResponsavel
	c.fieldMap["CEPResponsavel"] = c.CEPResponsavel
	c.fieldMap["BairroResponsavel"] = c.BairroResponsavel
	c.fieldMap["UFResponsavel"] = c.UFResponsavel
	c.fieldMap["CidadeResponsavel"] = c.CidadeResponsavel
	c.fieldMap["FoneResponsavel"] = c.FoneResponsavel
	c.fieldMap["CPFResponsavel"] = c.CPFResponsavel
	c.fieldMap["RGResponsavel"] = c.RGResponsavel
	c.fieldMap["MaeResponsavel"] = c.MaeResponsavel
	c.fieldMap["PontoReferencia"] = c.PontoReferencia
	c.fieldMap["NomeTerceiro"] = c.NomeTerceiro
	c.fieldMap["NascimentoTerceiro"] = c.NascimentoTerceiro
	c.fieldMap["EnderecoTerceiro"] = c.EnderecoTerceiro
	c.fieldMap["CEPTerceiro"] = c.CEPTerceiro
	c.fieldMap["BairroTerceiro"] = c.BairroTerceiro
	c.fieldMap["UFTerceiro"] = c.UFTerceiro
	c.fieldMap["CidadeTerceiro"] = c.CidadeTerceiro
	c.fieldMap["FoneTerceiro"] = c.FoneTerceiro
	c.fieldMap["CPFTerceiro"] = c.CPFTerceiro
	c.fieldMap["RGTerceiro"] = c.RGTerceiro
	c.fieldMap["MaeTerceiro"] = c.MaeTerceiro
	c.fieldMap["Obs"] = c.Obs
	c.fieldMap["RefBanco1"] = c.RefBanco1
	c.fieldMap["FoneBanco1"] = c.FoneBanco1
	c.fieldMap["RefBanco2"] = c.RefBanco2
	c.fieldMap["FoneBanco2"] = c.FoneBanco2
	c.fieldMap["Bloqueado"] = c.Bloqueado
	c.fieldMap["Inativo"] = c.Inativo
	c.fieldMap["Fisica"] = c.Fisica
	c.fieldMap["DataCompra"] = c.DataCompra
	c.fieldMap["ValorCompra"] = c.ValorCompra
	c.fieldMap["Foto"] = c.Foto
	c.fieldMap["Atualizado"] = c.Atualizado
	c.fieldMap["ObsRepresentante"] = c.ObsRepresentante
	c.fieldMap["OBSPedido"] = c.OBSPedido
	c.fieldMap["DescontoEspecial"] = c.DescontoEspecial
	c.fieldMap["IsentodeCNPJ"] = c.IsentodeCNPJ
	c.fieldMap["QtdeLojas"] = c.QtdeLojas
	c.fieldMap["ValorCredito"] = c.ValorCredito
	c.fieldMap["CodigoLoja"] = c.CodigoLoja
	c.fieldMap["EnderecoCobranca"] = c.EnderecoCobranca
	c.fieldMap["CepCobranca"] = c.CepCobranca
	c.fieldMap["BairroCobranca"] = c.BairroCobranca
	c.fieldMap["UFCobranca"] = c.UFCobranca
	c.fieldMap["CidadeCobranca"] = c.CidadeCobranca
	c.fieldMap["FoneCobranca"] = c.FoneCobranca
	c.fieldMap["VendedorLoja"] = c.VendedorLoja
	c.fieldMap["ChequeNaLoja"] = c.ChequeNaLoja
	c.fieldMap["CodigoPais"] = c.CodigoPais
	c.fieldMap["RGCliente"] = c.RGCliente
	c.fieldMap["FormaPagtoPadrao"] = c.FormaPagtoPadrao
	c.fieldMap["CondPagtoPadrao"] = c.CondPagtoPadrao
	c.fieldMap["msgOrcamento"] = c.MsgOrcamento
	c.fieldMap["DiasTroca"] = c.DiasTroca
	c.fieldMap["GrupoEmpresarial"] = c.GrupoEmpresarial
	c.fieldMap["Numero"] = c.Numero
	c.fieldMap["TaxaEntrega"] = c.TaxaEntrega
	c.fieldMap["Manequim"] = c.Manequim
}

func (c cliente) clone(db *gorm.DB) cliente {
	c.clienteDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cliente) replaceDB(db *gorm.DB) cliente {
	c.clienteDo.ReplaceDB(db)
	return c
}

type clienteDo struct{ gen.DO }

func (c clienteDo) Debug() *clienteDo {
	return c.withDO(c.DO.Debug())
}

func (c clienteDo) WithContext(ctx context.Context) *clienteDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c clienteDo) ReadDB() *clienteDo {
	return c.Clauses(dbresolver.Read)
}

func (c clienteDo) WriteDB() *clienteDo {
	return c.Clauses(dbresolver.Write)
}

func (c clienteDo) Session(config *gorm.Session) *clienteDo {
	return c.withDO(c.DO.Session(config))
}

func (c clienteDo) Clauses(conds ...clause.Expression) *clienteDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c clienteDo) Returning(value interface{}, columns ...string) *clienteDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c clienteDo) Not(conds ...gen.Condition) *clienteDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c clienteDo) Or(conds ...gen.Condition) *clienteDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c clienteDo) Select(conds ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c clienteDo) Where(conds ...gen.Condition) *clienteDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c clienteDo) Order(conds ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c clienteDo) Distinct(cols ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c clienteDo) Omit(cols ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c clienteDo) Join(table schema.Tabler, on ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c clienteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *clienteDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c clienteDo) RightJoin(table schema.Tabler, on ...field.Expr) *clienteDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c clienteDo) Group(cols ...field.Expr) *clienteDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c clienteDo) Having(conds ...gen.Condition) *clienteDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c clienteDo) Limit(limit int) *clienteDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c clienteDo) Offset(offset int) *clienteDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c clienteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *clienteDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c clienteDo) Unscoped() *clienteDo {
	return c.withDO(c.DO.Unscoped())
}

func (c clienteDo) Create(values ...*model.Cliente) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c clienteDo) CreateInBatches(values []*model.Cliente, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c clienteDo) Save(values ...*model.Cliente) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c clienteDo) First() (*model.Cliente, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cliente), nil
	}
}

func (c clienteDo) Take() (*model.Cliente, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cliente), nil
	}
}

func (c clienteDo) Last() (*model.Cliente, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cliente), nil
	}
}

func (c clienteDo) Find() ([]*model.Cliente, error) {
	result, err := c.DO.Find()
	return result.([]*model.Cliente), err
}

func (c clienteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cliente, err error) {
	buf := make([]*model.Cliente, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c clienteDo) FindInBatches(result *[]*model.Cliente, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c clienteDo) Attrs(attrs ...field.AssignExpr) *clienteDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c clienteDo) Assign(attrs ...field.AssignExpr) *clienteDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c clienteDo) Joins(fields ...field.RelationField) *clienteDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c clienteDo) Preload(fields ...field.RelationField) *clienteDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c clienteDo) FirstOrInit() (*model.Cliente, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cliente), nil
	}
}

func (c clienteDo) FirstOrCreate() (*model.Cliente, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cliente), nil
	}
}

func (c clienteDo) FindByPage(offset int, limit int) (result []*model.Cliente, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c clienteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c clienteDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c clienteDo) Delete(models ...*model.Cliente) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *clienteDo) withDO(do gen.Dao) *clienteDo {
	c.DO = *do.(*gen.DO)
	return c
}
