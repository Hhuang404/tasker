package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// test case to create delete bucket
func TestNewBucketSample(t *testing.T) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		handleError(err)
	}
	// 创建存储空间。
	err = client.CreateBucket(imagesBucket)
	if err != nil {
		handleError(err)
	}
	// 删除 bucket
	//client.DeleteBucket(bucket)
}
