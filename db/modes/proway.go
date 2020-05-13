package modes

import (
		"fmt"
		"time"
)

/*
 * 描述:推广方式
 * 表名: proway
 *	start	 0 启用  1 禁用
 ********************************************************************/
type Proway struct{
	Id		int64		`json:"proway_id" xorm:"id"`				// 表 ID
	Name    string  	`json:"proway_name" xorm:"name"`			// 推广方式
	Start	int			`json:"proway_start" xorm:"start"`			// 推广方式状态
	At 		int64 		`json:"proway_at" xorm:"at"`				//创建时间
	Creator	string 		`json:"proway_creator" xorm:"creator"`		//创建人
}

func (this *Proway) TableName() string {
	return  "proway"
}

func ( this *Proway )Save()( int64, error ){
	this.At = time.Now().Unix()
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Proway)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Proway)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Proway)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Proway )where( fage, count, page int, where, field string )( []Proway,error ){
	list := make( []Proway, 0 )
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

func ( this *Proway )StartList( count, page int )( []Proway,error ){
	// swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, "", "" )
}
//条件显示
func ( this *Proway )DisPlay( count, page int )( []Proway,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start)
	return this.where( 1, count, page, swhere, "" )
}