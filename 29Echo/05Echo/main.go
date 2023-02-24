package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	port := os.Getenv("MY_APP_PORT")

	if port == "" {
		port = "4500"
	}
	fmt.Println("installing echo")
	e := echo.New()
	v := validator.New() //it return validate, validate is struct

	// e.GET("/products/:vendor", func(c echo.Context) error {

	// 	// http://localhost:3000/products/:vendor?olderThan=iphone10

	// 	// return c.JSON(http.StatusOK,c.Param("vendor"))	// before this ?
	// 	return c.JSON(http.StatusOK, c.QueryParam("olderThaner")) // after this ?
	// })

	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}} //slice of map, assume you made call to db and get this data

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "well , hello there !!")
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	// instead of passing the name of the function i am defining the function here
	// id will be supply by the client and can be dynamic that's why :

	//get method
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string

		for _, p := range products {
			for k := range p {

				pId, err := strconv.Atoi(c.Param("id")) //param jo hai wo url  hai, atoi (str -> int),

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
		return c.JSON(http.StatusOK, product) //StatusOk == 200, product ko return kar rahe hai
	})

	//post method
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name            string `json:"product_name" validate:"required,min=4"`
			Vendor          string `json:"vendor" validate:"min=5,max=10"`
			Email           string `json:"email" validate:"required_with=Vendor,email"` // if vendor is provided then we have to provide this as well
			Website         string `json:"website" validate:"url"`
			DefaultDeviceIP string `json:"default_device_ip" validate:"ip"` //this field should be ip type
		}
		var reqBody body //jis type ka data aane wala hai us type ek data type banaliye and uska ek variable banadiye aur pass kardiye bind me
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}
		// validation
		err2 := v.Struct(reqBody) // it will check validation of every field and if all are correct then only it will not return eror
		if err2 != nil {
			return err2
		}

		CurrentProduct := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, CurrentProduct)
		return c.JSON(http.StatusOK, CurrentProduct) //we have to return something so that it can show the data is added successfullly
	})

	//put method
	e.PUT("/products/:id", func(c echo.Context) error {

		var product map[int]string
		pId, err5 := strconv.Atoi(c.Param("id"))

		if err5 != nil {
			return err5
		}
		for _, p := range products {
			for k := range p {
				if pId == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		type body struct {
			Name            string `json:"product_name" validate:"required,min=4"`
			Vendor          string `json:"vendor" validate:"min=5,max=10"`
			Email           string `json:"email" validate:"required_with=Vendor,email"` // if vendor is provided then we have to provide this as well
			Website         string `json:"website" validate:"url"`
			DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
		}
		var reqBody body
		err3 := c.Bind(&reqBody)
		if err3 != nil {
			return err3
		}
		// validation
		err2 := v.Struct(reqBody)
		if err2 != nil {
			return err2
		}

		product[pId] = reqBody.Name
		return c.JSON(http.StatusOK, product)

	})

	//delete method
	e.DELETE("/products/:id", func(c echo.Context) error {

		pid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		var idx int

		for index, p := range products {
			for k := range p {
				if pid == k {
					idx = index
				}
			}
		}
		if idx == 0 {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		rval := products[idx]
		slyce := append(products[0:idx], products[idx+1:]...)
		products = slyce

		return c.JSON(http.StatusOK, rval)
	})
	e.Start(":4500")
}
