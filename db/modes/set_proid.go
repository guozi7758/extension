package modes

 import (
	"fmt"
)

type SetProId struct {
	ProId 		`json:"proid" xorm:"extends"`
	Channel 	`json:"chan" xorm:"extends"`
	Methods 	`json:"meth" xorm:"extends"`
	Prothe 		`json:"pthe" xorm:"extends"`
	Project		`json:"ject" xorm:"extends"`
	Proway 		`json:"pway" xorm:"extends"`

}
func ( this *SetProId )Lookups(  val *Pager  ) error {
	list := make( []SetProId, 0 )
	val.TName = "proid"
	val.Page--
	var field, sort string
	switch val.Coll {
	case 0: field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
	case 1: field = fmt.Sprintf("%s.miso_id",	val.TName ) // 项目
	case 2: field = fmt.Sprintf("%s.state",	val.TName ) // 状态
	default :field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
	}
	if 1 == val.Rules {
		sort = fmt.Sprintf( "%s asc", field )
	}else{
		sort = fmt.Sprintf( "%s desc", field )
	}
	val.Total,_ = Db( 0 ).Table(  val.TName ).Where( val.ToWhere() ).Count( this )
	err := Db(0).Table( val.TName ).
		Join( "INNER", "channel", fmt.Sprintf( "%s.channelid = channel.id", val.TName ) ).
		Join( "INNER", "methods", fmt.Sprintf( "%s.methods = methods.id", val.TName ) ).
		Join( "INNER", "prothe", fmt.Sprintf( "%s.pro_the = prothe.id", val.TName ) ).
		Join( "INNER", "project", fmt.Sprintf( "%s.miso_id = project.id", val.TName ) ).
		Join( "INNER", "proway", fmt.Sprintf( "%s.pro_way = proway.id", val.TName ) ).
		Where( val.ToWhere() ).
		OrderBy( sort ).
		Limit( val.Count, val.Page * val.Count ).
		Find( &list )
	if err == nil {
		val.List = list
	}
	return  err
	
}


type SetPersonal  struct {
	Employee	`xorm:"extends"`
	Project   `xorm:"extends"`
	Team      `xorm:"extends"`
	Role		`xorm:"extends"`

}
// //个人资料
func (this *SetPersonal)Personal(id int64  ) ([]SetPersonal,error){
		list :=make([]SetPersonal,0)
		var err	error 
		err = Db(1).Table("employee").
		Join( "INNER", "project", "employee.proj_id = project.id" ).
		Join( "INNER", "team", "employee.team_id = team.id" ).
		Join("INNER","role","employee.role_id = role.id").
		Where("employee.id = ?", id).
		Find(&list)
		return list,err
}