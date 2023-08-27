package routes

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-db/config"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func RouteInsertTransaction(router *gin.Engine, db *sql.DB) {

	router.POST("/execute-insert-transaction", authorize(config.ConfigVar.Server.Authorization), func(c *gin.Context) {

		var requestBody config.InsertRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(requestBody.Records) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A lista de registros não pode estar vazia"})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao iniciar a transação"})
			return
		}

		var successfulInserts int
		querys := ""

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

				if reflect.TypeOf(val).String() == "string" {
					values += "'" + val.(string) + "'"
				} else {
					values += strconv.FormatFloat(val.(float64), 'f', -1, 64)
				}

				args = append(args, val)

				if !contains(primaryKeyColumns, col) {
					upsertColumns = append(upsertColumns, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
				}
			}

			query += "(" + columns + ") VALUES (" + values + ")"
			fmt.Println(query)

			if primaryKeyColumns != nil && len(primaryKeyColumns) > 0 {
				query += " ON CONFLICT (" + strings.Join(primaryKeyColumns, ", ") +
					") DO UPDATE SET " + strings.Join(upsertColumns, ", ")
			}

			querys += query + ";"

		}

		_, err = tx.Exec(querys)
		if err != nil {
			tx.Rollback()
			log.Printf("Erro na inserção: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro na transação"})
			return
		} else {
			successfulInserts++
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao confirmar a transação"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d registros inseridos/atualizados com sucesso", successfulInserts)})
	})
}
