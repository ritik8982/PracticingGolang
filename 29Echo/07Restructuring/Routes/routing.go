package Routes

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"os"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("unable to load configuration ")
	}
} //Start starts the application

func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "4500"
	}

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products?:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.POST("/products", createProduct)

	e.Logger.Print(fmt.Sprint("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
