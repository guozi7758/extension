package modes

import (
	"fmt"
	"time"
)
/*
 * 描述：员工表
 *
 ********************************************************************************************/
type Employee struct{
	Id      int64	`json:"emp_id" xorm:"id"`			// 员 工 id
	RoleId  int64   `json:"emp_role_id" xorm:"role_id"`	// 角 色 id
	TeamId  int64   `json:"emp_team_id" xorm:"team_id"`	// 小 组 id
	ProjId  int64   `json:"emp_proj_id" xorm:"proj_id"`	// 项 目 id
	Name    string  `json:"emp_name" xorm:"name"`		// 员工名称
	Phone   string  `json:"emp_phone" xorm:"phone"`		// 账    号
	Pass    string  `json:"-" xorm:"pass"`				// 密    码
	TQAcc   int     `json:"emp_tq_acc" xorm:"tq_acc"`	// TQ座席编号
	TQPw    string  `json:"emp_tq_pw" xorm:"tq_pw"`		// TQ 秘 密
	Start   int		`json:"emp_start" xorm:"start"`		// 状    态
	At		int64   `json:"pro_at" xorm:"at"`			// 创建时间
}

func (this *Employee ) TableName() string {
	return "employee"
}

func ( this *Employee )Save()(int64,error){
	this.At	= time.Now().Unix()
	return Db(0).Insert(this)
}

func ( this *Employee)Get()( bool, error){
	if _, err := Db(0).Get(this); err !=nil {
		return false, err
	}
	if this.RoleId != 0 {
		if this.RoleId % 100 == 3 || this.RoleId %100 ==0 {
			return true,nil
		}
		return false,nil
	}
	return true,nil
	//return Db(1).Get(this)
}

func (this  *Employee)update( where string, field string )(int64, error){
	return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Employee)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Employee )where( fage, count, page int, where, field string )( []Employee,error ){
	list := make( []Employee, 0 )
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

func ( this *Employee )StartList( count, page int )( []Employee,error ){
	swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, swhere, "" )
}
