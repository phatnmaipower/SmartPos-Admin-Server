package get

import (
	"app/db"
	"app/model"
	"github.com/labstack/echo"
	"net/http"
)

func Admins(c echo.Context) error {
	dbManager := db.GetDbManager()
	var admins []model.Admin
	dbManager.Find(&admins)

	return c.JSON(http.StatusOK, admins)
}

func Admin(c echo.Context) error {
	id := c.Param("id")

	dbManager := db.GetDbManager()
	var admins []model.Admin
	dbManager.Find(&admins, "id = ?", id)

	return c.JSON(http.StatusOK, admins)
}

func AdminsExisting(c echo.Context) error {
	dbManager := db.GetDbManager()
	var adminsExisting []model.Admin
	dbManager.Find(&adminsExisting, "type = ?", "existing")

	return c.JSON(http.StatusOK, adminsExisting)
}

func AdminsInviting(c echo.Context) error {
	dbManager := db.GetDbManager()
	var adminsInviting []model.Admin
	dbManager.Find(&adminsInviting, "type = ?", "inviting")

	return c.JSON(http.StatusOK, adminsInviting)
}
