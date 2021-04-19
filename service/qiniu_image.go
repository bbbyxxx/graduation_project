package services

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const (
	//本地保存的文件夹名称
	upload_path string = "/files/"
)

var (
	//BUCKET是你在存储空间的名称
	ACCESS_KEY = "XAxI5JB-88Thu86GyLAWYg7UImGuDP5pM-pZ6_MM"
	SECRET_KEY = "8yiUbe-tzNFVBo1VR_6P7shV0UBAbDQgMbzknZbl"
)

func UploadImageToQiNiuYun(localFile string, key string) (string, string, error) {
	bucket := "bbbyxxx"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ret.Key, "", err
	}
	return ret.Key, ret.Hash, nil
}
