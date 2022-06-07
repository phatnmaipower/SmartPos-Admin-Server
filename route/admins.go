package route

import (
	get "app/api/admins/get"
	update "app/api/admins/update"
	"github.com/labstack/echo"
)

func AdminsRouteInit(e *echo.Echo) {
	adminRoute := e.Group("/admins")
	adminRoute.GET("", get.Admins)
	adminRoute.GET("/:id", get.Admin)
	adminRoute.GET("/existing", get.AdminsExisting)
	adminRoute.GET("/inviting", get.AdminsInviting)
	adminRoute.PUT("/:id", update.Admin)
	adminRoute.PUT("/:id/name", update.Name)
	//
	///**
	//Update admin name
	//*/
	//e.PUT("/admins/:id/name", admins.UpdateAdminName)
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

}
