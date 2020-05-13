package modes

import (
	"fmt"
	"time"
)

type  Yesterday struct {
	YestDa  string  	`json:"yestday"`	//昨日首咨
}
type  Leyu struct {
	Ly		string 		`json:"leyu"`		//乐语注册
}

type Counts struct {
	 Count string `json:"count"`		//首咨投递数量
}
type Assigs struct {
	 Assig string `json:"assig"`		//已分配数量
}
type Report struct {
	Click		string 		`json:"click" `			// 	点击数量
	Cost		string		`json:"cost"`			//	所花费用
}
type Gather struct {
	Cout 	Counts 		`json:"cout"`		//首咨
	Repo	Report 		`json:"repo"`		//点击数量,所花费用
	Assi 	Assigs		`json:"assi"`		//首咨已经分配数量
	YestDay	Yesterday	`json:"yestday"`	//昨日首咨
	Ly 		Leyu		`josn:"leyu"`		//乐语
}
//type BaiDu struct {
//	Cost	string		`json:"cout"`		//百度费用
//}
//首咨已分配数量
func (this *Assigs)Assigned(s_time,e_time int64)([]Assigs,error){
	list := make([]Assigs,0)
	selec := "select count(1) as assig FROM cha_nat"
	where := fmt.Sprintf("where at >= %d AND at <= %d AND start = 1",s_time,e_time)
	sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
	if nil== err {
		for i,_ :=range sil {
			this.Assig = string(sil[i]["assig"])
			list = append(list, *this)
		}
	}
	fmt.Println(list)
	return list,err
}
//首咨投递数量
func (this *Counts)Countone(s_time,e_time int64)([]Counts,error){
	list := make([]Counts,0)
	selec := "select count(1) as count FROM cha_nat"
	where := fmt.Sprintf("where at >= %d AND at <= %d",s_time,e_time)
	sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
	if nil== err {
		for i,_ :=range sil {
			this.Count = string(sil[i]["count"])
			list = append(list, *this)
		}
	}
	return list,err
}

//点击数量、所花费用
func (this *Report)Reports(s_time,e_time int64)([]Report,error) {
	list := make([]Report,0)
	selec := "select  sum(click) AS click,sum(cost) AS cost FROM sics_data"
	where :=fmt.Sprintf("where at >= %d AND at <= %d",s_time,e_time)
	sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
	if nil== err {
		for i,_ :=range sil {

			this.Click 		= string(sil[i]["click"])
			this.Cost 		= string(sil[i]["cost"])

			list = append(list, *this)
		}
	}
	return list,err
}
//工作台数据分析
func ( this *Gather )Get(s_time,e_time int64)([]Gather,error){
	list := make([]Gather,0)
	this.Cout.Countone(s_time,e_time)
	this.Repo.Reports(s_time,e_time)
	this.Assi.Assigned(s_time,e_time)
	this.YestDay.YestDay()
	this.Ly.Leyu()
	list = append(list,*this)
	return  list,nil
}
func GetZero(checkTime int64) int64 {
	checkTime -= (checkTime + 28800) % 86400
	return checkTime
}
func (this *Yesterday)YestDay()([]Yesterday,error){
	list := make([]Yesterday,0)
	s_time := GetZero(time.Now().Unix()) - 86400
	e_time := time.Now().Unix()
	selec := "select count(1) as count FROM cha_nat"
	where := fmt.Sprintf("where at >= %d AND at <= %d",s_time,e_time)
	sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
	if nil== err {
		for i,_ :=range sil {
			this.YestDa = string(sil[i]["count"])
			list = append(list, *this)
		}
	}
	return list,err
}
//乐语注册数量	courses

func (this *Leyu)Leyu() ([]Leyu,error) {
	list := make([]Leyu,0)

	s_time := GetZero(time.Now().Unix())
	e_time := s_time + 86400
	selec := "select count(1) as count FROM cha_nat"
	where := fmt.Sprintf("where at >= %d AND at <= %d AND courses = 2",s_time,e_time)
	sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
	if nil== err {
		for i,_ :=range sil {
			this.Ly = string(sil[i]["leyu"])
			list = append(list, *this)
		}
	}
	fmt.Println(list)
	return list,err
}


