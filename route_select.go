package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeSelect(router *gin.Engine, db *sql.DB) {

	router.POST("/execute-select", authorize(config.Server.Authorization), func(c *gin.Context) {

		var requestBodySQL RequestBodySQL

		if err := c.ShouldBindJSON(&requestBodySQL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rows, err := db.Query(requestBodySQL.SQL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var results []map[string]interface{}

		for rows.Next() {
			values := make([]interface{}, len(columns))
			pointers := make([]interface{}, len(columns))
			for i := range columns {
				pointers[i] = &values[i]
			}

			if err := rows.Scan(pointers...); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			record := make(map[string]interface{})
			for i, colName := range columns {
				val := values[i]
				if b, ok := val.([]byte); ok {
					record[colName] = string(b)
				} else {
					record[colName] = val
				}
			}
			results = append(results, record)
		}

		c.JSON(http.StatusOK, results)
	})

}
