package modes

import (
        "fmt"
	"time"
)

/*
 * 描述:url管理表
 * 表名: ProId
 *  Methods  0 电销 1 微信
 *  pro_the  0 PC 1 移动浏览器 2 小程序 3 APP
 *  Pro_way  0 搜索 1 信息流 2 APP
 *
 ********************************************************************/
type ProId struct{
	Id			int64	`json:"pro_id" xorm:"id"`				//url id 
	Ursine		string	`json:"pro_ursine" xorm:"ursine"`		//url路径
	Urlone		string	`josn:"pro_urlone" xorm:"urlone"`		//域名
	ChannelId	int64	`json:"pro_channelid" xorm:"channelid"`	// 渠道id
	MisoId		int64	`json:"pro_miso_id" xorm:"miso_id"`		// 项目id
	Account		string	`json:"pro_account" xorm:"account"`		// 推广账户
	At			int64   `json:"pro_at" xorm:"at"`				// 创建时间 
	Methods		int		`json:"pro_methods" xorm:"methods"`		// 营销方式
	ProThe		int		`json:"pro_pro_the" xorm:"pro_the"`		// 推 广 端
	ProWay		int		`json:"pro_pro_way" xorm:"pro_way"`		// 推广方式
	KeyFile		string	`json:"pro_key_file" xorm:"key_file"`	// 文 件UR
	CName		string 	`json:"pro_c_name" xorm:"c_name"`		//员 工 ID
	Start		int64	`json:"pro_start" xorm:"start"`			//状态 0启用 1禁用
	Code		string 	`josn:"-" xorm:"code"`					//hash唯一值

}

func (this *ProId) TableName() string {
	return  "proid"
}

func ( this *ProId )Save()( int64, error ){
	this.Start = 0
	this.At	= time.Now().Unix()
	return Db(0).Insert(this)
}

func ( this *ProId)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *ProId)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *ProId) IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *ProId )list_count( where string )(int64, error){
	return Db(0).Where( where ).Count( this )
}

func ( this *ProId )where( fage, count, page int, where, field string )( []ProId,error ){
	list := make( []ProId, 0 )
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
func ( this *ProId )AccountList( count, page int )( []ProId,int64,error ){
	list := make( []ProId, 0 )
	var err error
	var ncoun int64
	swhere := fmt.Sprintf( "account = %s and channelid = %d", this.Account,this.ChannelId)
	ncoun, err = this.list_count( swhere )
	if err == nil || ncoun > 0 {
		list, err = this.where( 1, count, page, swhere, "" )
	}
	return list, ncoun, err
}
