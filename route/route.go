package route

import (
	"app/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))
	e.Use(middleware.Recover())
	e.GET("/", api.Home)

	AdminsRouteInit(e)

	///**
	//Update admin info
	//*/
	//e.PUT("/admins/:id", admins.UpdateAdmin)
	//

	///**
	//Update admin name
	//*/
	//e.PUT("/admins/updateOne/name", admins.UpdateName)
	//
	///**
	//Update admin password
	//*/
	//e.PUT("/admins/:id/password", admins.UpdatePassword)
	//
	///**
	//Update admin email
	//*/
	//e.PUT("/admins/:id/email", admins.UpdateEmail)
	//
	///**
	//Update admin role
	//*/
	//e.PUT("/admins/:id/role", admins.UpdateRole)
	//
	///**
	//Update admin role
	//*/
	//e.PUT("/admins/:id/image", admins.UpdateImage)

	return e
}
