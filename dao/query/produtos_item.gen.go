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

func newProdutosItem(db *gorm.DB, opts ...gen.DOOption) produtosItem {
	_produtosItem := produtosItem{}

	_produtosItem.produtosItemDo.UseDB(db, opts...)
	_produtosItem.produtosItemDo.UseModel(&model.ProdutosItem{})

	tableName := _produtosItem.produtosItemDo.TableName()
	_produtosItem.ALL = field.NewAsterisk(tableName)
	_produtosItem.Produto = field.NewString(tableName, "Produto")
	_produtosItem.Cor = field.NewInt64(tableName, "Cor")
	_produtosItem.Ent01 = field.NewInt64(tableName, "Ent01")
	_produtosItem.Ent02 = field.NewInt64(tableName, "Ent02")
	_produtosItem.Ent03 = field.NewInt64(tableName, "Ent03")
	_produtosItem.Ent04 = field.NewInt64(tableName, "Ent04")
	_produtosItem.Ent05 = field.NewInt64(tableName, "Ent05")
	_produtosItem.Ent06 = field.NewInt64(tableName, "Ent06")
	_produtosItem.Ent07 = field.NewInt64(tableName, "Ent07")
	_produtosItem.Ent08 = field.NewInt64(tableName, "Ent08")
	_produtosItem.Ent09 = field.NewInt64(tableName, "Ent09")
	_produtosItem.Ent10 = field.NewInt64(tableName, "Ent10")
	_produtosItem.Sai01 = field.NewInt64(tableName, "Sai01")
	_produtosItem.Sai02 = field.NewInt64(tableName, "Sai02")
	_produtosItem.Sai03 = field.NewInt64(tableName, "Sai03")
	_produtosItem.Sai04 = field.NewInt64(tableName, "Sai04")
	_produtosItem.Sai05 = field.NewInt64(tableName, "Sai05")
	_produtosItem.Sai06 = field.NewInt64(tableName, "Sai06")
	_produtosItem.Sai07 = field.NewInt64(tableName, "Sai07")
	_produtosItem.Sai08 = field.NewInt64(tableName, "Sai08")
	_produtosItem.Sai09 = field.NewInt64(tableName, "Sai09")
	_produtosItem.Sai10 = field.NewInt64(tableName, "Sai10")
	_produtosItem.Preco1 = field.NewField(tableName, "Preco1")
	_produtosItem.Preco2 = field.NewField(tableName, "Preco2")
	_produtosItem.Preco3 = field.NewField(tableName, "Preco3")
	_produtosItem.Preco4 = field.NewField(tableName, "Preco4")
	_produtosItem.Preco5 = field.NewField(tableName, "Preco5")
	_produtosItem.Preco6 = field.NewField(tableName, "Preco6")
	_produtosItem.DataBalanco = field.NewTime(tableName, "DataBalanco")
	_produtosItem.KitCor1 = field.NewInt64(tableName, "KitCor1")
	_produtosItem.KitCor2 = field.NewInt64(tableName, "KitCor2")
	_produtosItem.KitCor3 = field.NewInt64(tableName, "KitCor3")
	_produtosItem.KitCor4 = field.NewInt64(tableName, "KitCor4")
	_produtosItem.Inativo = field.NewBool(tableName, "Inativo")
	_produtosItem.PorcentagemCorte = field.NewFloat64(tableName, "PorcentagemCorte")
	_produtosItem.Imagem = field.NewString(tableName, "Imagem")

	_produtosItem.fillFieldMap()

	return _produtosItem
}

type produtosItem struct {
	produtosItemDo produtosItemDo

	ALL              field.Asterisk
	Produto          field.String
	Cor              field.Int64
	Ent01            field.Int64
	Ent02            field.Int64
	Ent03            field.Int64
	Ent04            field.Int64
	Ent05            field.Int64
	Ent06            field.Int64
	Ent07            field.Int64
	Ent08            field.Int64
	Ent09            field.Int64
	Ent10            field.Int64
	Sai01            field.Int64
	Sai02            field.Int64
	Sai03            field.Int64
	Sai04            field.Int64
	Sai05            field.Int64
	Sai06            field.Int64
	Sai07            field.Int64
	Sai08            field.Int64
	Sai09            field.Int64
	Sai10            field.Int64
	Preco1           field.Field
	Preco2           field.Field
	Preco3           field.Field
	Preco4           field.Field
	Preco5           field.Field
	Preco6           field.Field
	DataBalanco      field.Time
	KitCor1          field.Int64
	KitCor2          field.Int64
	KitCor3          field.Int64
	KitCor4          field.Int64
	Inativo          field.Bool
	PorcentagemCorte field.Float64
	Imagem           field.String

	fieldMap map[string]field.Expr
}

