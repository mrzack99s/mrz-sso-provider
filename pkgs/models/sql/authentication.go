package models

import (
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
)

type Authentication struct {
	ID        string `gorm:"primaryKey;size:36;column:auth_id;not null"`
	Username  string `gorm:"column:username;size:100;uniqueIndex"`
	Password  string `gorm:"column:password;size:255;not null"`
	Attribute string `gorm:"column:attribute;size:100;default:no-action"`
	UserID    string `gorm:"column:user_id;size:36;not null"`
}

//Create --
func (obj *Authentication) Create() error {
	result := sql_db.SQL_DB.Create(obj)
	return result.Error
}

//Update
func (obj *Authentication) UpdateFirstName() error {
	result := sql_db.SQL_DB.Update("first_name", obj)
	return result.Error
}

func (obj *Authentication) UpdateLastName() error {
	result := sql_db.SQL_DB.Update("last_name", obj)
	return result.Error
}

//Delete --
func (obj *Authentication) Delete() error {
	result := sql_db.SQL_DB.Delete(obj)
	return result.Error
}
