package dddgen

const tmplProtocolDomainModel = `package domain

type {{.Model}} struct {
{{.Items}}
}

type {{.Model}}Repository interface {
	Save(dm *{{.Model}}) (*{{.Model}}, error)
	Remove(dm *{{.Model}}) (*{{.Model}}, error)
	FindOne(queryOptions map[string]interface{}) (*{{.Model}}, error)
	Find(queryOptions map[string]interface{}) (int64, []*{{.Model}}, error)
}

func (m *{{.Model}}) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
`

const tmplProtocolDomainPgRepository = `package repository

type {{.Model}}Repository struct {
	transactionContext *transaction.TransactionContext
}

func (repository *{{.Model}}Repository) Save(dm *domain.{{.Model}}) (*domain.{{.Model}}, error) {
	var (
		err error
		m   = &models.{{.Model}}{}
		tx  = repository.transactionContext.PgTx
	)
	if err = GobModelTransform(m, dm); err != nil {
		return nil, err
	}
	if dm.Identify() == nil {
		if err = tx.Insert(m); err != nil {
			return nil, err
		}
		return dm, nil
	}
	if err = tx.Update(m); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *{{.Model}}Repository) Remove({{.Model}} *domain.{{.Model}}) (*domain.{{.Model}}, error) {
	var (
		tx          = repository.transactionContext.PgTx
		{{.Model}}Model = &models.{{.Model}}{Id: {{.Model}}.Identify().(int64)}
	)
	if _, err := tx.Model({{.Model}}Model).Where("id = ?", {{.Model}}.Id).Delete(); err != nil {
		return {{.Model}}, err
	}
	return {{.Model}}, nil
}

func (repository *{{.Model}}Repository) FindOne(queryOptions map[string]interface{}) (*domain.{{.Model}}, error) {
	tx := repository.transactionContext.PgTx
	{{.Model}}Model := new(models.{{.Model}})
	query := NewQuery(tx.Model({{.Model}}Model), queryOptions)
	query.SetWhere("id = ?", "id")
	if err := query.First(); err != nil {
		return nil, domain.QueryNoRow
	}
	if {{.Model}}Model.Id == 0 {
		return nil, domain.QueryNoRow
	}
	return repository.transformPgModelToDomainModel({{.Model}}Model)
}

func (repository *{{.Model}}Repository) Find(queryOptions map[string]interface{}) (int64, []*domain.{{.Model}}, error) {
	tx := repository.transactionContext.PgTx
	var {{.Model}}Models []*models.{{.Model}}
	{{.Model}}s := make([]*domain.{{.Model}}, 0)
	query := NewQuery(tx.Model(&{{.Model}}Models), queryOptions).
		SetOrder("create_time", "sortByCreateTime").
		SetOrder("update_time", "sortByUpdateTime")
	var err error
	if query.AffectRow, err = query.SelectAndCount(); err != nil {
		return 0, {{.Model}}s, err
	}
	for _, {{.Model}}Model := range {{.Model}}Models {
		if {{.Model}}, err := repository.transformPgModelToDomainModel({{.Model}}Model); err != nil {
			return 0, {{.Model}}s, err
		} else {
			{{.Model}}s = append({{.Model}}s, {{.Model}})
		}
	}
	return int64(query.AffectRow), {{.Model}}s, nil
}

func (repository *{{.Model}}Repository) transformPgModelToDomainModel({{.Model}}Model *models.{{.Model}}) (*domain.{{.Model}}, error) {
	m := &domain.{{.Model}}{}
	err := GobModelTransform(m, {{.Model}}Model)
	return m, err
}

func New{{.Model}}Repository(transactionContext *transaction.TransactionContext) (*{{.Model}}Repository, error) {
	if transactionContext == nil {
		return nil,fmt.Errorf("transactionContext参数不能为nil")
	}
	return &{{.Model}}Repository{transactionContext: transactionContext}, nil
}
`

const tmplProtocolPgModel = `package models

type {{.Model}} struct {
{{.Items}}
}
`
