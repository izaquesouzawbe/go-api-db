package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeExecuteSQL(router *gin.Engine, db *sql.DB) {

	router.POST("/execute-sql", authorize(config.Server.Authorization), func(c *gin.Context) {

		var requestBody RequestBodySQL

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
