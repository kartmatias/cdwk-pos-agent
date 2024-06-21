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

func newEmpresa(db *gorm.DB, opts ...gen.DOOption) empresa {
	_empresa := empresa{}

	_empresa.empresaDo.UseDB(db, opts...)
	_empresa.empresaDo.UseModel(&model.Empresa{})

	tableName := _empresa.empresaDo.TableName()
	_empresa.ALL = field.NewAsterisk(tableName)
	_empresa.Codigo = field.NewInt64(tableName, "Codigo")
	_empresa.RazaoSocial = field.NewString(tableName, "RazaoSocial")
	_empresa.Fantasia = field.NewString(tableName, "Fantasia")
	_empresa.Endereco = field.NewString(tableName, "Endereco")
	_empresa.CEP = field.NewString(tableName, "CEP")
	_empresa.Bairro = field.NewString(tableName, "Bairro")
	_empresa.UF = field.NewString(tableName, "UF")
	_empresa.Cidade = field.NewString(tableName, "Cidade")
	_empresa.CNPJCPF = field.NewString(tableName, "CNPJCPF")
	_empresa.CGFRG = field.NewString(tableName, "CGFRG")
	_empresa.Email = field.NewString(tableName, "Email")
	_empresa.Fone = field.NewString(tableName, "Fone")
	_empresa.Fax = field.NewString(tableName, "Fax")
	_empresa.ImprimirDuplicatacomoNumerodaNF = field.NewBool(tableName, "ImprimirDuplicatacomoNumerodaNF")
	_empresa.CNPJResponsavelTecnico = field.NewString(tableName, "CNPJResponsavelTecnico")
	_empresa.ContatoResponsavelTecnico = field.NewString(tableName, "ContatoResponsavelTecnico")
	_empresa.EmailResponsavelTecnico = field.NewString(tableName, "EmailResponsavelTecnico")
	_empresa.FoneResponsavelTecnico = field.NewString(tableName, "FoneResponsavelTecnico")

	_empresa.fillFieldMap()

	return _empresa
}

type empresa struct {
	empresaDo empresaDo

	ALL                             field.Asterisk
	Codigo                          field.Int64
	RazaoSocial                     field.String
	Fantasia                        field.String
	Endereco                        field.String
	CEP                             field.String
	Bairro                          field.String
	UF                              field.String
	Cidade                          field.String
	CNPJCPF                         field.String
	CGFRG                           field.String
	Email                           field.String
	Fone                            field.String
	Fax                             field.String
	ImprimirDuplicatacomoNumerodaNF field.Bool
	CNPJResponsavelTecnico          field.String
	ContatoResponsavelTecnico       field.String
	EmailResponsavelTecnico         field.String
	FoneResponsavelTecnico          field.String

	fieldMap map[string]field.Expr
}

func (e empresa) Table(newTableName string) *empresa {
	e.empresaDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e empresa) As(alias string) *empresa {
	e.empresaDo.DO = *(e.empresaDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *empresa) updateTableName(table string) *empresa {
	e.ALL = field.NewAsterisk(table)
	e.Codigo = field.NewInt64(table, "Codigo")
	e.RazaoSocial = field.NewString(table, "RazaoSocial")
	e.Fantasia = field.NewString(table, "Fantasia")
	e.Endereco = field.NewString(table, "Endereco")
	e.CEP = field.NewString(table, "CEP")
	e.Bairro = field.NewString(table, "Bairro")
	e.UF = field.NewString(table, "UF")
	e.Cidade = field.NewString(table, "Cidade")
	e.CNPJCPF = field.NewString(table, "CNPJCPF")
	e.CGFRG = field.NewString(table, "CGFRG")
	e.Email = field.NewString(table, "Email")
	e.Fone = field.NewString(table, "Fone")
	e.Fax = field.NewString(table, "Fax")
	e.ImprimirDuplicatacomoNumerodaNF = field.NewBool(table, "ImprimirDuplicatacomoNumerodaNF")
	e.CNPJResponsavelTecnico = field.NewString(table, "CNPJResponsavelTecnico")
	e.ContatoResponsavelTecnico = field.NewString(table, "ContatoResponsavelTecnico")
	e.EmailResponsavelTecnico = field.NewString(table, "EmailResponsavelTecnico")
	e.FoneResponsavelTecnico = field.NewString(table, "FoneResponsavelTecnico")

	e.fillFieldMap()

	return e
}

func (e *empresa) WithContext(ctx context.Context) *empresaDo { return e.empresaDo.WithContext(ctx) }

func (e empresa) TableName() string { return e.empresaDo.TableName() }

func (e empresa) Alias() string { return e.empresaDo.Alias() }

func (e empresa) Columns(cols ...field.Expr) gen.Columns { return e.empresaDo.Columns(cols...) }

func (e *empresa) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *empresa) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 18)
	e.fieldMap["Codigo"] = e.Codigo
	e.fieldMap["RazaoSocial"] = e.RazaoSocial
	e.fieldMap["Fantasia"] = e.Fantasia
	e.fieldMap["Endereco"] = e.Endereco
	e.fieldMap["CEP"] = e.CEP
	e.fieldMap["Bairro"] = e.Bairro
	e.fieldMap["UF"] = e.UF
	e.fieldMap["Cidade"] = e.Cidade
	e.fieldMap["CNPJCPF"] = e.CNPJCPF
	e.fieldMap["CGFRG"] = e.CGFRG
	e.fieldMap["Email"] = e.Email
	e.fieldMap["Fone"] = e.Fone
	e.fieldMap["Fax"] = e.Fax
	e.fieldMap["ImprimirDuplicatacomoNumerodaNF"] = e.ImprimirDuplicatacomoNumerodaNF
	e.fieldMap["CNPJResponsavelTecnico"] = e.CNPJResponsavelTecnico
	e.fieldMap["ContatoResponsavelTecnico"] = e.ContatoResponsavelTecnico
	e.fieldMap["EmailResponsavelTecnico"] = e.EmailResponsavelTecnico
	e.fieldMap["FoneResponsavelTecnico"] = e.FoneResponsavelTecnico
}

