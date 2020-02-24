package infrastructure

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/Okinawas/blog_api/app/interfaces/database"
)

type sqlHandler struct {
  Conn *sql.DB
}

func NewSqlHandler() database.SqlHnadler {
  conn, err = sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/blog_dev")
  if err != nil {
    panic(err.Error)
  }

  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn

  return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
  res := SqlResult{}
  result, err := handler.Conn.Exec(statement, args...)

  if err != nil {
    return res, err
  }

  res.Result = result
  return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
  rows, err := handler.Conn.Query(statement, args...)
  if err != nil {
    return new(SqlRow), err
  }

  row := new(SqlRow)

  row.Rows = rows
  return row, nil
}

type SqlResult struct {
  Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
  return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
  return r.Result.RowAffected()
}

type SqlRow struct {
  Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
  return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
  return r.Rows.Next()
}


func (r SqlRow) Close() error {
  return r.Rows.Close()
}


