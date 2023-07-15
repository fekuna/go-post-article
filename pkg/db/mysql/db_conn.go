package mysql

import (
	"fmt"
	"time"

	"github.com/fekuna/go-post-article/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpensConns   = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance
func NewMysqlDB(c *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s",
		c.Mysql.MysqlUsername,
		c.Mysql.MysqlPassword,
		c.Mysql.MysqlHost,
		c.Mysql.MysqlPort,
		c.Mysql.MysqlDbName,
	)

	db, err := sqlx.Connect(c.Mysql.Driver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpensConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
