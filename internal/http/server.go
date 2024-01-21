package http

import (
	"log"
	"net/http"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/core/interfaces"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/core/utils"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/handlers"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	echo           *echo.Echo
	productHandler *handlers.ProductHandler
}

func NewServer(productRepository interfaces.ProductRespositoryInterface) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI::${uri}\n, METHOD::${method},  STATUS::${status}, HEADER::${header}\n, QUERY::${query}\n, ERROR::${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           utils.ColorLoggerOutput(),
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())

	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	return &Server{
		echo:           e,
		productHandler: productHandler,
	}
}

func (s *Server) RouteInit(address string) {
	s.routeConfig()

	err := s.echo.Start(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	api := s.echo.Group("/api/v1")

	products := api.Group("/products")
	products.GET("", s.productHandler.Index)
	products.POST("", s.productHandler.Create)
	products.GET("/:id", s.productHandler.Show)
	products.DELETE("/:id", s.productHandler.Delete)
}
