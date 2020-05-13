package modes

import (
		"fmt"
		"time"
)

/*
 * 描述:推广端
 * 表名: prothe
 *	start	 0 启用  1 禁用
 ********************************************************************/
type Prothe struct{
	Id		int64		`json:"prothe_id" xorm:"id"`				// 表 ID
	Name    string  	`json:"prothe_name" xorm:"name"`			// 推广端
	Start	int			`json:"prothe_start" xorm:"start"`			// 推广端状态
	At 		int64 		`json:"prothe_at" xorm:"at"`				//创建时间
	Creator	string 		`json:"prothe_creator" xorm:"creator"`		//创建人
}

func (this *Prothe) TableName() string {
	return  "prothe"
}

func ( this *Prothe )Save()( int64, error ){
	this.At = time.Now().Unix()
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Prothe)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Prothe)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Prothe)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Prothe )where( fage, count, page int, where, field string )( []Prothe,error ){
	list := make( []Prothe, 0 )
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

func ( this *Prothe )StartList( count, page int )( []Prothe,error ){
	// swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, "", "" )
}
//条件返回
func ( this *Prothe )DisPlay( count, page int )( []Prothe,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, swhere, "" )
}