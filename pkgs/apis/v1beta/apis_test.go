package apis_test

import (
	"testing"

	apis "github.com/mrzack99s/mrz-sso-provider/pkgs/apis/v1beta"
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	repository "github.com/mrzack99s/mrz-sso-provider/pkgs/repository/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/tools"
	"github.com/stretchr/testify/assert"
)

func TestAddUserinfo(t *testing.T) {

	system.ParseSystemConfig("../../../config.yaml")
	db := sql_db.MySQL{
		Username: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Username,
		Password: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Password,
		Hostname: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Hostname,
		DBName:   system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.DBName,
	}
	db.Initial()
	tools.SQL_CleanAndMigrate()

	//data from client
	//raw Data
	userInfo := apis.APIUserInformation{
		FirstName:  "firstname",
		LastName:   "lastname",
		PersonalID: "1234567890",
	}

	//generate uuid
	newUUID := ""
	for {
		newUUID = system.NewUUID()
		if !repository.FindUniqueUserID(newUUID) {
			break
		}
	}

	newUserInfo := models.UserInformation{
		ID:         newUUID,
		PersonalID: userInfo.PersonalID,
		FirstName:  userInfo.FirstName,
		LastName:   userInfo.LastName,
	}

	//create user
	err := newUserInfo.Create()
	assert.Nil(t, err)

	//delete user
	err = newUserInfo.Delete()
	assert.Nil(t, err)

	//stop db
	objDB, _ := sql_db.SQL_DB.DB()
	objDB.Close()

}

func TestAddAuthentication(t *testing.T) {

	system.ParseSystemConfig("../../../config.yaml")
	db := sql_db.MySQL{
		Username: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Username,
		Password: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Password,
		Hostname: system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.Hostname,
		DBName:   system.SystemConfigVar.MRZ_SSO.DB.SQL.UAT.DBName,
	}
	db.Initial()
	tools.SQL_CleanAndMigrate()

	//data from client
	//raw Data
	userInfo := apis.APIUserInformation{
		FirstName:  "firstname",
		LastName:   "lastname",
		PersonalID: "1234567890",
	}

	//generate uuid
	newUUID := ""
	for {
		newUUID = system.NewUUID()
		if !repository.FindUniqueUserID(newUUID) {
			break
		}
	}

	newUserInfo := models.UserInformation{
		ID:         newUUID,
		PersonalID: userInfo.PersonalID,
		FirstName:  userInfo.FirstName,
		LastName:   userInfo.LastName,
	}

	//create user
	err := newUserInfo.Create()
	assert.Nil(t, err)

	//authentication
	auth_input := &apis.APIAddAuthentication{
		Username:   "test",
		Password:   "password",
		Attribute:  "auth;",
		PersonalID: newUserInfo.PersonalID,
	}

	findUserInfo := &models.UserInformation{
		PersonalID: auth_input.PersonalID,
	}
	//find user
	err = repository.FindUserByPersonalId(findUserInfo)
	assert.Nil(t, err)

	//generate uuid
	newUUID = ""
	for {
		newUUID = system.NewUUID()
		if !repository.FindUniqueAuthID(newUUID) {
			break
		}
	}

	newAuthentication := models.Authentication{
		ID:        newUUID,
		Username:  auth_input.Username,
		Password:  auth_input.Password,
		Attribute: auth_input.Attribute,
		UserID:    findUserInfo.ID,
	}
	//create authentication
	err = newAuthentication.Create()
	assert.Nil(t, err)

	//delete user
	err = newUserInfo.Delete()
	assert.Nil(t, err)

	//close db
	objDB, _ := sql_db.SQL_DB.DB()
	objDB.Close()

}
