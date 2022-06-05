package api

import (
	"../db"
	"../model"
	"../reponse"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func GetAdmins(c echo.Context) error {
	dbManager := db.DbManager()
	var admins []model.Admin
	dbManager.Find(&admins)

	return c.JSON(http.StatusOK, admins)
}

func GetAdminsExisting(c echo.Context) error {
	dbManager := db.DbManager()
	var adminsExisting []model.Admin
	dbManager.Find(&adminsExisting, "type = ?", "existing")

	return c.JSON(http.StatusOK, adminsExisting)
}

func GetAdminsInviting(c echo.Context) error {
	dbManager := db.DbManager()
	var adminsInviting []model.Admin
	dbManager.Find(&adminsInviting, "type = ?", "inviting")

	return c.JSON(http.StatusOK, adminsInviting)
}

func UpdateAdmin(c echo.Context) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to parse body")
		}
	}(c.Request().Body)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body:%s", err)
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			Status: "failed", Code: http.StatusInternalServerError, Data: "Failed reading the request body",
		})
	}

	requestData := new(model.Admin)
	err = json.Unmarshal(b, &requestData)

	if err != nil {
		log.Printf("Failed unmarshalling")
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			Status: "failed", Code: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.DbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			Status: "failed",
			Code:   http.StatusNotFound,
			Data:   "Admin does not exist, cannot update",
		})
	}

	if requestData.Password != "" {
		admin.Password = requestData.Password
	}
	if requestData.Name != "" {
		admin.Name = requestData.Name
	}
	if requestData.Email != "" {
		admin.Email = requestData.Email
	}
	if requestData.Role != "" {
		admin.Role = requestData.Role
	}

	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "failed",
		Code:   http.StatusOK,
		Data:   admin,
	})
}
