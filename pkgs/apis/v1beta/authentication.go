package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mrzack99s/mrz-sso-provider/pkgs/models/sql"
	repository "github.com/mrzack99s/mrz-sso-provider/pkgs/repository/sql"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/security"
)

type AuthAPI struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authentication(c *gin.Context) {

	var cipherPayload APICipherPayload
	if err := c.ShouldBindJSON(&cipherPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input AuthAPI
	json.Unmarshal([]byte(security.Decrypt([]byte(cipherPayload.Payload))), &input)

	auth := &models.Authentication{
		Username: input.Username,
		Password: input.Password,
	}

	err := repository.GetAuthentication(auth)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": auth,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

}
