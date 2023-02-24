package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

// ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate // validator is a variable of validator.Validate
}

// validate validates product request body\
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}
func main() {
	port := os.Getenv("MY_APP_PORT")

	if port == "" {
		port = "4500"
	}
	fmt.Println("installing echo")
	e := echo.New()
	v := validator.New() //we used to do v.Struct()

	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "well , hello there !!")
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string

		for _, p := range products {
			for k := range p {

				pId, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					return err
				}

				if pId == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found") //statusNotFOund == 404, "" is optional message
		}
		return c.JSON(http.StatusOK, product)
	})

	//post method
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}

		var reqBody body
		e.Validator = &ProductValidator{validator: v} //intialising the Validator field of echo	//validator is one of the field in the echo struct, validator is a interface us interface me ek method hai Validate that takes an empty interface and return an error

		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}
		// validation
		//err2 := v.Struct(reqbody)
		err2 := c.Validate(reqBody) // ab jo validate method hogi wo apne aap check karlegi e.Validator ko

		if err2 != nil {
			return err2
		}
		CurrentProduct := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, CurrentProduct)
		return c.JSON(http.StatusOK, CurrentProduct)
	})
	e.Start(":4500")
}
