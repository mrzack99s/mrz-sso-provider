package tools

import (
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
)

var entities = []interface{}{
	&models.UserInformation{},
	&models.Authentication{},
}

func SQL_Migrate() {

	for _, entity := range entities {
		if !sql_db.SQL_DB.Migrator().HasTable(entity) {
			sql_db.SQL_DB.Migrator().CreateTable(entity)
		}
	}

}

func SQL_CleanAndMigrate() {

	for _, entity := range entities {
		sql_db.SQL_DB.Migrator().DropTable(entity)
		sql_db.SQL_DB.Migrator().CreateTable(entity)
	}

}
