package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"../db/modes"
	//"reflect"
)

type SLZServers struct {
	Server  	string  `json:"server"`				//登陆接口
	Report  	string  `json:"report"`				//信息接口
	Apikey  	string	`json:"apikey"`				//账户的Apikey
	UserName 	string	`json:"username"`			//360账户
	PassWd 		string 	`json:"passwd"`				//加密后的密码	AES(MD5(password),apiSecret)
	PId 		int64 	`json:"pid"`
	SToken		string	`json:"-"`

}
type Tokens struct {
	Uid		string	`json:"uid"`
	AccessToken   string	`json:"accessToken"`
}
type SllBodys struct {
	DailyList []BodySlls
}
type BodySlls struct {
	Views      		string	`json:"views"` 			//当天总展现数
	Clicks   		string  `json:"clicks"`  		//当天总点击数
	TotalCost 		string  `json:"totalCost"` 		//总消费
	Date  			string 	`json:"date"`     		//日期
	Type 			string 	`json:"type"`      		//投放端类型
}

var SanLiuZeros []SLZServers

/*
 描述 ： 获取360accessToken
 ******************************************************************************/
func ( this *SLZServers )GetToken() error{
	//fmt.Println("12312312312312312312")
	client := http.Client{}
	api  := fmt.Sprintf("%s", this.Server )
	para := fmt.Sprintf("username=%s&passwd=%s&format=JSON", this.UserName, this.PassWd)
	request, err := http.NewRequest("POST", api, strings.NewReader(para)) //请求
	if err != nil {
		return err // handle error
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") //设置Content-Type

	request.Header.Set("apiKey", this.Apikey)

	response, err := client.Do(request)                //返回
	//fmt.Println(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var val Tokens
	json.Unmarshal( body, &val)
	//fmt.Println(string(body))
	if val.AccessToken != "" {
		this.SToken = val.AccessToken
	}else{
		return nil
	}
	return nil
}
/*
 描述 ： 通过360accessToken获取每天的数据
 ******************************************************************************/
func ( this *SLZServers )SllTime() (int64,int64){
		//fmt.Println("1 13123123123123 1231 2")
	err := this.GetToken()
	client := http.Client{}
	api  := fmt.Sprintf("%s", this.Report )
	time  := fmt.Sprintf("%d-%02d-%02d",time.Now().Year(),time.Now().Month() ,time.Now().Day()-1 )
	para := fmt.Sprintf("startDate=%s&endDate=%s&format=JSON&type=all", time, time)
	request, err := http.NewRequest("POST", api, strings.NewReader(para)) //请求
	if err != nil {
		return 0,0 // handle error
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") //设置Content-Type

	request.Header.Set("apiKey", this.Apikey)
	request.Header.Set("accessToken", this.SToken)

	response, err := client.Do(request)                //返回
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0,0
	}
	//fmt.Println(string(body))
	param := new(SllBodys)
	err = json.Unmarshal(body, &param)
	if err != nil {
		return 0,0
	}
	var sll modes.SicsData

	for _, v := range param.DailyList {
		sll.Name = this.UserName
		sll.Impression, _ = strconv.ParseInt(v.Views,10,64)
		sll.Click, _ = strconv.ParseInt(v.Clicks,10,64)
		cost, _ := strconv.ParseFloat(v.TotalCost,64)
		sll.Cost = int64(cost * 100)
		ctr, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sll.Click)/float64(sll.Impression) ),64)
		ctr = ctr * 100
		sll.Ctr = int64(ctr)
		sll.Cpc =sll.Cost / sll.Click
		sll.Cpc = int64(math.Ceil(float64(sll.Cpc)))
	}
	sll.PId = this.PId
	sll.StartAt = time
	time = time + " 00:00:00"
	sll.At = StringToTimeEx(time)
	sll.Start = 6
	click := sll.Click

	cout := sll.Cost
	//fmt.Println(cout)
	//fmt.Println(click)
	fmt.Println(sll)
	return  cout,click

}

func init(){
	jsonFile, err := os.Open( "./config/sll.json" )
	if err != nil {
		fmt.Println("打开文件错误，请查看" + err.Error())
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal( data, &SanLiuZeros)
	for i, _ := range SanLiuZeros  {
		if err :=SanLiuZeros[i].GetToken(); err !=nil{
			fmt.Println("获取token失败",err.Error())
		}
	}
}
