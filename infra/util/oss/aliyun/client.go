package aliyun

//import (
//	"crgo/infra/log"
//	"io"
//	"path"
//	"strings"
//	"sync"
//
//	"github.com/aliyun/aliyun-oss-go-sdk/oss"
//	"github.com/spf13/viper"
//
//	"crgo/infra/bizerror"
//)
//
//var instance *oss.Client
//var once sync.Once
//
//// PICTRUE 图片类型
//const PICTRUE = 1
//
//// VIDEO 视频类型
//const VIDEO = 2
//
//// PROTOCOL 协议
//const PROTOCOL string = "https"
//
//// ProviderID 供应商ID，阿里云分配 1 号
//const ProviderID = 1
//
//// New ... 创建oss client对
//func New() *oss.Client {
//	once.Do(func() {
//		// 创建 oss client 实例
//		conf := viper.Sub("oss.aliyun")
//		endpoint := conf.GetString("endpoint")
//		accessKeyID := conf.GetString("access_key_id")
//		accessKeySecret := conf.GetString("access_key_secret")
//		instance, _ = oss.New(endpoint, accessKeyID, accessKeySecret)
//	})
//	return instance
//}
//
////// Bucket : 获取bucket存储空间
//func Bucket() *oss.Bucket {
//	cli := New()
//	if cli != nil {
//		bucket, _ := cli.Bucket(viper.Sub("oss.aliyun").GetString("bucket"))
//		return bucket
//	}
//	return nil
//}
//
//// GetPictureURL 根据图片 URI 返回 URL
//func GetPictureURL(uri string, stypeName string) string {
//	conf := viper.Sub("cdn")
//	return PROTOCOL + "://" + conf.GetString("domain") + uri + "_" + stypeName
//}
//
//// GetOriginURL 返回原始 URL
//func GetOriginURL(uri string) string {
//	if uri == "" {
//		return uri
//	}
//	conf := viper.Sub("cdn")
//	return PROTOCOL + "://" + conf.GetString("domain") + uri
//}
//
//// ParseURI 返回原始 URI
//func ParseURI(url string) string {
//	if url == "" {
//		return url
//	}
//	prefix := PROTOCOL + "://" + viper.Sub("cdn").GetString("domain")
//	return strings.TrimPrefix(url, prefix)
//}
//
//// ParseDigest 从 URL 返回文件名
//func ParseDigest(url string) string {
//	filename := path.Base(url)
//	digest := strings.TrimSuffix(filename, path.Ext(filename))
//	return digest
//}
//
//// 获取 bucket
//func GetBucket(bucketName string) *oss.Bucket {
//	var b string
//	conf := viper.Sub("oss.aliyun")
//	if bucketName == "" {
//		b = conf.GetString("bucket")
//	} else {
//		b = conf.GetString(bucketName)
//	}
//	log.Debug(b)
//	// 上传
//	ossCli := New()
//	bucket, err := ossCli.Bucket(b)
//	if err != nil {
//		panic(bizerror.NewError(1003, "获取 bucket 失败", err.Error()))
//	}
//	return bucket
//}
//
//// 请求 oss 上传
//func Upload(bucketName, path string, r io.Reader) error {
//	// 上传
//	bucket := GetBucket(bucketName)
//	err := bucket.PutObject(path, r)
//	if err != nil {
//		return err
//	}
//	return nil
//}
