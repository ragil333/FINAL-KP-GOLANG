package routes

import (
	"kp-elibrary-golang/controllers/RoleController"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Initialize() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	g := e.Group("api/")

	AuthRoutes(g)
	g.Add("GET", "/", RoleController.Index)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/auth/login" || c.Path() == "/api/auth/register" {
				return true
			}
			return false
		},
	}))

	RoleRoutes(g)
	CategoryRoutes(g)
	BookRoutes(g)
	e.Logger.Fatal(e.Start(GetPort()))
}
func GetPort() string {
	port := os.Getenv("PORT")
	return ":" + port
}
