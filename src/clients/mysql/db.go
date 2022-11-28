package mysql

import (
	"podcast/src/zlog"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Client DBIFace
)

func ConnectToDatabase(dbUser, dbPass, dbHost, dbPort, dbName, useTLS string) *sqlx.DB {
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&rejectReadOnly=true&tls=%s", dbUser, dbPass, dbHost, dbPort, dbName, useTLS)
	client, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		zlog.Logger.Panic(fmt.Sprintf("database connection failed: %s", err))
	}

	if err = client.Ping(); err != nil {
		zlog.Logger.Panic(fmt.Sprintf("database ping failed: %s", err))
	}
	client.SetMaxIdleConns(50)
	client.SetMaxOpenConns(0)
	client.SetConnMaxLifetime(time.Minute * 10)
	client.SetConnMaxIdleTime(time.Minute * 10)

	Client = client

	return client
}

type DBIFace interface {
	DriverName() string
	MapperFunc(mf func(string) string)
	Rebind(query string) string
	Unsafe() *sqlx.DB
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	MustBegin() *sqlx.Tx
	Beginx() (*sqlx.Tx, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	MustExec(query string, args ...interface{}) sql.Result
	Preparex(query string) (*sqlx.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sqlx.Tx
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	Connx(ctx context.Context) (*sqlx.Conn, error)

	PingContext(ctx context.Context) error
	Ping() error
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Stats() sql.DBStats
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	Driver() driver.Driver
	Conn(ctx context.Context) (*sql.Conn, error)
}
