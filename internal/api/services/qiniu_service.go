package services

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"im-services/internal/config"
)

type QiNiuService struct {
}

func (QiNiuService) UploadFile(localFile string, fileName string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.Conf.QiNiu.Bucket,
	}
	mac := qbox.NewMac(config.Conf.QiNiu.AccessKey, config.Conf.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return config.Conf.QiNiu.Domain + "/" + fileName, nil

}
