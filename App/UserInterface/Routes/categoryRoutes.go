package routes

import (
	interactors "saleinvoice/App/Application/Interactors"
	mysqlrepository "saleinvoice/App/Infrastructure/Repositories/MysqlRepository"
	"saleinvoice/App/Infrastructure/database"
	controllers "saleinvoice/App/UserInterface/Controllers"

	"github.com/gin-gonic/gin"
)

var RegisterCategoryRouter = func(router *gin.Engine) {
	conteroller := controllers.NewCategoryController(*interactors.NewCatergoryInteractor(mysqlrepository.NewCategoryRepository(database.NewDB())))
	router.GET("/api/categories", conteroller.Index)
	router.POST("/api/category/create", conteroller.Store)

	//router.HandleFunc("/book/test", controllers.Test).Methods("POST")
}
