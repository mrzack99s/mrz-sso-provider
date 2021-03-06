package sql_db_test

import (
	"testing"

	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
	"github.com/stretchr/testify/assert"
)

// func TestConfigInitialMysql(t *testing.T) {
// 	system.ParseSystemConfig("../../config.yaml")
// 	db := sql_db.MySQL{
// 		Username: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Username,
// 		Password: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Password,
// 		Hostname: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Hostname,
// 		DBName:   system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.DBName,
// 	}
// 	db.Initial()

// 	objDB, _ := sql_db.SQL_DB.DB()
// 	assert.Nil(t, objDB.Ping())
// 	objDB.Close()
// }

func TestConfigInitialSQLLite(t *testing.T) {
	db := sql_db.SQLLite{
		Path: "../../test.db",
	}
	db.Initial()

	objDB, _ := sql_db.SQL_DB.DB()
	assert.Nil(t, objDB.Ping())
	objDB.Close()
}
