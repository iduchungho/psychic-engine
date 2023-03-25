// Package database cloudinary
package cloudinary

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"sync"
)

var cld *cloudinary.Cloudinary
var lockCld = &sync.Mutex{}

func GetConnCloudinary() *cloudinary.Cloudinary {
	if cld == nil {
		lockCld.Lock()
		defer lockCld.Unlock()
		if cld == nil {
			cldLocal, err := cloudinary.New()
			if err != nil {
				log.Fatal(err)
			}
			cld = cldLocal
			fmt.Println("Cloudinary connected !!")
			return cld
		} else {
			return cld
		}
	}
	return cld
}

func UpdateImages(cld *cloudinary.Cloudinary, file interface{}) (*uploader.UploadResult, error) {
	var ctx = context.Background()
	updateResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	return updateResult, err
}
