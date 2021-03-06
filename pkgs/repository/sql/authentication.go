package repository

import (
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
)

func GetAuthentication(auth *models.Authentication) error {
	result := sql_db.SQL_DB.Where("username = ? AND password = ?", auth.Username, auth.Password).First(auth)
	return result.Error
}

func FindUniqueAuthID(id string) bool {
	auth := &models.Authentication{}
	sql_db.SQL_DB.Where("auth_id = ?", id).First(auth)

	return auth.ID != ""
}
