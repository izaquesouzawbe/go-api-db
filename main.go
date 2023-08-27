package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-api-db/config"
	"go-api-db/routes"
)

func main() {

	db := config.GetDB()
	router := gin.Default()

	routes.RouteSelect(router, db)
	routes.RouteExecuteSQL(router, db)
	routes.RouteExecuteSQLTransaction(router, db)
	routes.RouteInsert(router, db)
	routes.RouteInsertTransaction(router, db)

	err := router.Run(":" + config.ConfigVar.Server.Port)
	if err != nil {
		return
	}
}
