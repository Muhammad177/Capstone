package database

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func UploadImageCloud(fileHeader *multipart.FileHeader) (imageUrl string, err error) {

	file, _ := fileHeader.Open()

	cld := Credentials()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       "Uploads" + "/" + fileHeader.Filename,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}
