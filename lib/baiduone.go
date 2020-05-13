package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"../db/modes"
	"time"
)
var err error
type Url struct {
	TestApi string `json:"test_api"`
}
var url Url


var parmas = make(map[string]interface{})

//读取json文件 返回json字符串
func loadJson(fileName string) (json string, err error) {
	filePtr2, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件错误，请查看:", err.Error())
		return
	}
	defer filePtr2.Close()
	res, err := ioutil.ReadAll(filePtr2)
	if err != nil {
		fmt.Println("json read err", err.Error())
		return
	}
	json = string(res)
	return
}
func init(){
	urlJson, err := loadJson("./config/url.json")
	if err != nil {
		fmt.Println("Open file url.json [Err:%s]", err.Error())
		return

	}
	parmasJson, err := loadJson("./config/parmas.json")
	if err != nil {
		fmt.Println("Open file parmas.json [Err:%s]", err.Error())
		return

	}
	err = json.Unmarshal([]byte(urlJson), &url)
	if err != nil {
		fmt.Println("urlJson  failed [Err:%s]", err.Error())
		return
	}
	err = json.Unmarshal([]byte(parmasJson), &parmas)
	if err != nil {
		fmt.Println("parmasJson failed [Err:%s]", err.Error())
		return
	}
	return
}

type Header struct {
	HEADER Header_F
	BODY   BodyF
}
type Header_F struct {
	DESC     string   `json:"desc"`
	FAILURES []string `json:"failures"`
	OPRS     int      `json:"oprs"`
	SUCC     int      `json:"succ"`
	OPRTIME  int      `json:"oprtime"`
	QUOTA    int      `json:"quota"`
	RQUOTA   int      `json:"rquota"`
	STATUS   int      `json:"status"`
}
type BodyF struct {
	DATA []Body_F_C
}

type Body_F_C struct {
	ID     int `json:"id"`
	KPIS   []string
	NAME   []string
	DATE   string `json:"date"`
	KPIS_S *FieldsData

}

type FieldsData struct {
	Impression 	int64			//展示次数
	Click 		int64			//点击数量
	Cost 		int64			//所花费用
	Ctr 		int64			//点击率
	Cpc 		int64			//平均点击价格
	Cpm 		int64			//单元出价
}

func (this *FieldsData)Click_Time() string {
	var baidu modes.SicsData
	testApi := url.TestApi
	time  := fmt.Sprintf("%d-%02d-%02d",time.Now().Year(),time.Now().Month() ,time.Now().Day())
	startDate := time
	endDate := time
	parmas["body"].(map[string]interface{})["realTimeRequestType"].(map[string]interface{})["startDate"] = startDate
	parmas["body"].(map[string]interface{})["realTimeRequestType"].(map[string]interface{})["endDate"] = endDate
	//再把map 转为json
	jsonStr, err := json.Marshal(parmas)
	if err != nil {
		fmt.Println("json write err", err.Error())
		return ""
	}
	//发送http  post 请求
	req, err := http.NewRequest("POST", testApi, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http请求异常", err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	param := new(Header)
	err = json.Unmarshal(body, &param)
	if err != nil {
		return ""
	}

	for _, v := range param.BODY.DATA {

			baidu.Name = strings.Join(v.NAME," ")
			s := v.KPIS
			baidu.Impression, _ = strconv.ParseInt(s[0],10,64)
			baidu.Click, _ = strconv.ParseInt(s[1],10,64)
			//总费用
			cost, _ := strconv.ParseFloat(s[2],64)
			cost = cost * 100
			baidu.Cost = int64(cost)
			//点击率
			crts, _ := strconv.ParseFloat(s[3], 64)
			crts = crts * 100
			baidu.Ctr = int64(crts)
			//评价点击价格
			cpc, _ := strconv.ParseFloat(s[4], 64)
		    cpc = cpc * 100
			baidu.Cpc = int64(cpc)
			baidu.PId = 1001001
			//baidu.Cpm, _ = strconv.ParseInt(s[5],10,64)

	}
	//time := start_time + " 00:00:00"
	//baidu.At = StringToTimeEx(time)
	baidu.StartAt = startDate
	baidu.Start = 3
	click :=strconv.FormatInt(baidu.Click,10)
	//fmt.Println("百度clcik",click)
	//baidu.Save()
	return click
}
func (this *FieldsData)Cost_Time() string {
	var baidu modes.SicsData
	testApi := url.TestApi
	time  := fmt.Sprintf("%d-%02d-%02d",time.Now().Year(),time.Now().Month() ,time.Now().Day())
	startDate := time
	endDate := time
	parmas["body"].(map[string]interface{})["realTimeRequestType"].(map[string]interface{})["startDate"] = startDate
	parmas["body"].(map[string]interface{})["realTimeRequestType"].(map[string]interface{})["endDate"] = endDate
	//再把map 转为json
	jsonStr, err := json.Marshal(parmas)
	if err != nil {
		fmt.Println("json write err", err.Error())
		return ""
	}
	//发送http  post 请求
	req, err := http.NewRequest("POST", testApi, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http请求异常", err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	param := new(Header)
	err = json.Unmarshal(body, &param)
	if err != nil {
		return ""
	}

	for _, v := range param.BODY.DATA {

		baidu.Name = strings.Join(v.NAME," ")
		s := v.KPIS
		baidu.Impression, _ = strconv.ParseInt(s[0],10,64)
		baidu.Click, _ = strconv.ParseInt(s[1],10,64)
		//总费用
		cost, _ := strconv.ParseFloat(s[2],64)
		cost = cost * 100
		baidu.Cost = int64(cost)
		//点击率
		crts, _ := strconv.ParseFloat(s[3], 64)
		crts = crts * 100
		baidu.Ctr = int64(crts)
		//评价点击价格
		cpc, _ := strconv.ParseFloat(s[4], 64)
		cpc = cpc * 100
		baidu.Cpc = int64(cpc)
		baidu.PId = 1001001
		//baidu.Cpm, _ = strconv.ParseInt(s[5],10,64)

	}
	//time := start_time + " 00:00:00"
	baidu.At = StringToTimeEx(time)
	baidu.StartAt = startDate
	baidu.Start = 3
	fmt.Println("百度数据",baidu)
	costmoery := baidu.Cost
	return strconv.FormatInt(costmoery,10)
}

//{
//"header": {
//"username": "伟伟道11",
//"password": "Kk1K4U",
//"token": "dd170692a939181ae1adb27ecd3b4964"
//},
//"body": {
//"realTimeRequestType": {
//"unitOfTime": 5,
//"startDate": "2020-3-31",
//"reportType": 2,
//"performanceData": [
//"impression",
//"click",
//"cost",
//"ctr",
//"cpc",
//"cpm"
//],
//"endDate": "2020-3-31",
//"levelOfDetails": 2
//},
//"pid" : 1001001
//}
//}


