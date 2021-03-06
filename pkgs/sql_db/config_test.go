package sql_db_test

import (
	"testing"

	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
	"github.com/stretchr/testify/assert"
)

func TestConfigInitial(t *testing.T) {
	system.ParseSystemConfig("../../config.yaml")
	db := sql_db.MySQL{
		Username: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Username,
		Password: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Password,
		Hostname: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Hostname,
		DBName:   system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.DBName,
	}
	db.Initial()

	objDB, _ := sql_db.SQL_DB.DB()
	assert.Nil(t, objDB.Ping())
	objDB.Close()
}
