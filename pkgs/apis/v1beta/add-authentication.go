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

type APIAddAuthentication struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Attribute  string `json:"attribute"`
	PersonalID string `json:"personal_id"`
}

func AddAuthentication(c *gin.Context) {

	var cipherPayload APICipherPayload
	if err := c.ShouldBindJSON(&cipherPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input APIAddAuthentication
	json.Unmarshal([]byte(security.Decrypt([]byte(cipherPayload.Payload))), &input)

	userInfo := &models.UserInformation{
		PersonalID: input.PersonalID,
	}
	err := repository.FindUserByPersonalId(userInfo)

	if err == nil {

		newUUID := ""
		for {
			newUUID = system.NewUUID()
			if !repository.FindUniqueAuthID(newUUID) {
				break
			}
		}

		newAuthentication := models.Authentication{
			ID:        newUUID,
			Username:  input.Username,
			Password:  input.Password,
			Attribute: input.Attribute,
			UserID:    userInfo.ID,
		}
		err = newAuthentication.Create()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": newAuthentication,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

}
