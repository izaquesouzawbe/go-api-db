package routes

import (
	"database/sql"
	"go-api-db/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteExecuteSQL(router *gin.Engine, db *sql.DB) {

	router.POST("/execute-sql", authorize(config.ConfigVar.Server.Authorization), func(c *gin.Context) {

		var requestBody config.RequestBodySQL

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec(requestBody.SQL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})

	})

}
