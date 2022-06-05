package route

import (
	"../api"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
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

	/**
	Get all Admin in DB
	*/
	e.GET("/admins", api.GetAdmins)

	/**
	Get all admin existing
	*/
	e.GET("/admins/existing", api.GetAdminsExisting)

	/**
	Get all admin inviting
	*/
	e.GET("/admins/inviting", api.GetAdminsInviting)

	/**
	Update admin info
	*/
	e.GET("/admin/update", api.UpdateAdmin)
	
	return e
}
