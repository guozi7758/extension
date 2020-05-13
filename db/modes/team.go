package modes

import (
        "fmt"
)
/*
 * 描述：小组表
 *
 *	state 0 启用  1 禁用
 *		
 ********************************************************************************************/
type Team struct{
	Id	int64	`json:"te_id" xorm:"id"`				//小组id 
	ProId	int64	`json:"te_pro_id" xorm:"pro_id"`	//项目id 
	Name	string	`json:"te_name" xorm:"name"`		//小组名称
	Start	int64	`json:"te_start" xorm:"start"`		//状态 
}
func (this *Team )TableName() string {
	return "team"
}

func (this *Team )Save()(int64,error){
	return Db(0).Insert(this)
}

func ( this *Team)Get()( bool, error){
	return Db(0).Get(this)
}

func (this *Team)update( where, field string )(int64, error){
	return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Team)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Team )where( fage, count, page int, where, field string )( []Team,error ){
	list := make( []Team, 0 )
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

func ( this *Team )PIdList()( []Team,error ){
	var swhere string
	if this.ProId % 1000 != 0 {
		swhere = fmt.Sprintf( "start = 0 AND pro_id = %d", this.ProId )
	}else{
		start, end := PIdWhere( this.ProId )
		swhere = fmt.Sprintf( "start = 0 AND pro_id > %d and pro_id < %d", start, end )
	}
	return this.where( 0, 200, 0, swhere, "" )
}

