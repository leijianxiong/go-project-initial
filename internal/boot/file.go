package boot

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-project-initial/configs"
	"mime/multipart"
)

// LocalFileUploader is an Uploader of local file engine.
type LocalFileUploader struct {
	BasePath   string
	BasePrefix string
}

// GetLocalFileUploader return the default Uploader.
func GetLocalFileUploader() file.Uploader {
	return &LocalFileUploader{
		configs.ProjectDir() + "/uploads",
		configs.Conf.App.StorePrefix,
	}
}

// Upload implements the Uploader.Upload.
func (local *LocalFileUploader) Upload(form *multipart.Form) error {
	return file.Upload(func(fileObj *multipart.FileHeader, filename string) (string, error) {
		if err := file.SaveMultipartFile(fileObj, (*local).BasePath+"/"+filename); err != nil {
			return "", err
		}
		return local.BasePrefix + "/" + filename, nil
	}, form)
}

type AliyunOssFileUploader struct {
	client *oss.Client
	conf   *configs.AliyunOss
}

func GetAliyunOssFileUploader() file.Uploader {
	up := new(AliyunOssFileUploader)
	up.conf = configs.Conf.AliyunOss
	client, err := oss.New(up.conf.Endpoint, up.conf.AccessKeyId, up.conf.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	up.client = client
	return up
}

func (aliup *AliyunOssFileUploader) Upload(form *multipart.Form) error {
	return file.Upload(func(fileObj *multipart.FileHeader, filename string) (path string, err error) {
		bucket, err := aliup.client.Bucket(aliup.conf.BucketName)
		if err != nil {
			return
		}
		f, err := fileObj.Open()
		if err != nil {
			return
		}
		err = bucket.PutObject(filename, f)
		if err != nil {
			return
		}

		path = fmt.Sprintf("https://%s.%s/%s", aliup.conf.BucketName, aliup.conf.Endpoint, filename)
		return
	}, form)
}

func init() {
	file.AddUploader("local-prefix", GetLocalFileUploader)
	file.AddUploader("aliyun-oss", GetAliyunOssFileUploader)
}
