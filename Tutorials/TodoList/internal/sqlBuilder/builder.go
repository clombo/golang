package sqlbuilder

import (
	"database/sql"
	"reflect"
)

//TODO: As you get better use interfaces instead so that it can be one struct and choosen according to the operation
//like SELECT,CREATE,DELETE,UPDATE

/*
	builder := sqlbuilder.NewBuilder()

	SELECT
	builder
		.Select(struct here)
		.From(table here) //optional
		.Where(where clause here with args)
		.ExecuteOne() OR .ExecuteMany()

	CREATE TABLE
	builder.CreateTable(struct here)

	ANY
	builder
		.HasAny(pass ref to bool)
		.From(either a struct or a string) //This is required!
		.Where() //This is required for the Any to work

	//INSERT
	builder
		.Add(new record here) //If Into is not specified use the struct passed here
		.Into(either a struct or a string) //This is optional

	//DELETE
	builder
		.Delete(struct or string of table name)
		.Where() //This is required for delete to work

	builder
		.DeleteBulk(struct or string of table name)

	Documentation:
		Document each parts usages
		Document tags

	Future plans:
	- Check different ways how to build strings efficiently
	- Add tags for structs for tings like PRIMARY KEY, FOREIGN KEY, UNIQUE
	- Get insparation from other ORM's if exists
	- Get insparation from other Linq like modules
	- Possibly investigate context like struct holding db sets
	- Add logging support like EfCore has where you can see the query built when enabled
	- Add SP support
	- Add transaction support
	- Add repository and UnitOfWork pattern support
	- Add database support for MsSql,PostgreSql,MySql
*/

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
