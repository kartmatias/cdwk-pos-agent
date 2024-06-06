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

func newSubGrupo(db *gorm.DB, opts ...gen.DOOption) subGrupo {
	_subGrupo := subGrupo{}

	_subGrupo.subGrupoDo.UseDB(db, opts...)
	_subGrupo.subGrupoDo.UseModel(&model.SubGrupo{})

	tableName := _subGrupo.subGrupoDo.TableName()
	_subGrupo.ALL = field.NewAsterisk(tableName)
	_subGrupo.Codigo = field.NewInt64(tableName, "Codigo")
	_subGrupo.Descricao = field.NewString(tableName, "Descricao")
	_subGrupo.Grupo = field.NewInt64(tableName, "Grupo")

	_subGrupo.fillFieldMap()

	return _subGrupo
}

type subGrupo struct {
	subGrupoDo subGrupoDo

	ALL       field.Asterisk
	Codigo    field.Int64
	Descricao field.String
	Grupo     field.Int64

	fieldMap map[string]field.Expr
}

func (s subGrupo) Table(newTableName string) *subGrupo {
	s.subGrupoDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subGrupo) As(alias string) *subGrupo {
	s.subGrupoDo.DO = *(s.subGrupoDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subGrupo) updateTableName(table string) *subGrupo {
	s.ALL = field.NewAsterisk(table)
	s.Codigo = field.NewInt64(table, "Codigo")
	s.Descricao = field.NewString(table, "Descricao")
	s.Grupo = field.NewInt64(table, "Grupo")

	s.fillFieldMap()

	return s
}

func (s *subGrupo) WithContext(ctx context.Context) *subGrupoDo { return s.subGrupoDo.WithContext(ctx) }

func (s subGrupo) TableName() string { return s.subGrupoDo.TableName() }

func (s subGrupo) Alias() string { return s.subGrupoDo.Alias() }

func (s subGrupo) Columns(cols ...field.Expr) gen.Columns { return s.subGrupoDo.Columns(cols...) }

func (s *subGrupo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subGrupo) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["Codigo"] = s.Codigo
	s.fieldMap["Descricao"] = s.Descricao
	s.fieldMap["Grupo"] = s.Grupo
}

func (s subGrupo) clone(db *gorm.DB) subGrupo {
	s.subGrupoDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subGrupo) replaceDB(db *gorm.DB) subGrupo {
	s.subGrupoDo.ReplaceDB(db)
	return s
}

type subGrupoDo struct{ gen.DO }

func (s subGrupoDo) Debug() *subGrupoDo {
	return s.withDO(s.DO.Debug())
}

func (s subGrupoDo) WithContext(ctx context.Context) *subGrupoDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subGrupoDo) ReadDB() *subGrupoDo {
	return s.Clauses(dbresolver.Read)
}

func (s subGrupoDo) WriteDB() *subGrupoDo {
	return s.Clauses(dbresolver.Write)
}

func (s subGrupoDo) Session(config *gorm.Session) *subGrupoDo {
	return s.withDO(s.DO.Session(config))
}

func (s subGrupoDo) Clauses(conds ...clause.Expression) *subGrupoDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subGrupoDo) Returning(value interface{}, columns ...string) *subGrupoDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subGrupoDo) Not(conds ...gen.Condition) *subGrupoDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subGrupoDo) Or(conds ...gen.Condition) *subGrupoDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subGrupoDo) Select(conds ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subGrupoDo) Where(conds ...gen.Condition) *subGrupoDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subGrupoDo) Order(conds ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subGrupoDo) Distinct(cols ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subGrupoDo) Omit(cols ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subGrupoDo) Join(table schema.Tabler, on ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subGrupoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subGrupoDo) RightJoin(table schema.Tabler, on ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subGrupoDo) Group(cols ...field.Expr) *subGrupoDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subGrupoDo) Having(conds ...gen.Condition) *subGrupoDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subGrupoDo) Limit(limit int) *subGrupoDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subGrupoDo) Offset(offset int) *subGrupoDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subGrupoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subGrupoDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subGrupoDo) Unscoped() *subGrupoDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subGrupoDo) Create(values ...*model.SubGrupo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subGrupoDo) CreateInBatches(values []*model.SubGrupo, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subGrupoDo) Save(values ...*model.SubGrupo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subGrupoDo) First() (*model.SubGrupo, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubGrupo), nil
	}
}

func (s subGrupoDo) Take() (*model.SubGrupo, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubGrupo), nil
	}
}

func (s subGrupoDo) Last() (*model.SubGrupo, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubGrupo), nil
	}
}

func (s subGrupoDo) Find() ([]*model.SubGrupo, error) {
	result, err := s.DO.Find()
	return result.([]*model.SubGrupo), err
}

func (s subGrupoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SubGrupo, err error) {
	buf := make([]*model.SubGrupo, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subGrupoDo) FindInBatches(result *[]*model.SubGrupo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subGrupoDo) Attrs(attrs ...field.AssignExpr) *subGrupoDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subGrupoDo) Assign(attrs ...field.AssignExpr) *subGrupoDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subGrupoDo) Joins(fields ...field.RelationField) *subGrupoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subGrupoDo) Preload(fields ...field.RelationField) *subGrupoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subGrupoDo) FirstOrInit() (*model.SubGrupo, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubGrupo), nil
	}
}

func (s subGrupoDo) FirstOrCreate() (*model.SubGrupo, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubGrupo), nil
	}
}

func (s subGrupoDo) FindByPage(offset int, limit int) (result []*model.SubGrupo, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subGrupoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subGrupoDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subGrupoDo) Delete(models ...*model.SubGrupo) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subGrupoDo) withDO(do gen.Dao) *subGrupoDo {
	s.DO = *do.(*gen.DO)
	return s
}
