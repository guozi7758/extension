package modes

import (
	"fmt"
	// "time"
	// "strings"
)

type SetChanat struct {
	Cha  ChaNat 	`json:"cha" xorm:"extends"`
	// Channel 	`json:"chan" xorm:"extends"`
	//Prot Prothe 	`json:"prot" xorm:"extends"`
	//Prow Proway 	`json:"prow" xorm:"extends"`
	// Proj Project	`json:"proj" xorm:"extends"`
}

func ( this *SetChanat )ChanatLists(  val *Pager  ) error {
	list := make( []SetChanat, 0 )
	val.TName = "cha_nat"
	val.Page--
	var field, sort string
	switch val.Coll {
		case 0: field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
		case 1: field = fmt.Sprintf("%s.pro_id",	val.TName ) // 项目
		case 2: field = fmt.Sprintf("%s.start",	val.TName ) // 状态
		default :field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
	}
	if 1 == val.Rules {
		sort = fmt.Sprintf( "%s asc", field )
	}else{
		sort = fmt.Sprintf( "%s desc", field )
	}
	val.Total,_ = Db( 0 ).Table(  val.TName ).Where( val.ToWhere() ).Count( this )
	err := Db(0).Table( val.TName ).
		//Join( "INNER", "channel", fmt.Sprintf( "%s.channelid = channel.id", val.TName ) ).
		//Join( "INNER", "prothe", fmt.Sprintf( "%s.pro_the = prothe.id", val.TName ) ).
		//Join( "INNER", "proway", fmt.Sprintf( "%s.pro_way = proway.id", val.TName ) ).
		Where( val.ToWhere() ).
		OrderBy( sort ).
		Limit( val.Count, val.Page * val.Count ).
		Find( &list )
	if err == nil {
		val.List = list
	}
	return  err
}

