package lib

import (
	"encoding/json"
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/tools/godoc/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	_ "strings"
)

type FileData struct {
	Num      		int64	`json:"num"` 			//序列号
	CustomerName    string  `json:"customer_name"`  //客户姓名
	CustomerPhone 	string  `json:"customer_phone"` //客户电话
	ChannelId  		string 	`json:"channel_id"`     //客户来源
	AreaCode 		int 	`json:"area_code"`      //识别手机号码城市code
	ProThe			int 	`json:"pro_the"`  		//推广端
	ProWay			int 	`json:"pro_way"`  		//推广方式
}


func (this *FileData)ParsingFile(c *gin.Context)([] FileData,error) {
	//创建一个集合
	list := make([]FileData,0)
	// 首先读excel
	xlsx, err := excelize.OpenFile(c.PostForm("file_url"))
	if err != nil {
		return nil, nil
	}
	rows ,_:= xlsx.GetRows("Sheet1")
	for i, row := range rows {
		// 去掉第一行，第一行是表头
		var val FileData
		if i == 0 {
			continue
		}
		for j, colCell := range row {

			//fmt.Println(colCell)
			// 去掉前后空格
			colCell := strings.Trim(colCell, " ")
			val.Num = int64(i)
			val.ProThe, _ = strconv.Atoi(c.PostForm("pro_the"))
			val.ProWay, _ = strconv.Atoi(c.PostForm("pro_way"))
			// 排除第一列为Null
			if j == 0 && colCell == "Null" {
				continue
			}
			if j == 0 {
				val.CustomerName = colCell
			}
			if j == 1 {
				if err := IsPhone(colCell); err != true {
					fmt.Println("PHONE ERR :", colCell)
					break
				}
				val.CustomerPhone = colCell
				var citycode CityCode
				ad, err := citycode.ToAreaCode(colCell)
				if err == nil {
					val.AreaCode = ad
				} else {
					fmt.Printf("%s ERR: %s", colCell, err.Error())
					continue
				}
			}
			if j == 2 {
				val.ChannelId = colCell
				list = append(list, val)
			}
		}
	}
	return list,nil
}

type CityCode struct {
	Province    	string    	`json:"province"`
	City 			string  	`json:"city"`
	PostCode 		string 		`json:"postcode"`
	AreaCode  		string 		`json:"areacode"`
	MobileType  	string 		`json:"mobiletype"`
}
/*
 * 描述: 识别手机号地址areacode
 *
 ***************************************************************************/
func (this *CityCode)ToAreaCode(phone string)(int,error)  {
	//?number=15966600367

	resp, err := http.Get(fmt.Sprintf("http://api.online-service.vip/phone?number=%s",phone))
	if err != nil {
		return 0,err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("body:",string(body))
	if err != nil {
		return 0,err
	}
	json.Unmarshal( body, this)
	fmt.Println("jsonbode:",this)
	code,_ :=strconv.Atoi(this.AreaCode)
	return code,nil
}
