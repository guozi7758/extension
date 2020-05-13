package modes

import (
	//"time"
	//"../../lib"
	"fmt"
)
/*
 * 描述:推广数据报表
 * 表名: sics_data
 *	start	 0 启用  1 禁用
 ********************************************************************/
type SicsData struct{
	Id			int64		`json:"bd_id" xorm:"id"`					// 	ID
	PId			int64		`json:"bd_pid" xorm:"pid"`					// 	项目ID
	Name   		string  	`json:"bd_name" xorm:"name"`				// 	账户名称
	At			int64 		`json:"bd_at" xorm:"at"`					// 	插入时间
	StartAt		string 		`josn:"bd_start_at" xorm:"start_at"`		// 	统计时间
	Impression	int64		`json:"bd_impression" xorm:"impression"`	// 	展示次数
	Click		int64 		`json:"bd_click" xorm:"click"`				// 	点击数量
	Cost		int64		`json:"bd_cost" xorm:"cost"`				//	所花费用
	Ctr			int64		`json:"bd_ctr" xorm:"ctr"`					//	点击率
	Cpc			int64		`json:"bd_cpc" xorm:"cpc"`					//	平均点击价格
	Cpm			int64		`json:"bd_cpm" xorm:"cpm"`					//	单元出价
	Start 		int			`josn:"bd_start" xorm:"start"`				//类别  1.百度  2.360  3.搜狗 4.神马
}



func (this *SicsData) TableName() string {
	return  "sics_data"
}
func ( this *SicsData )Save()( int64, error ){
	return Db(0).Insert(this)
}
func ( this *SicsData)Get(s_time,e_time int64)( bool, error){
	//昨日工作台数据
	where := fmt.Sprintf("at >= %d and at <= %d", s_time,e_time)
	field := fmt.Sprintf("(sum(cost) AS cost)")
	return Db(0).Cols(field).Where(where).Get(this)
}

