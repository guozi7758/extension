package modes

import (
	"fmt"
)

/*
 * 描述：角色表
 *
 *	state 0 启用  1 禁用
 *		
 ********************************************************************************************/
 type Role struct{
	Id	int64	`json:"ro_id" xorm:"id"`	//角色id 
	Name	string	`json:"ro_name" xorm:"name"`	//角色名称
	Start	int64	`json:"ro_start" xorm:"start"`	//状态 
}

func ( this *Role ) TableName() string {
	return "role"
}

func (this *Role )Save()(int64,error){
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Role)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Role)update( where string, field string )(int64, error){
	return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Role)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Role )where( fage, count, page int, where, field string )( []Role,error ){
	list := make( []Role, 0 )
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

func ( this *Role )StartList()( []Role,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, 100, 0, swhere, "" )
}

