package upload

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// AliYun ...
var AliYun = new(aliyun)

type aliyun struct {
	err    error
	file   multipart.File
	client *oss.Client
	bucket *oss.Bucket
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传文件到aliyun对象存储
func (a *aliyun) Upload(file *multipart.FileHeader) (path string, key string, err error) {
	if a.file, a.err = file.Open(); a.err != nil { // 读取文件
		g.Log().Error("function file.Open() Failed!", g.Map{"err": a.err})
		return path, key, a.err
	}

	defer func() { // multipart.File 对象 defer 关闭
		_ = a.file.Close()
	}()

	objectType := oss.ContentType(file.Header.Get("content-type")) // 获取文件类型
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)                  // 指定访问权限为公共读，缺省为继承bucket的权限。
	objectName := GetObjectName(file.Filename)                     // 文件对象名

	if a.err = a.bucket.PutObject(objectName, a.file, a.StorageClassType("Standard"), objectType, objectAcl); a.err != nil { // 上传文件到阿里云
		g.Log().Error("function bucket.PutObject() Failed!", g.Map{"err": a.err})
		return path, key, a.err
	}

	return objectName, objectName, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据key删除在aliyun对象存储的文件
func (a *aliyun) Delete(key string) error {
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	if a.err = a.bucket.DeleteObject(key); a.err != nil {
		fmt.Println("Delete File Failed!", a.err)
		return a.err
	}

	return nil
}

// AliYunInit ...
func AliYunInit() (result *aliyun, err error) {
	var info aliyun
	if info.client, info.err = oss.New(OSSEndpoint, OSSAccessKeyID, OSSAccessKeySecret, oss.Timeout(10, 120)); info.err != nil {
		return nil, info.err
	}

	if info.bucket, info.err = info.client.Bucket(OSSBucket); info.err != nil { // 获取存储空间。
		return nil, info.err
	}

	return &info, nil
}

// GetObjectName ...
func GetObjectName(filename string) string {
	folder := time.Now().Format("20060102")
	return fmt.Sprintf("%s/%d%s", folder, time.Now().Unix(), filename) // 文件名格式 自己可以改 建议保证唯一性
}

func (a *aliyun) StorageClassType(storageClassType string) oss.Option {
	switch storageClassType { // 根据配置文件进行指定存储类型
	case "Standard": // 指定存储类型为标准存储，默认也为标准存储。
		return oss.ObjectStorageClass(oss.StorageStandard)
	case "IA": // 指定存储类型为很少访问存储
		return oss.ObjectStorageClass(oss.StorageIA)
	case "Archive": // 指定存储类型为归档存储。
		return oss.ObjectStorageClass(oss.StorageArchive)
	case "ColdArchive": // 指定存储类型为归档存储。
		return oss.ObjectStorageClass(oss.StorageColdArchive)
	default:
		return oss.ObjectStorageClass(oss.StorageStandard)
	}
}
