package route

import (
	"app/api"
	"app/api/admins"
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

	/**
	Get all Admin in DB
	*/
	e.GET("/admins", admins.GetAdmins)

	/**
	Get a Admin in DB
	*/
	e.GET("/admins/:id", admins.GetAdmin)

	/**
	Get all admin existing
	*/
	e.GET("/admins/existing", admins.GetAdminsExisting)

	/**
	Get all admin inviting
	*/
	e.GET("/admins/inviting", admins.GetAdminsInviting)

	/**
	Update admin info
	*/
	e.PUT("/admins/updateOne", admins.UpdateAdmin)

	/**
	Update admin name
	*/
	e.PUT("/admins/updateOne/name", admins.UpdateName)

	/**
	Update admin password
	*/
	e.PUT("/admins/updateOne/password", admins.UpdatePassword)

	/**
	Update admin email
	*/
	e.PUT("/admins/updateOne/email", admins.UpdateEmail)

	/**
	Update admin role
	*/
	e.PUT("/admins/updateOne/role", admins.UpdateRole)

	/**
	Update admin role
	*/
	e.PUT("/admins/updateOne/image", admins.UpdateImage)

	return e
}
