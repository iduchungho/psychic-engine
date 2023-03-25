package cloud

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"os"
	"sync"
)

var cld *cloudinary.Cloudinary
var lockCld = &sync.Mutex{}

func GetConnCloudinary() *cloudinary.Cloudinary {
	if cld == nil {
		lockCld.Lock()
		defer lockCld.Unlock()
		if cld == nil {
			cldLocal, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
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
	updateResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       "smart-home",
	})
	return updateResult, err
}
