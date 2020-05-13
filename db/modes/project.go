package modes

import (
	"fmt"

)

/*
 * 描述：项目表
 *
 *	id =  00    00       000       000     
 *	==================================================================================
 *           大区 事业部  项目一级 项目二级
 *
 *	state 0 启用  1 禁用
 *		
 * 报考条件: 报考条件为HTML文件，ID是表名
 * 
 *
 ********************************************************************************************/
type Project struct{
	Id int64	`json:"pr_id" xorm:"id"`	// 项目id
	Name string	`json:"pr_name" xorm:"name"`	// 项目名称
	Start int64	`json:"pr_start" xorm:"start"`	// 状态 0 正常 1
}


func (this *Project ) TableName() string {
	return "project"
}

func ( this *Project )Save()(int64,error){
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Project)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Project)update( where string, field string )(int64, error){
	return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Project)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Project )where( fage, count, page int, where, field string )( []Project,error ){
	list := make( []Project, 0 )
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

func ( this *Project )IdList(id int64)( []Project,error ){
	// start, end := PIdWhere( this.Id  )
	// swhere := fmt.Sprintf( "start = 0 and id > %d And id < %d", start, end )
	// return this.where( 0, 100, 0, swhere, "" )
	var where string
	if id == 0 {
		where = fmt.Sprintf( "id %% 100000 = 0 and start = 0" )
	}else if (id % 100000) > 0 {
		where = fmt.Sprintf( "id > %d and id < %d and start = 0", id, id + 1000 )
	}else{
		where = fmt.Sprintf( "id > %d and id < %d and start = 0 and id %% 1000 = 0", id, id + 100000 )
	}
	if id == 1 {
		where = fmt.Sprintf("id%%1000 > 0")
	}

	return this.where( 0, 100, 0, where, "" )
}
