package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func routeInsert(router *gin.Engine, db *sql.DB) {

	router.POST("/execute-insert", authorize(config.Server.Authorization), func(c *gin.Context) {

		var requestBody InsertRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(requestBody.Records) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A lista de registros não pode estar vazia"})
			return
		}

		var successfulInserts int

		for _, record := range requestBody.Records {
			columns := ""
			values := ""
			var args []interface{}

			primaryKeyColumns := requestBody.PrimaryKeys
			upsertColumns := make([]string, 0)

			query := fmt.Sprintf("INSERT INTO %s ", requestBody.Table)

			for col, val := range record {
				if len(args) > 0 {
					columns += ", "
					values += ", "
				}
				columns += col
				values += fmt.Sprintf("$%d", len(args)+1)
				args = append(args, val)

				if !contains(primaryKeyColumns, col) {
					upsertColumns = append(upsertColumns, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
				}
			}

			query += "(" + columns + ") VALUES (" + values + ")"

			if primaryKeyColumns != nil && len(primaryKeyColumns) > 0 {
				query += "ON CONFLICT (" + strings.Join(primaryKeyColumns, ", ") +
					") DO UPDATE SET " + strings.Join(upsertColumns, ", ")
			}

			_, err := db.Exec(query, args...)
			if err != nil {
				log.Printf("Erro na inserção: %s", err.Error())
			} else {
				successfulInserts++
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d registros inseridos/atualizados com sucesso", successfulInserts)})
	})
}
