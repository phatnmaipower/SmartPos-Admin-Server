package update

import (
	"app/cloudinary"
	"app/db"
	"app/model"
	"app/reponse"
	"fmt"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
)

// UpdateAdmin Only update: password, name, email and role. NO IMG
func Admin(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			StatusText: "failed", Status: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
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
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func Email(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			StatusText: "failed", Status: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
		})
	}

	if requestData.Email == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			StatusText: "BadRequest",
			Status:     http.StatusBadRequest,
			Data:       "Invalid input",
		})
	}

	admin.Email = requestData.Email
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func CurrentName(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			StatusText: "failed", Status: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
		})
	}

	if requestData.Name == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			StatusText: "BadRequest",
			Status:     http.StatusBadRequest,
			Data:       "Invalid input",
		})
	}

	admin.Name = requestData.Name
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func Name(c echo.Context) error {
	fmt.Print("/admins/:id/name")

	id := c.Param("id")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")

	fmt.Print(id, firstName, lastName)

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", id)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
		})
	}

	admin.Name = firstName + " " + lastName

	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func UpdatePassword(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			StatusText: "failed", Status: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
		})
	}

	if requestData.Password == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			StatusText: "BadRequest",
			Status:     http.StatusBadRequest,
			Data:       "Invalid input",
		})
	}
	admin.Password = requestData.Password
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func UpdateRole(c echo.Context) error {

	requestData := new(model.Admin)

	err := c.Bind(requestData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, reponse.UpdateResponse{
			StatusText: "failed", Status: http.StatusInternalServerError, Data: err.Error(),
		})
	}

	dbManager := db.GetDbManager()
	admin := new(model.Admin)
	dbManager.Find(&admin, "ID = ?", requestData.ID)

	if admin == nil {
		return c.JSON(http.StatusNotFound, reponse.UpdateResponse{
			StatusText: "failed",
			Status:     http.StatusNotFound,
			Data:       "Admin does not exist, cannot update",
		})
	}

	if requestData.Role == "" {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			StatusText: "BadRequest",
			Status:     http.StatusBadRequest,
			Data:       "Invalid input",
		})
	}

	admin.Role = requestData.Role
	dbManager.Save(&admin)

	return c.JSON(http.StatusOK, reponse.UpdateResponse{
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}

func UpdateImage(c echo.Context) error {
	// Read id
	id := c.FormValue("id")

	dbManager := db.GetDbManager()
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
			StatusText: "InternalServerError",
			Status:     http.StatusInternalServerError,
			Data:       err,
		})
	}

	image, err := cloudinary.UploadImage(id, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, reponse.UpdateResponse{
			StatusText: "BadRequest",
			Status:     http.StatusBadRequest,
			Data:       err,
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
		StatusText: "success",
		Status:     http.StatusOK,
		Data:       admin,
	})
}
