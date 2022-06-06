package admins

import (
	"../../cloudinary"
	"../../db"
	"../../model"
	"../../reponse"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"mime/multipart"
	"net/http"
)

func GetAdmins(c echo.Context) error {
	dbManager := db.DbManager()
	var admins []model.Admin
	dbManager.Find(&admins)

	return c.JSON(http.StatusOK, admins)
}

func GetAdmin(c echo.Context) error {
	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		log.Printf("Failed unmarshalling")
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			Status: "failed", Code: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.DbManager()
	var admins []model.Admin
	dbManager.Find(&admins, "id = ?", requestData.ID)

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

// UpdateAdmin Only update: password, name, email and role. NO IMG
func UpdateAdmin(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

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
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}

func UpdateEmail(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

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

	if requestData.Email == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			Status: "BadRequest",
			Code:   http.StatusBadRequest,
			Data:   "Invalid input",
		})
	}

	admin.Email = requestData.Email
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}

func UpdateName(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

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

	if requestData.Name == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			Status: "BadRequest",
			Code:   http.StatusBadRequest,
			Data:   "Invalid input",
		})
	}

	admin.Name = requestData.Name
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}

func UpdatePassword(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

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

	if requestData.Password == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			Status: "BadRequest",
			Code:   http.StatusBadRequest,
			Data:   "Invalid input",
		})
	}
	admin.Password = requestData.Password
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}

func UpdateRole(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

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

	if requestData.Role == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			Status: "BadRequest",
			Code:   http.StatusBadRequest,
			Data:   "Invalid input",
		})
	}

	admin.Role = requestData.Role
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}

func UpdateImage(c echo.Context) error {
	// Read id
	id := c.FormValue("id")

	dbManager := db.DbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", id)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			Status: "InternalServerError",
			Code:   http.StatusInternalServerError,
			Data:   err,
		})
	}

	image, err := cloudinary.UploadImage(id, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			Status: "BadRequest",
			Code:   http.StatusBadRequest,
			Data:   err,
		})
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Print("src.Close throw an exception")
			return
		}
	}(src)

	admin.Img = image
	dbManager.Save(admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   admin,
	})
}
