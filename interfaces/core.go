package interfaces

import (
	"backend-project/conf"
	"backend-project/utils/service"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"strings"
	"time"

	// "log"
	"backend-project/domain"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

var CacheSearch = cache.New(5*time.Minute, 10*time.Minute)

// function to get reply from health check endpoint
func Check(c *gin.Context) {
	c.String(http.StatusOK, "health check it's ok")
}

func CreateProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var product domain.Products
	c.ShouldBindJSON(&product)
	if err := db.Create(&product).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	CacheSearch.Flush()
	c.JSON(200, product)
}

func GetDetailedProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var product domain.Products
	ProductId := c.Param("product_id")
	// SELECT * FROM products WHERE id = ?
	if err := db.Where("id = ?", ProductId).First(&product).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var products []domain.Products
	if err := db.Find(&products).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, products)
}

func LimitProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var products []domain.Products
	limit := c.DefaultQuery("Limit", "25")
	table := "products"
	query := db.Select(table + ".*")
	query = query.Limit(service.Limit(limit))

	if err := query.Find(&products).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, products)
}

func OffsetProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var products []domain.Products
	offset := c.DefaultQuery("Offset", "0")
	table := "products"
	query := db.Select(table + ".*")
	query = query.Offset(service.Offset(offset))

	if err := query.Find(&products).Error; err != nil {
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, products)
}

func SearchProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var products []domain.Products
	search := c.DefaultQuery("Search", "")
	table := "products"
	query := db.Select(table + ".*")
	query = query.Scopes(service.Search(search))

	if err := query.Find(&products).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, products)
}

func UpdateProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var product domain.Products
	ProductId := c.Params.ByName("product_id")
	if err := db.Where("id= ?", ProductId).First(&product).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.ShouldBindJSON(&product)
	db.Save(&product)
	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context) {
	db, err := conf.ValidateDB()
	if err != nil {
		panic("Can't connect the database")
	}

	var product domain.Products
	ProductId := c.Params.ByName("product_id")
	if err := db.Where("id= ?", ProductId).Delete(&product).Error; err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, err)
		fmt.Println("Getting an error")
		return
	}
	c.JSON(200, product)
}

func GetCache(cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		ignore := strings.ToLower(c.Query("ignoreCache")) == "true"
		getCache, cacheExists := cache.Get(c.Request.RequestURI)
		if !ignore && cacheExists {
			c.Data(200, "application/json", getCache.([]byte))
		} else {
			type bodyCache domain.Cache
			cacheWrite := &bodyCache{Cache: cache, ResponseWriter: c.Writer, URlPath: c.Request.RequestURI}
			c.Writer = cacheWrite
			c.Next()
		}
	}
}