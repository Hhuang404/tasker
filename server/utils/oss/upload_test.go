package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"testing"
)

func TestUpload(t *testing.T) {

	client, err := oss.New(endpoint, accessID, accessKey)
	// 获取存储空间。
	bucket, err := client.Bucket(imagesBucket)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile("animals_pandas_bears_cute_1920x1080.jpg", "C:\\Users\\黄浩\\Pictures\\Saved Pictures\\cats_blue_eyes_animals_pets_4288x2848.jpg")
	if err != nil {
		handleError(err)
	}
}
