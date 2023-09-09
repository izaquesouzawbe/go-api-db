package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-api-db/internal/config"
	"go-api-db/internal/routes"
)

func main() {

	db := config.GetDB()
	defer db.Close()

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
