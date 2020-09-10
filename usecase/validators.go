package usecase

import (
	"backend-project/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	v "gopkg.in/go-playground/validator.v8"
	"reflect"
)

func ProductStructValidator(validator *v.Validate, structLevel *v.StructLevel) {
	type products domain.Products

	product := structLevel.CurrentStruct.Interface().(products)
	if len(product.Name) == 0 && len(product.Description) == 0 {
		structLevel.ReportError(
			reflect.ValueOf(product.Name), "Name", "name", "productname",
		)
		structLevel.ReportError(
			reflect.ValueOf(product.Description), "Description", "description", "productdescription",
		)
	}

	if validator, ok := binding.Validator.Engine().(*v.Validate); ok {
		validator.RegisterStructValidation(ProductStructValidator, products{})
	}
}

func ValidateProduct(c *gin.Context) {
	var products domain.Products

	if err := c.ShouldBindJSON(products); err != nil {
		fmt.Println("Validate products it's okay")
	} else {
		fmt.Println("Getting and error for validate product")
	}
}
