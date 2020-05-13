package modes

import (
		"fmt"
		"time"
)

/*
 * 描述:营销方式
 * 表名: methods
 *	start	 0 启用  1 禁用
 ********************************************************************/
type Methods struct{
	Id		int64		`json:"methods_id" xorm:"id"`				// 表 ID
	Name    string  	`json:"methods_name" xorm:"name"`			// 营销方式
	Start	int			`json:"methods_start" xorm:"start"`			// 营销方式状态
	At 		int64 		`json:"methods_at" xorm:"at"`				//创建时间
	Creator	string 		`json:"methods_creator" xorm:"creator"`		//创建人
}

func (this *Methods) TableName() string {
	return  "methods"
}

func ( this *Methods )Save()( int64, error ){
	this.At = time.Now().Unix()
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Methods)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Methods)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Methods)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Methods )where( fage, count, page int, where, field string )( []Methods,error ){
	list := make( []Methods, 0 )
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

func ( this *Methods )StartList( count, page int )( []Methods,error ){
	// swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, "", "" )
}

//条件显示
func ( this *Methods )DisPlay( count, page int )( []Methods,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, swhere, "" )
}