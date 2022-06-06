package cloudinary

import (
	cloudinary2 "app/config/cloudinary"
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"time"
)

var cld *cloudinary.Cloudinary
var err error

func Init() bool {
	cld, err = cloudinary.NewFromParams(cloudinary2.EnvCloudName(), cloudinary2.EnvCloudAPIKey(), cloudinary2.EnvCloudAPISecret())

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
			Folder:       cloudinary2.EnvCloudUploadFolder(),
			ResourceType: "image",
		})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
