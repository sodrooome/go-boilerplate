package http

import (
	"backend-project/infrastructure"
	"backend-project/interfaces"
	"backend-project/usecase"
	"github.com/gin-gonic/gin"
	// "golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(infrastructure.SetHeader())

	// passing our cached files
	var search = interfaces.GetCache(interfaces.CacheSearch)

	// https://github.com/gin-gonic/gin/issues/1301
	routes := router.Group("/api/v1/product")
	routes.GET("/cache", search)
	routes.GET("/offset", interfaces.OffsetProduct)
	routes.GET("/search", interfaces.SearchProduct)
	routes.GET("/limit", interfaces.LimitProduct)
	routes.GET("/all-products", interfaces.GetAllProducts)
	routes.GET("/detailed/:product_id", interfaces.GetDetailedProduct)
	routes.POST("/create", interfaces.CreateProduct)
	routes.PUT("/update/:product_id", interfaces.UpdateProduct)
	routes.DELETE("/delete/:product_id", interfaces.DeleteProduct)

	log.Fatal(http.ListenAndServe(":8001", router))

	return router
}

func RouterCheck() *gin.Engine {
	router := gin.New()
	router.Use(infrastructure.SetHeader())
	router.Use(gin.Recovery())

	// router for health check
	router.GET("/check", interfaces.Check)

	log.Fatal(http.ListenAndServe(":8002", router))

	return router
}

func RouterValidate() *gin.Engine {
	router := gin.Default()
	router.Use(infrastructure.SetHeader())
	router.POST("/validate-products", usecase.ValidateProduct)

	log.Fatal(http.ListenAndServe(":8003", router))

	return router
}