package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1beta "github.com/mrzack99s/mrz-sso-provider/pkgs/apis/v1beta"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
)

var apiRootPath = "/v1beta/mrz-sso/apis"
var mgmtRootPath = "/v1beta/mrz-sso/mgmt"

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORS)

	var authorized_api_account = make(gin.Accounts)
	for _, user := range system.SystemConfigVar.MRZ_SSO.Security.HttpAuthUsers {
		authorized_api_account[user.Username] = user.Password
	}

	var authorized_mgmt_account = make(gin.Accounts)
	for _, user := range system.SystemConfigVar.MRZ_SSO.Security.MGMTAuthUsers {
		authorized_mgmt_account[user.Username] = user.Password
	}

	authorized_api := r.Group(apiRootPath, gin.BasicAuth(authorized_api_account))
	authorized_api.POST("/authentication", v1beta.Authentication)

	authorized_mgmt := r.Group(mgmtRootPath, gin.BasicAuth(authorized_mgmt_account))
	authorized_mgmt.POST("/add/userinfo", v1beta.AddUserInformation)
	authorized_mgmt.POST("/add/authentication", v1beta.AddAuthentication)
	// authorized_mgmt.POST("/delete/userinfo", v1beta.Authentication)
	// authorized_mgmt.POST("/delete/authentication", v1beta.Authentication)

	return r
}

func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