func (p produtosItem) Table(newTableName string) *produtosItem {
	p.produtosItemDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p produtosItem) As(alias string) *produtosItem {
	p.produtosItemDo.DO = *(p.produtosItemDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *produtosItem) updateTableName(table string) *produtosItem {
	p.ALL = field.NewAsterisk(table)
	p.Produto = field.NewString(table, "Produto")
	p.Cor = field.NewInt64(table, "Cor")
	p.Ent01 = field.NewInt64(table, "Ent01")
	p.Ent02 = field.NewInt64(table, "Ent02")
	p.Ent03 = field.NewInt64(table, "Ent03")
	p.Ent04 = field.NewInt64(table, "Ent04")
	p.Ent05 = field.NewInt64(table, "Ent05")
	p.Ent06 = field.NewInt64(table, "Ent06")
	p.Ent07 = field.NewInt64(table, "Ent07")
	p.Ent08 = field.NewInt64(table, "Ent08")
	p.Ent09 = field.NewInt64(table, "Ent09")
	p.Ent10 = field.NewInt64(table, "Ent10")
	p.Sai01 = field.NewInt64(table, "Sai01")
	p.Sai02 = field.NewInt64(table, "Sai02")
	p.Sai03 = field.NewInt64(table, "Sai03")
	p.Sai04 = field.NewInt64(table, "Sai04")
	p.Sai05 = field.NewInt64(table, "Sai05")
	p.Sai06 = field.NewInt64(table, "Sai06")
	p.Sai07 = field.NewInt64(table, "Sai07")
	p.Sai08 = field.NewInt64(table, "Sai08")
	p.Sai09 = field.NewInt64(table, "Sai09")
	p.Sai10 = field.NewInt64(table, "Sai10")
	p.Preco1 = field.NewField(table, "Preco1")
	p.Preco2 = field.NewField(table, "Preco2")
	p.Preco3 = field.NewField(table, "Preco3")
	p.Preco4 = field.NewField(table, "Preco4")
	p.Preco5 = field.NewField(table, "Preco5")
	p.Preco6 = field.NewField(table, "Preco6")
	p.DataBalanco = field.NewTime(table, "DataBalanco")
	p.KitCor1 = field.NewInt64(table, "KitCor1")
	p.KitCor2 = field.NewInt64(table, "KitCor2")
	p.KitCor3 = field.NewInt64(table, "KitCor3")
	p.KitCor4 = field.NewInt64(table, "KitCor4")
	p.Inativo = field.NewBool(table, "Inativo")
	p.PorcentagemCorte = field.NewFloat64(table, "PorcentagemCorte")
	p.Imagem = field.NewString(table, "Imagem")

	p.fillFieldMap()

	return p
}

func (p *produtosItem) WithContext(ctx context.Context) *produtosItemDo {
	return p.produtosItemDo.WithContext(ctx)
}

func (p produtosItem) TableName() string { return p.produtosItemDo.TableName() }

func (p produtosItem) Alias() string { return p.produtosItemDo.Alias() }

func (p produtosItem) Columns(cols ...field.Expr) gen.Columns {
	return p.produtosItemDo.Columns(cols...)
}

func (p *produtosItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *produtosItem) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 36)
	p.fieldMap["Produto"] = p.Produto
	p.fieldMap["Cor"] = p.Cor
	p.fieldMap["Ent01"] = p.Ent01
	p.fieldMap["Ent02"] = p.Ent02
	p.fieldMap["Ent03"] = p.Ent03
	p.fieldMap["Ent04"] = p.Ent04
	p.fieldMap["Ent05"] = p.Ent05
	p.fieldMap["Ent06"] = p.Ent06
	p.fieldMap["Ent07"] = p.Ent07
	p.fieldMap["Ent08"] = p.Ent08
	p.fieldMap["Ent09"] = p.Ent09
	p.fieldMap["Ent10"] = p.Ent10
	p.fieldMap["Sai01"] = p.Sai01
	p.fieldMap["Sai02"] = p.Sai02
	p.fieldMap["Sai03"] = p.Sai03
	p.fieldMap["Sai04"] = p.Sai04
	p.fieldMap["Sai05"] = p.Sai05
	p.fieldMap["Sai06"] = p.Sai06
	p.fieldMap["Sai07"] = p.Sai07
	p.fieldMap["Sai08"] = p.Sai08
	p.fieldMap["Sai09"] = p.Sai09
	p.fieldMap["Sai10"] = p.Sai10
	p.fieldMap["Preco1"] = p.Preco1
	p.fieldMap["Preco2"] = p.Preco2
	p.fieldMap["Preco3"] = p.Preco3
	p.fieldMap["Preco4"] = p.Preco4
	p.fieldMap["Preco5"] = p.Preco5
	p.fieldMap["Preco6"] = p.Preco6
	p.fieldMap["DataBalanco"] = p.DataBalanco
	p.fieldMap["KitCor1"] = p.KitCor1
	p.fieldMap["KitCor2"] = p.KitCor2
	p.fieldMap["KitCor3"] = p.KitCor3
	p.fieldMap["KitCor4"] = p.KitCor4
	p.fieldMap["Inativo"] = p.Inativo
	p.fieldMap["PorcentagemCorte"] = p.PorcentagemCorte
	p.fieldMap["Imagem"] = p.Imagem
}

