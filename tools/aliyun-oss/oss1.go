package upload

import (
	"log"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	OSSEndpoint        = ""
	OSSAccessKeyID     = ""
	OSSAccessKeySecret = ""
	OSSBucket          = ""
)

var (
	Bucket *oss.Bucket
	once   = sync.Once{}
)

// InitBucket ...
func InitBucket() {
	once.Do(func() {
		client, err := oss.New(OSSEndpoint, OSSAccessKeyID, OSSAccessKeySecret)
		if err != nil {
			log.Print("oss.New error", err)
		}

		Bucket, err = client.Bucket(OSSBucket)
		if err != nil {
			log.Print("client.Bucket error", err)
		}
	})
}

// UploadOSSApk ...
func UploadOSSApk(filepath, objectKey string) (string, error) {
	err := Bucket.PutObjectFromFile(objectKey, filepath)
	if err != nil {
		log.Print("UploadFileToOSS error ", err)
		return "", err
	}

	downloadURL := "https://" + OSSBucket + "." + OSSEndpoint + "/" + objectKey
	return downloadURL, nil
}
