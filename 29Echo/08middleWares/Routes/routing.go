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
}

// next is the function of echo.HandlerFunc type mtlb serverMessage ek function leta hai as an argu and return echo.handlerFunc
func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside  middleware")
		return next(c)
	}
}

//type MiddlewareFunc func(HandlerFunc) HandlerFunc
// type HandlerFunc func(Context) error

//Start starts the application

func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "4500"
	}
	//common middle for all the route
	e.Use(serverMessage) // one doubt i have only passed the name of the func, not handlerFunc argument then how it is calling
	// route level middleware, ek route pe middleWare
	//e.GET("/products", getProducts, serverMessage) // first middleWare will be called then Function

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products?:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.POST("/products", createProduct)

	e.Logger.Print(fmt.Sprint("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
