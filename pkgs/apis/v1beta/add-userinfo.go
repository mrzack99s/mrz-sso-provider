package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	repository "github.com/mrzack99s/mrz-sso-provider/pkgs/repository/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/security"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
)

type APIUserInformation struct {
	PersonalID string `json:"personal_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

func AddUserInformation(c *gin.Context) {

	var cipherPayload APICipherPayload
	if err := c.ShouldBindJSON(&cipherPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input APIUserInformation
	json.Unmarshal([]byte(security.Decrypt([]byte(cipherPayload.Payload))), &input)

	newUUID := ""
	for {
		newUUID = system.NewUUID()
		if !repository.FindUniqueUserID(newUUID) {
			break
		}
	}

	newUserInfo := models.UserInformation{
		ID:         newUUID,
		PersonalID: input.PersonalID,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
	}
	err := newUserInfo.Create()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": newUserInfo,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
