package modes

import (
		"fmt"
		"time"
)

/*
 * 描述:渠道表
 * 表名: channel
 *	start	 0 启用  1 禁用
 ********************************************************************/
type Channel struct{
	Id		int64		`json:"cha_id" xorm:"id"`				// 表 ID
	Name    string  	`json:"cha_name" xorm:"name"`			// 渠道名称
	Start	int			`json:"cha_start" xorm:"start"`			// 渠道状态
	At 		int64 		`json:"cha_at" xorm:"at"`				//创建时间
	Creator	string 		`json:"cha_creator" xorm:"creator"`		//创建人
}

func (this *Channel) TableName() string {
	return  "channel"
}

func ( this *Channel )Save()( int64, error ){
	this.At = time.Now().Unix()
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Channel)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Channel)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Channel)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Channel )where( fage, count, page int, where, field string )( []Channel,error ){
	list := make( []Channel, 0 )
	var err error
	if field == "" {
		field = "id"
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

func ( this *Channel )StartList( count, page int )( []Channel,error ){
	// swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, "", "" )
}
func( this *Channel)DisPlay( count, page int)([]Channel,error){
	swhere := fmt.Sprintf("start =%d",this.Start)
	return this.where( 1, count, page, swhere, "" )
}
