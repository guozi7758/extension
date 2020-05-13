package modes

import (
	"fmt"
	"time"
)

/*
 * 描述:推广账户
 * 	
 *
 ********************************************************************/
type Account struct{
	Id 			int64	`json:"ac_id" xorm:"id"`				//id
	UId 		int64	`json:"ac_uid" xorm:"uid"`				//用户id
	ProJect		int64	`json:"ac_project" xorm:"project"`		//项目id
	Name 		string 	`json:"ac_name" xorm:"name"` 			//推广账户 
	ChannelId 	int64	`json:"ac_channelid" xorm:"channelid"` 	//渠道id 
	Proway		int64	`json:"ac_proway" xorm:"proway"`			//推广方式
	Prothe		int64	`json:"ac_prothe" xorm:"prothe"`			//推广端
	Url			string	`json:"ac_url" xorm:"url"`				//url 
	At 			int64 	`json:"ac_at" xorm:"at"`				//时间
	Start 		int		`json:"ac_start" xorm:"start"`			//状态  0 启用 1 禁用
}

func (this *Account) TableName() string {
	return  "account"
}

func ( this *Account )Save()( int64, error ){
	this.At = time.Now().Unix()
	return Db(0).Insert(this)
}

func ( this *Account)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Account)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}

func ( this *Account) IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Account )list_count( where string )(int64, error){
	return Db(0).Where( where ).Count( this )
}

func ( this *Account )where( fage, count, page int, where, field string )( []Account,error ){
	list := make( []Account, 0 )
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
func ( this *Account )AccountList( count, page int )( []Account,int64,error ){
	list := make( []Account, 0 )
	var err error
	var ncoun int64
	swhere := fmt.Sprintf( "start = %d", this.Start)
	ncoun, err = this.list_count( swhere )
	if err == nil || ncoun > 0 {
		list, err = this.where( 1, count, page, swhere, "" )
	}
	return list, ncoun, err
}