func (p produtosItem) clone(db *gorm.DB) produtosItem {
	p.produtosItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p produtosItem) replaceDB(db *gorm.DB) produtosItem {
	p.produtosItemDo.ReplaceDB(db)
	return p
}

type produtosItemDo struct{ gen.DO }

func (p produtosItemDo) Debug() *produtosItemDo {
	return p.withDO(p.DO.Debug())
}

func (p produtosItemDo) WithContext(ctx context.Context) *produtosItemDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p produtosItemDo) ReadDB() *produtosItemDo {
	return p.Clauses(dbresolver.Read)
}

func (p produtosItemDo) WriteDB() *produtosItemDo {
	return p.Clauses(dbresolver.Write)
}

func (p produtosItemDo) Session(config *gorm.Session) *produtosItemDo {
	return p.withDO(p.DO.Session(config))
}

func (p produtosItemDo) Clauses(conds ...clause.Expression) *produtosItemDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p produtosItemDo) Returning(value interface{}, columns ...string) *produtosItemDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p produtosItemDo) Not(conds ...gen.Condition) *produtosItemDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p produtosItemDo) Or(conds ...gen.Condition) *produtosItemDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p produtosItemDo) Select(conds ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p produtosItemDo) Where(conds ...gen.Condition) *produtosItemDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p produtosItemDo) Order(conds ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p produtosItemDo) Distinct(cols ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p produtosItemDo) Omit(cols ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p produtosItemDo) Join(table schema.Tabler, on ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p produtosItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p produtosItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p produtosItemDo) Group(cols ...field.Expr) *produtosItemDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p produtosItemDo) Having(conds ...gen.Condition) *produtosItemDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p produtosItemDo) Limit(limit int) *produtosItemDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p produtosItemDo) Offset(offset int) *produtosItemDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p produtosItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *produtosItemDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p produtosItemDo) Unscoped() *produtosItemDo {
	return p.withDO(p.DO.Unscoped())
}

func (p produtosItemDo) Create(values ...*model.ProdutosItem) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p produtosItemDo) CreateInBatches(values []*model.ProdutosItem, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p produtosItemDo) Save(values ...*model.ProdutosItem) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p produtosItemDo) First() (*model.ProdutosItem, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosItem), nil
	}
}

func (p produtosItemDo) Take() (*model.ProdutosItem, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosItem), nil
	}
}

func (p produtosItemDo) Last() (*model.ProdutosItem, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosItem), nil
	}
}

func (p produtosItemDo) Find() ([]*model.ProdutosItem, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProdutosItem), err
}

func (p produtosItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProdutosItem, err error) {
	buf := make([]*model.ProdutosItem, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p produtosItemDo) FindInBatches(result *[]*model.ProdutosItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p produtosItemDo) Attrs(attrs ...field.AssignExpr) *produtosItemDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p produtosItemDo) Assign(attrs ...field.AssignExpr) *produtosItemDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p produtosItemDo) Joins(fields ...field.RelationField) *produtosItemDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p produtosItemDo) Preload(fields ...field.RelationField) *produtosItemDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p produtosItemDo) FirstOrInit() (*model.ProdutosItem, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosItem), nil
	}
}

func (p produtosItemDo) FirstOrCreate() (*model.ProdutosItem, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProdutosItem), nil
	}
}

func (p produtosItemDo) FindByPage(offset int, limit int) (result []*model.ProdutosItem, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p produtosItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p produtosItemDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p produtosItemDo) Delete(models ...*model.ProdutosItem) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *produtosItemDo) withDO(do gen.Dao) *produtosItemDo {
	p.DO = *do.(*gen.DO)
	return p
}
