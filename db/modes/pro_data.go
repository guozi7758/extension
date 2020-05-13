package modes

import (
		"fmt"
		// "time"
)

/*
 * 描述:url管理表
 *
 ********************************************************************/
type ProData struct{
	Id		 int64	`json:"pr_pro_id" xorm:"pro_id"`	//id 管理url管理表
	ClickNum int64	`json:"pr_click_num" xorm:"click_num"`	//点击数
	ShowNum	 int64	`json:"pr_show_num" xorm:"show_num"`	//展示数
	Submit	 int64  `json:"pr_submit" xorm:"submit"`	//提交量
	Start    int    `json:"pr_start" xorm:"start"`    //状态 0 显示 1 禁用
}

func (this *ProData) TableName() string {
	return  "prodata"
}
func (this *ProData )Save()( int64, error ){

	return Db(0).Insert(this)
}
func ( this *ProData)Get()( bool, error){
	return Db(0).Get(this)
}
func (this  *ProData)Update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}
func ( this *ProData) IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("pro_id = %d", this.Id)
	return this.Update(where,field)
}

func ( this *ProData )where( fage, count, page int, where, field string )( []ProData,error ){
	list := make( []ProData, 0 )
	var err error
	if field == "" {
		field = "pro_id"
	}
	if 0 == fage { // 从小到大排序
		err = Db(0).Where( where ).
			Asc( field ).
			Limit( count, page * count ).
			Find( &list )
	}else {	// 从大到小排序
		err = Db(0).Where( where ).
			Desc( field ).
			Limit( count, page * count ).
			Find( &list )
	}
	return list, err
}
func ( this *ProData )IdList( count, page int )( []ProData,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, swhere, "" )
}

func (this *ProData) List() ([]ProData, error) {
	list := make([]ProData, 0)
		err := Db(0).
			Find(&list)
	return list, err
}
