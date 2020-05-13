package modes

import (
	"fmt"
	//  "strings"

	 // "time"
)

type SetAccount struct {
	Account		`json:"acco" xorm:"extends"`
	Channel 	`json:"chan" xorm:"extends"`
	Prothe 		`json:"pthe" xorm:"extends"`
	Proway 		`json:"pway" xorm:"extends"`
	Project		`json:"ject" xorm:"extends"`

}
func ( this *SetAccount )Lookups(  val *Pager  ) error {
	list := make( []SetAccount, 0 )
	val.TName = "account"
	val.Page--
	var field, sort string
	switch val.Coll {
		case 0: field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
		case 1: field = fmt.Sprintf("%s.pro_id",	val.TName ) // 项目
		case 2: field = fmt.Sprintf("%s.state",	val.TName ) // 状态
		default :field = fmt.Sprintf("%s.at",	val.TName ) // 状态
	}
	if 1 == val.Rules {
		sort = fmt.Sprintf( "%s asc", field )
	}else{
		sort = fmt.Sprintf( "%s desc", field )
	}
	val.Total,_ = Db( 0 ).Table(  val.TName ).Where( val.ToWhere() ).Count( this )
	err := Db(0).Table( val.TName ).
		Join( "LEFT", "channel", fmt.Sprintf( "%s.channelid = channel.id", val.TName ) ).
		Join( "LEFT", "prothe", fmt.Sprintf( "%s.prothe = prothe.id", val.TName ) ).
		Join( "LEFT", "proway", fmt.Sprintf( "%s.proway = proway.id", val.TName ) ).
		Join( "LEFT", "project", fmt.Sprintf( "%s.project = project.id", val.TName ) ).
		Where( val.ToWhere() ).
		OrderBy( sort ).
		Limit( val.Count, val.Page * val.Count ).
		Find( &list )
	if err == nil {
		val.List = list
	}
	return  err
	
}



