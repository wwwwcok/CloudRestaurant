package tool

import (
	"fmt"

	"github.com/tedcy/fdfs_client"
)

//b
func UploadFile(filename string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		return ""
	}
	fileId, err := client.UploadByFilename(filename)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fileId
}
