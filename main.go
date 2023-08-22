package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db := getDB()
	router := gin.Default()

	routeSelect(router, db)
	routeExecuteSQL(router, db)
	routeExecuteSQLTransaction(router, db)
	routeInsert(router, db)
	routeInsertTransaction(router, db)

	err := router.Run(":" + config.Server.Port)
	if err != nil {
		return
	}
}
