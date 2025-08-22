package sqlbuilder

import (
	"database/sql"
	"reflect"
)

//TODO: As you get better use interfaces instead so that it can be one struct and choosen according to the operation
//like SELECT,CREATE,DELETE,UPDATE

// struct that holds everything for SELECT
type QueryBuilder struct {
	db          *sql.DB
	selectCols  []string
	tableName   string
	whereClause string
	whereArgs   []any
	structType  reflect.Type
}

// Init new query builder
func NewQueryBuilder(db *sql.DB) *QueryBuilder {
	return &QueryBuilder{
		db:        db,
		whereArgs: make([]interface{}, 0),
	}
}

//

// struct that holds everything for CREATE
type CreateBuilder struct {
	db         *sql.DB
	tableName  string
	structType reflect.Type
}

// struct that holds everything for UPDATE
type UpdateBuilder struct {
	db           *sql.DB
	updateClause string
	updateArgs   []any
	whereClause  string
	whereArgs    []any
	tableName    string
	structType   reflect.Type
}

// struct that holds everything for DELETE
type DeleteBuilder struct {
	db          *sql.DB
	tableName   string
	structType  reflect.Type
	whereClause string
	whereArgs   []any
}
