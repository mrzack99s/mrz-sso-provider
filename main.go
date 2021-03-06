package main

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	apis "github.com/mrzack99s/mrz-sso-provider/pkgs/apis/routes"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/sql_db"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/tools"
)

func main() {

	mode := gin.DebugMode

	gin.SetMode(mode)

	system.ParseSystemConfig("./config.yaml")

	db := sql_db.MySQL{
		Username: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Username,
		Password: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Password,
		Hostname: system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.Hostname,
		DBName:   system.SystemConfigVar.MRZ_SSO.DB.SQL.Production.DBName,
	}
	db.Initial()

	tools.SQL_Migrate()

	router := apis.SetupRouter()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Run(":8888")

	// Close DB
	objDB, _ := sql_db.SQL_DB.DB()
	objDB.Close()

}
