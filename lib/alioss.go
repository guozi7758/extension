package lib

import (
	"bytes"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"os"
)

type AliOSS struct {
	Endpoint        string      `json:"endpoint"`
	AccessKeyId     string      `json:"keyid"`
	AccessKeySecret string      `json:"keysecret"`
	Bucket          string      `json:"bucket"`
	ClientHead      *oss.Client // Oss链接具柄
}

/*
 * 描述：读取OSS链接参数
 *
 *  strName : 文件的路径与名称
 *
 *****************************************************************************************/
func (this *AliOSS) ReadConfigFile(strName string) {

	jsonFile, err := os.Open(strName)
	if err != nil {
		panic("打开文件错误，请查看:" + strName)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("读取文件错误:" + strName)
	}

	json.Unmarshal(jsonData, &this)
}

/*
 * 描述：OSS链接
 *
 *****************************************************************************************/
func (this *AliOSS) Link() error {
	client, err := oss.New(this.Endpoint, this.AccessKeyId, this.AccessKeySecret)
	if err != nil {
		return err
	}
	this.ClientHead = client
	this.ClientHead.Bucket(this.Bucket)
	return nil
}

/*
 * 描述：上传本地文件
 *
 *  strOssName  : 到OSS上存放的名字
 *  strLocalFile: 本地文件路径与名字
 *
 *****************************************************************************************/
func (this *AliOSS) Upload(strOssName, strLocalFile string) error {
	bucket, err := this.ClientHead.Bucket(this.Bucket)
	if err != nil {
		return err
	}
	return bucket.UploadFile(strOssName, strLocalFile, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
}

/*
 * 描述：删除文件
 *
 *  silFiles: 文件名字.
 *
 *	注：如果删除失败，就添加到/log/alioss.log下，人工核对并删除操作。
 *
 *
 *****************************************************************************************/
//func (this *AliOSS) Delete(silFile []string) {
//	bucket, err := this.ClientHead.Bucket(this.Bucket)
//	if err != nil {
//		OssError("BUCKET 连接失败：", this.Bucket)
//		return
//	}
//	_, err = bucket.DeleteObjects(silFile)
//	if nil != err {
//		for _, val := range silFile {
//			strData := fmt.Sprintf("BUCKET-- %s : %s\n", this.Bucket, val)
//			OssError("回收OSS资源", strData)
//		}
//	}
//}

/*
 * 描述：上传本地文件
 *
 *  strOssName  : 到OSS上存放的名字
 *  strLocalFile: 本地文件路径与名字
 *
 *****************************************************************************************/
func (this *AliOSS) UploadByte(strOssName string, b []byte) error {
	bucket, err := this.ClientHead.Bucket(this.Bucket)
	if err != nil {
		return err
	}
	return bucket.PutObject(strOssName, bytes.NewReader(b))
}

/*
 * 描述：查看Bucket存不存在
 *
 *  strBucketName : Bucket的名字
 *
 *****************************************************************************************/
func (this *AliOSS) IsBucket(strBucketName string) bool {
	isExist, _ := this.ClientHead.IsBucketExist(strBucketName)
	return isExist
}

/*
func main() {
    var oss AliOSS
    oss.ReadConfigFile("./oss.json")
    err := oss.Link()
    fmt.Println(oss)
    if err != nil{
        fmt.Println("链接失败")
    }else{
        fmt.Println("链接成功")
    }
    //fmt.Println(oss.Upload("oss.json","oss.json"))
    fmt.Println( oss.Delete([]string{"oss.json"}))
}
*/