func (e empresa) clone(db *gorm.DB) empresa {
	e.empresaDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e empresa) replaceDB(db *gorm.DB) empresa {
	e.empresaDo.ReplaceDB(db)
	return e
}

type empresaDo struct{ gen.DO }

func (e empresaDo) Debug() *empresaDo {
	return e.withDO(e.DO.Debug())
}

func (e empresaDo) WithContext(ctx context.Context) *empresaDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e empresaDo) ReadDB() *empresaDo {
	return e.Clauses(dbresolver.Read)
}

func (e empresaDo) WriteDB() *empresaDo {
	return e.Clauses(dbresolver.Write)
}

func (e empresaDo) Session(config *gorm.Session) *empresaDo {
	return e.withDO(e.DO.Session(config))
}

func (e empresaDo) Clauses(conds ...clause.Expression) *empresaDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e empresaDo) Returning(value interface{}, columns ...string) *empresaDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e empresaDo) Not(conds ...gen.Condition) *empresaDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e empresaDo) Or(conds ...gen.Condition) *empresaDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e empresaDo) Select(conds ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e empresaDo) Where(conds ...gen.Condition) *empresaDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e empresaDo) Order(conds ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e empresaDo) Distinct(cols ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e empresaDo) Omit(cols ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e empresaDo) Join(table schema.Tabler, on ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e empresaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *empresaDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e empresaDo) RightJoin(table schema.Tabler, on ...field.Expr) *empresaDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e empresaDo) Group(cols ...field.Expr) *empresaDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e empresaDo) Having(conds ...gen.Condition) *empresaDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e empresaDo) Limit(limit int) *empresaDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e empresaDo) Offset(offset int) *empresaDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e empresaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *empresaDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e empresaDo) Unscoped() *empresaDo {
	return e.withDO(e.DO.Unscoped())
}

func (e empresaDo) Create(values ...*model.Empresa) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e empresaDo) CreateInBatches(values []*model.Empresa, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e empresaDo) Save(values ...*model.Empresa) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e empresaDo) First() (*model.Empresa, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Empresa), nil
	}
}

func (e empresaDo) Take() (*model.Empresa, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Empresa), nil
	}
}

func (e empresaDo) Last() (*model.Empresa, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Empresa), nil
	}
}

func (e empresaDo) Find() ([]*model.Empresa, error) {
	result, err := e.DO.Find()
	return result.([]*model.Empresa), err
}

func (e empresaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Empresa, err error) {
	buf := make([]*model.Empresa, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e empresaDo) FindInBatches(result *[]*model.Empresa, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e empresaDo) Attrs(attrs ...field.AssignExpr) *empresaDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e empresaDo) Assign(attrs ...field.AssignExpr) *empresaDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e empresaDo) Joins(fields ...field.RelationField) *empresaDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e empresaDo) Preload(fields ...field.RelationField) *empresaDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e empresaDo) FirstOrInit() (*model.Empresa, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Empresa), nil
	}
}

func (e empresaDo) FirstOrCreate() (*model.Empresa, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Empresa), nil
	}
}

func (e empresaDo) FindByPage(offset int, limit int) (result []*model.Empresa, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e empresaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e empresaDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e empresaDo) Delete(models ...*model.Empresa) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *empresaDo) withDO(do gen.Dao) *empresaDo {
	e.DO = *do.(*gen.DO)
	return e
}
