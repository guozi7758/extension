package modes

import (
	"fmt"
	"time"
	//"reflect"
	"github.com/Luxurioust/excelize"
)

/*
 * 描述:推广机会属性表
 * 	 s_type  0 搜索推广 1 乐语推广 2 微信推广 3 自然流量
 * 	 education 0 未提交学历， 1 大专 2 中专 3 本科 4 其他
 *
 ********************************************************************/
type ChaNat struct{
	Id		int64		`json:"net_id" xorm:"id"`		//机会id
	Phone		int64		`json:"net_phone" xorm:"phone"`		//手机号
	ProWay		int64		`json:"net_pro_way" xorm:"pro_way"`	// 推广方式
	ChannelId	int64		`json:"net_channelid" xorm:"channelid"`	// 渠道id
	ProThe		int64		`json:"net_pro_the" xorm:"pro_the"`	// 推 广 端
	Education	int64		`json:"net_education" xorm:"education"`	//学历	-- 0 无学历, 1 中专 2 大专， 3 本科 4 研究生
	City		int64		`json:"net_city" xorm:"city"`		//城市ADCODE
	UId			int64		`json:"net_uid" xorm:"uid"`		//首次指派坐席
	Pro_id		int64		`json:"net_pro_id" xorm:"pro_id"`	//项目id
	Start		int		`json:"net_start" xorm:"start"`		//状态( 0 未分配 1 已分配 2 已成交)
	At			int64		`json:"net_at" xorm:"at"`		//创建时间
	Url_code	string		`josn:"net_url_code" xorm:"url_code"`	//url
	Courses		string		`json:"net_courses" xorm:"courses"`	//
	Major		string		`json:"net_major" xorm:"major"`		//专业
	Uname		string		`json:"net_uname" xorm:"uname"`		//坐席姓名
	Pro_name	string		`json:"net_pro_name" xorm:"pro_name"`	//项目名称
}

func (this *ChaNat) TableName() string {
	return  "cha_nat"
}

func ( this *ChaNat )Save()( int64, error ){
	this.At = time.Now().Unix()
	return Db(0).Insert(this)
}

func ( this *ChaNat)Get()( bool, error){
	return Db(0).Get(this)
}
func (this  *ChaNat)delete( where string)(int64, error){
	return Db(0).Where(where).Delete(this)
}
func (this * ChaNat)Del()(int64, error){
	where := fmt.Sprintf("id =%d",this.Id)
	return this.delete(where)
}

func (this  *ChaNat)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *ChaNat) IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	//this.At = time.Now().Unix()
	return this.update(where,field)
}

func ( this *ChaNat )list_count( where string )(int64, error){
	return Db(0).Where( where ).Count( this )
}

func ( this *ChaNat )where( fage, count, page int, where, field string )( []ChaNat,error ){
	list := make( []ChaNat, 0 )
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
func ( this *ChaNat )AccountList( count, page int )( []ChaNat,int64,error ){
	list := make( []ChaNat, 0 )
	var err error
	var ncoun int64
	swhere := fmt.Sprintf( "start = %d", this.Start)
	ncoun, err = this.list_count( swhere )
	if err == nil || ncoun > 0 {
		list, err = this.where( 1, count, page, swhere, "" )
	}
	return list, ncoun, err
}

func to_line( char string, line int )string{
	return fmt.Sprintf( "%s%d", char, line )
}

func ToString( i int64 )string{

	return fmt.Sprintf( "%d", i )

}
func ToStringint( i int )string{

	return fmt.Sprintf( "%d", i )

}
func ( this *ChaNat )GPassPhone()string{
	str := []byte( fmt.Sprintf( "%d", this.Phone ) )
	str[3] = '*'
	str[4] = '*'
	str[5] = '*'
	str[6] = '*'
	return string( str )
}
func ( this *ChaNat )ToExcelFile( val *Pager  )(string, error){
	list := make( []ChaNat, 0 )
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
	err := Db(0).Table(val.TName).
		Where(val.ToWhere()).
		OrderBy(sort).
		Limit(val.Count, val.Page*val.Count).
		Find(&list)
	if nil == err {
		xlsx := excelize.NewFile()
	        xlsx.NewSheet("Sheet1")
		xlsx.SetCellValue("Sheet1", to_line( "A", 1  ),"手机号")
		xlsx.SetCellValue("Sheet1", to_line( "B", 1  ),"推广方式")
		xlsx.SetCellValue("Sheet1", to_line( "C", 1  ),"渠道")
		xlsx.SetCellValue("Sheet1", to_line( "D", 1  ),"推广端")
		xlsx.SetCellValue("Sheet1", to_line( "E", 1  ),"行政区号")
		xlsx.SetCellValue("Sheet1", to_line( "F", 1  ),"当前状态")
		xlsx.SetCellValue("Sheet1", to_line( "H", 1  ),"投递时间")
		xlsx.SetCellValue("Sheet1", to_line( "I", 1  ),"账户id")
		xlsx.SetCellValue("Sheet1", to_line( "J", 1  ),"当前坐席")
		xlsx.SetCellValue("Sheet1", to_line( "K", 1  ),"项目名称")
		xlsx.SetCellValue("Sheet1", to_line( "L", 1  ),"投递地址")
		for i, v := range list {
			xlsx.SetCellValue("Sheet1", to_line( "A", i + 2 ), v.Phone )
			//xlsx.SetCellValue("Sheet1", to_line( "A", i + 2 ), v.Phone )
			xlsx.SetCellValue("Sheet1", to_line( "B", i + 2 ), ToString( v.ProWay ) )
			xlsx.SetCellValue("Sheet1", to_line( "C", i + 2 ), ToString( v.ChannelId ))
			xlsx.SetCellValue("Sheet1", to_line( "D", i + 2 ), ToString( v.ProThe ) )
			xlsx.SetCellValue("Sheet1", to_line( "E", i + 2 ), ToString( v.City ) )
			xlsx.SetCellValue("Sheet1", to_line( "F", i + 2 ), ToStringint( v.Start ) )
			xlsx.SetCellValue("Sheet1", to_line( "H", i + 2 ), ToString( v.At ) )
			xlsx.SetCellValue("Sheet1", to_line( "I", i + 2 ), v.Courses )
			xlsx.SetCellValue("Sheet1", to_line( "J", i + 2 ), v.Uname )
			xlsx.SetCellValue("Sheet1", to_line( "K", i + 2 ), v.Pro_name )
			xlsx.SetCellValue("Sheet1", to_line( "L", i + 2 ), v.Url_code )
		}
		filepath :=  fmt.Sprintf( "./upload/%d.xlsx",time.Now().Day())
		_ = xlsx.SaveAs(filepath)
		return filepath, err
	}
	return "", err
}
