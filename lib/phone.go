package lib

import (
	"encoding/json"
	"fmt"
	_ "golang.org/x/tools/godoc/util"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "strings"
)


type Citys struct {
	Province    	string    	`json:"province"`
	City 			string  	`json:"city"`
	PostCode 		string 		`json:"postcode"`
	AreaCode  		string 		`json:"areacode"`
	MobileType  	string 		`json:"mobiletype"`
}
type Juhe struct {
	Resultcode 		string `json:"resultcode"`
	Reason 			string `json:"reason"`
	Result struct {
		Province 	string `json:"province"`
		City 		string `json:"city"`
		Areacode 	string `json:"areacode"`
		Zip 		string `json:"zip"`
		Company 	string `json:"company"`
		Card 		string `json:"card"`
		} 				   `json:"result"`
	ErrorCode 		int `json:"error_code"`
}

type PhoneArea struct {
	Success string `json:"success"`
	Result struct {
		Status string `json:"status"`
		Phone string `json:"phone"`
		Area string `json:"area"`
		Postno string `json:"postno"`
		Att string `json:"att"`
		Ctype string `json:"ctype"`
		Par string `json:"par"`
		Prefix string `json:"prefix"`
		Operators string `json:"operators"`
		StyleSimcall string `json:"style_simcall"`
		StyleCitynm string `json:"style_citynm"`
	} `json:"result"`
}

func (this *PhoneArea)Area(phone string)(int64,error)  {
	resp, err := http.Get(fmt.Sprintf("http://api.k780.com:88/?app=phone.get&phone=%s&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json",phone))
	if 	err != nil {
		return 0,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0,err
	}
	json.Unmarshal( body, this)
	fmt.Println("jsonbode:",this)
	return strconv.ParseInt(this.Result.Area,10,64)
}

func (this *Juhe)Area(phone string)(int64,error)  {
	resp, err := http.Get(fmt.Sprintf("http://apis.juhe.cn/mobile/get?phone=%s&dtype=josn&key=baf59408d09965a7d45003e24cd23626",phone))
	if 	err != nil {
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
	return strconv.ParseInt(this.Result.Areacode,10,64)
}


/*
 * 描述: 识别手机号地址areacode
 *
 ***************************************************************************/
func (this *Citys)ToArea(phone string)(int64,error)  {
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
	return strconv.ParseInt(this.AreaCode,10,64)
}