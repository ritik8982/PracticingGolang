package Routes

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
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
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
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
		Email           string `json:"email" validate:"required_with=Vendor,email"` // if vendor diya to ye bhi dena hoga
		Website         string `json:"website" validate:"url"`
		DefaultDeviceIP string `json:"default_device_ip" validate:"ip"` //check karo ki ip hai ke nhi
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

}

func createProduct(c echo.Context) error {
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
	err2 := c.Validate(reqBody) // now the validate method will check that reqBody is having all the validatoin that we have applied, if not then it will return error

	if err2 != nil {
		return err2
	}
	CurrentProduct := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, CurrentProduct)
	return c.JSON(http.StatusOK, CurrentProduct)
}

func deleteProduct(c echo.Context) error {
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
}
