package models

import (
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
)

type UserInformation struct {
	ID         string           `gorm:"primaryKey;size:36;column:user_id;not null"`
	PersonalID string           `gorm:"column:personal_id;size:20;uniqueIndex;not null"`
	FirstName  string           `gorm:"column:first_name;size:255;not null"`
	LastName   string           `gorm:"column:last_name;size:255;not null"`
	Auths      []Authentication `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

//Create --
func (obj *UserInformation) Create() error {
	result := sql_db.SQL_DB.Create(obj)
	return result.Error
}

//Update
func (obj *UserInformation) UpdateFirstName() error {
	result := sql_db.SQL_DB.Update("first_name", obj)
	return result.Error
}

func (obj *UserInformation) UpdateLastName() error {
	result := sql_db.SQL_DB.Update("last_name", obj)
	return result.Error
}

//Delete --
func (obj *UserInformation) Delete() error {
	result := sql_db.SQL_DB.Delete(obj)
	return result.Error
}

func (obj *UserInformation) DeleteFromFirstName() error {
	result := sql_db.SQL_DB.Where("first_name = ?", obj.FirstName).Delete(obj)
	return result.Error
}

func (obj *UserInformation) DeleteFromLastName() error {
	result := sql_db.SQL_DB.Where("last_name = ?", obj.LastName).Delete(obj)
	return result.Error
}
