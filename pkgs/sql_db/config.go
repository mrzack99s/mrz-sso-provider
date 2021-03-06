package sql_db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	SQL_DB *gorm.DB
	Err    error
)

type MySQL struct {
	Username string
	Password string
	Hostname string
	DBName   string
}

func (mdb *MySQL) Initial() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mdb.Username, mdb.Password, mdb.Hostname, mdb.DBName)
	SQL_DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sso_",
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	SQL_DB.DB()

}
