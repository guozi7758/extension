package modes

 import (
	"fmt"
)

type TgData struct {
	SicsData 		`json:"tg_data" xorm:"extends"`
	Channel			`json:"chan" xorm:"extends"`
}
func ( this *TgData )Lookups(  val *Pager  ) error {
	list := make( []TgData, 0 )
	val.TName = "sics_data"
	val.Page--
	var field, sort string
	switch val.Coll {
	case 0: field = fmt.Sprintf("%s.at",	val.TName ) //时间排序
	case 1: field = fmt.Sprintf("%s.start",val.TName ) 	// 渠道
	default :field = fmt.Sprintf("%s.at",	val.TName ) // 时间排序
	}
	if 1 == val.Rules {
		sort = fmt.Sprintf( "%s asc", field )
	}else{
		sort = fmt.Sprintf( "%s desc", field )
	}
	val.Total,_ = Db( 0 ).Table(  val.TName ).Where( val.ToWhere() ).Count( this )
	err := Db(0).Table( val.TName ).
		Join( "INNER", "channel", fmt.Sprintf( "%s.start = channel.id", val.TName ) ).
		Where( val.ToWhere() ).
		OrderBy( sort ).
		Limit( val.Count, val.Page * val.Count ).
		Find( &list )
	if err == nil {
		val.List = list
	}
	return  err
}