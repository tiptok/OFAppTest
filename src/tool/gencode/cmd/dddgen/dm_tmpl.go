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
		return nil, fmt.Errorf("query row not found")
	}
	if {{.Model}}Model.Id == 0 {
		return nil, fmt.Errorf("query row not found")
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

const tmplConstantPg = `package constant

import "os"

var POSTGRESQL_DB_NAME = "postgres"
var POSTGRESQL_USER = "postgres"      
var POSTGRESQL_PASSWORD = "123456"  
var POSTGRESQL_HOST = "127.0.0.1"  
var POSTGRESQL_PORT = "5432"          
var DISABLE_CREATE_TABLE = false
var DISABLE_SQL_GENERATE_PRINT = false

func init() {
	if os.Getenv("POSTGRESQL_DB_NAME") != "" {
		POSTGRESQL_DB_NAME = os.Getenv("POSTGRESQL_DB_NAME")
	}
	if os.Getenv("POSTGRESQL_USER") != "" {
		POSTGRESQL_USER = os.Getenv("POSTGRESQL_USER")
	}
	if os.Getenv("POSTGRESQL_PASSWORD") != "" {
		POSTGRESQL_PASSWORD = os.Getenv("POSTGRESQL_PASSWORD")
	}
	if os.Getenv("POSTGRESQL_HOST") != "" {
		POSTGRESQL_HOST = os.Getenv("POSTGRESQL_HOST")
	}
	if os.Getenv("POSTGRESQL_PORT") != "" {
		POSTGRESQL_PORT = os.Getenv("POSTGRESQL_PORT")
	}
	if os.Getenv("DISABLE_CREATE_TABLE") != "" {
		DISABLE_CREATE_TABLE = true
	}
	if os.Getenv("DISABLE_SQL_GENERATE_PRINT") != "" {
		DISABLE_SQL_GENERATE_PRINT = true
	}
}
`

const tmplPgInit = `package pg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     constant.POSTGRESQL_USER,
		Password: constant.POSTGRESQL_PASSWORD,
		Database: constant.POSTGRESQL_DB_NAME,
		Addr:     fmt.Sprintf("%s:%s", constant.POSTGRESQL_HOST, constant.POSTGRESQL_PORT),
	})
	if !constant.DISABLE_SQL_GENERATE_PRINT {
		DB.AddQueryHook(SqlGeneratePrintHook{})
	}
	//orm.RegisterTable((*models.OrderGood)(nil))
	if !constant.DISABLE_CREATE_TABLE {
		for _, model := range []interface{}{
{{.models}}
		} {
			err := DB.CreateTable(model, &orm.CreateTableOptions{
				Temp:          false,
				IfNotExists:   true,
				FKConstraints: true,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

type SqlGeneratePrintHook struct{}

func (hook SqlGeneratePrintHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (hook SqlGeneratePrintHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	data, err := q.FormattedQuery()
	if len(string(data)) > 8 { //BEGIN COMMIT
		//log.Debug(string(data))
	}
	return err
}

`

const tmplPgTransaction = `package transaction

import "github.com/go-pg/pg/v10"

type TransactionContext struct {
	PgDd *pg.DB
	PgTx *pg.Tx
}

func (transactionContext *TransactionContext) StartTransaction() error {
	tx, err := transactionContext.PgDd.Begin()
	if err != nil {
		return err
	}
	transactionContext.PgTx = tx
	return nil
}

func (transactionContext *TransactionContext) CommitTransaction() error {
	err := transactionContext.PgTx.Commit()
	return err
}

func (transactionContext *TransactionContext) RollbackTransaction() error {
	err := transactionContext.PgTx.Rollback()
	return err
}

func NewPGTransactionContext(pgDd *pg.DB) *TransactionContext {
	return &TransactionContext{
		PgDd: pgDd,
	}
}
`
