package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	//"math"
	//"strconv"
	//"strings"
	"net/http"
	"os"
	"../db/modes"
	"time"

)

type SmServer struct {
	GetReport  	string  `json:"getReport"`			//获取taskid
	GetTaskState  	string  `json:"getTaskState"`	//获取文件id
	DownLoad  	string  `json:"download"`			//下载文件
	Apikey  	string	`json:"token"`				//账户的Apikey
	UserName 	string	`json:"username"`			//360账户
	PassWd 		string 	`json:"password"`			//加密后的密码	AES(MD5(password),apiSecret)
	TaskId  int64 `json:"taskId"`   //taskid
	FileId  int64 `json:"fileId"`   //fileId

	PId  int64 `json:"pid"`   //fileId
}
type SenMaBody struct {
	Header HeaderSenMa `json:"header"`
	Body  BodySenMa `json:"body"`
}
type HeaderSenMa struct {
	UserName  string `json:"username"`   //用户账号
	PassWord  string `json:"password"`   //状态
	Token string  `json:"token"` 		 //用户token
}
type BodySenMa struct {
	StartDate  string `json:"startDate"`   //用户账号
	EndDate    string `json:"endDate"`     //状态
	ReportType int  `json:"reportType"`    //账号报告
	TaskId  int64 `json:"taskId"`      //taskid
	FileId  int64 `json:"fileId"`      //fileid
	Status  string `json:"status"`     //状态
	ProRecs float32  `json:"progress"` //0.0
	Success bool `json:"success"`      //状态
}
type ResPonBodySenMa struct {
	TaskId  int64 `json:"taskId"`      //taskid
	Status  string `json:"status"`     //状态
	ProRecs float32  `json:"progress"` //0.0
	Success bool `json:"success"`      //状态
}

var SmZero SmServer

/*
 描述 ： 创建什么会话ID taksID
 ******************************************************************************/
func ( this *SmServer )GetTaskId() int64{

	client := http.Client{}
	api  := fmt.Sprintf("%s", SmZero.GetReport )
	time  := fmt.Sprintf("%d-%02d-%02d",time.Now().Year(),time.Now().Month() ,time.Now().Day() )

	nba := new(SenMaBody)
	nba.Header.UserName = SmZero.UserName
	nba.Header.PassWord = SmZero.PassWd
	nba.Header.Token = SmZero.Apikey
	nba.Body.StartDate = time
	nba.Body.EndDate = time
	nba.Body.ReportType = 2
	//再把map 转为json
	jsonStr, err := json.Marshal(nba)

	request, err := http.NewRequest("POST", api, bytes.NewBuffer([]byte(jsonStr))) //请求
	if err != nil {
		return 0 // handle error
	}
	request.Header.Set("Content-Type", "application/json") //设置Content-Type

	response, err := client.Do(request)                //返回
	if err != nil {
		return  0
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0
	}

	err = json.Unmarshal(body, &nba)
	if err != nil {
		return 0
	}
	fmt.Println("TaskId",nba.Body.TaskId)
	//if nba.Body.TaskId >= 0 {
	//	this.TaskId = nba.Body.TaskId
	//	SmZero.GetFileId()
	//}else{
	//	return nil
	//}
	//return nil
	return nba.Body.TaskId
}
/*
 描述 ： 获取文件FileId
 ******************************************************************************/
func ( this *SmServer )GetFileId(TaskId int64) int64{

	client := http.Client{}
	api  := fmt.Sprintf("%s", SmZero.GetTaskState )
	nba := new(SenMaBody)
	nba.Header.UserName = SmZero.UserName
	nba.Header.PassWord = SmZero.PassWd
	nba.Header.Token = SmZero.Apikey
	nba.Body.TaskId = TaskId
	//再把map 转为json
	jsonStr, err := json.Marshal(nba)
	request, err := http.NewRequest("POST", api, bytes.NewBuffer([]byte(jsonStr))) //请求
	if err != nil {
		return 0 // handle error
	}
	request.Header.Set("Content-Type", "application/json") //设置Content-Type
	response, err := client.Do(request)                //返回
	if err != nil {
		fmt.Println("err:", err.Error())
		return 0
	}
	defer response.Body.Close()

	time.Sleep(time.Second *5)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &nba)
	if err != nil {
		return 0
	}

	//if nba.Body.Success == true && nba.Body.FileId >=0 {
	//	this.FileId = nba.Body.FileId
	//	SmZero.GetFileDload()
	//}else{
	//	return nil
	//}
	//return nil
	return nba.Body.FileId
}

/*
 描述 ： 获取文件数据
 ******************************************************************************/
func ( this *SmServer )GetFileDload(FileId int64) (string){
	client := http.Client{}
	api  := fmt.Sprintf("%s", SmZero.DownLoad )
	nba := new(SenMaBody)
	nba.Header.UserName = SmZero.UserName
	nba.Header.PassWord = SmZero.PassWd
	nba.Header.Token = SmZero.Apikey
	nba.Body.FileId = FileId
	//再把map 转为json
	jsonStr, err := json.Marshal(nba)
	request, err := http.NewRequest("POST", api, bytes.NewBuffer([]byte(jsonStr))) //请求
	if err != nil {
		return "" // handle error
	}
	request.Header.Set("Content-Type", "application/json") //设置Content-Type
	response, err := client.Do(request)                //返回
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	rcdata := string(body)
	str := strings.Split(rcdata,",")
	fmt.Println("******",str)

	cos,_ :=strconv.ParseFloat(str[12],64)
	cost := cos * 100

	costs := strconv.FormatFloat( cost, 'f', -1, 64)
	//click ,_ := strconv.ParseInt(str[11],10,64)
	return costs
}

/*
 描述 ： 获取文件数据
 ******************************************************************************/
func ( this *SmServer )GetFileCilck(FileId int64) (int64,int64){
	var SenMa modes.SicsData
	client := http.Client{}
	api  := fmt.Sprintf("%s", SmZero.DownLoad )
	nba := new(SenMaBody)
	nba.Header.UserName = SmZero.UserName
	nba.Header.PassWord = SmZero.PassWd
	nba.Header.Token = SmZero.Apikey
	nba.Body.FileId = FileId
	//再把map 转为json
	jsonStr, err := json.Marshal(nba)
	request, err := http.NewRequest("POST", api, bytes.NewBuffer([]byte(jsonStr))) //请求
	if err != nil {
		return 0,0 // handle error
	}
	request.Header.Set("Content-Type", "application/json") //设置Content-Type
	response, err := client.Do(request)                //返回
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println("神马数据",string(body))
	if err != nil {
		fmt.Println("err:", err.Error())
		return 0,0
	}
	str := strings.Split(string(body),",")
	if len(str)  < 12 {
		fmt.Println("数据请求失败", string(body) )
		return 0,0
	}
	if str[11] == ""{
		SenMa.Click = 0
	}else{
		SenMa.Click, _ = strconv.ParseInt(str[11],10,64)
	}
	co, _ := strconv.ParseFloat(str[12],64)
	if str[12] ==""{
		SenMa.Cost = 0
	}else{
		cos := co * 100
		SenMa.Cost = int64(cos)
	}
	return SenMa.Click,SenMa.Cost
}

func init(){
	jsonFile, err := os.Open( "./config/shenma.json" )
	if err != nil {
		fmt.Println("打开文件错误，请查看:" + err.Error())
		return
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)

	json.Unmarshal( data, &SmZero)
	SmZero.GetTaskId()
	//if err:= SmZero.GetTaskId(); err !=nil {
	//	panic("打开文件错误，请查看:" + err.Error())
	//}
}
