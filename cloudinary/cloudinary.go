package cloudinary

import (
	"../config"
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	_ "github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	_ "github.com/cloudinary/cloudinary-go/api/uploader"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	_ "log"
	"time"
)

var cld *cloudinary.Cloudinary
var err error

func Init() bool {
	cld, err = cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())

	if err != nil {
		fmt.Print("Cannot connect to cloudinary with err", err)
		return false
	}
	fmt.Print("Connect successfully\n")
	return true

}

func UploadImage(publicID string, input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !Init() {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(
		ctx,
		input,
		uploader.UploadParams{
			PublicID:     publicID,
			Folder:       config.EnvCloudUploadFolder(),
			ResourceType: "image",
		})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
