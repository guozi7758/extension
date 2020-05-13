package modes

import (
	"fmt"
)
type BaiduFirst struct {
	BaiduCount 	string `json:"baiducount"`	//投递数量
}

type BaiduSum struct {
	Bdsum		string	`json:"baiducount"`	//投递数量
}

//百度投递数量
func (this *BaiduFirst)BaiduCountone(s_time,e_time,channelid int64)([]BaiduFirst,error){
	list := make([]BaiduFirst,0)
	selec := "select count(1) as baiducount FROM cha_nat"
	fmt.Println(channelid)
		where := fmt.Sprintf("where channelid = %d AND at >= %d AND at <= %d",channelid,s_time,e_time)
		sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
		if nil== err {
			for i,_ :=range sil {
				this.BaiduCount = string(sil[i]["baiducount"])
				list = append(list, *this)
			}
		}

	return list,nil
}
func (this *BaiduSum)BaiduSum(s_time,e_time int64)([]BaiduSum,error){
	list :=make([]BaiduSum,0)
	selec := "select count(1) as baiducount FROM cha_nat"
		where := fmt.Sprintf("where at >= %d AND at <= %d",s_time,e_time)
		sil,err :=Db(0).Query(fmt.Sprintf(" %s %s",selec,where))
		if nil== err {
			for i,_ :=range sil {
				this.Bdsum = string(sil[i]["baiducount"])
				list = append(list, *this)
			}
		}
		return list,nil
}
