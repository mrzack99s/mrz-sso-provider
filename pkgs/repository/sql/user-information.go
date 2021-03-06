package repository

import (
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
)

func FindUserByPersonalId(obj *models.UserInformation) error {
	result := sql_db.SQL_DB.First(obj)
	return result.Error
}

func FindUniqueUserID(id string) bool {
	user := &models.UserInformation{}
	sql_db.SQL_DB.Where("user_id = ?", id).First(user)

	return user.ID != ""
}
