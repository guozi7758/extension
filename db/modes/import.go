package modes

 import	(
	"time"
	"fmt"
 )

/*
 * 描述:导入数据
 * 	
 *
 ********************************************************************/
 
type Import struct{
	Id 			int64		`json:"im_id" xorm:"id" `				// 'id'
	Pid			int64		`json:"im_pid" xorm:"pid" `				// 项目id
	CName		string		`json:"im_c_name" xorm:"c_name" `		// 操作人
	At			int64		`json:"im_at" xorm:"at" `				// 插入时间
	Start		int64		`json:"im_start" xorm:"start" `			// 是否执行/*0未执行	1已执行*/
	ImCount		int64		`json:"im_im_count" xorm:"im_count" `	// 导入总条数
	Wrong 		int64 		`json:"im_wrong" xorm:"wrong" `			// 错误条数
	Actual		int64		`json:"im_actual" xorm:"actual" `		// 实际执行条数
	Exec_at		int64		`json:"im_exec_at" xorm:"exec_at" `		// 执行时间
	FileExec	string		`json:"im_file_exec" xorm:"file_exec" `	// execl文件
	Code 		string 		`json:"-" xorm:"code"`					// hashcode
}
 
func (this *Import) TableName() string {
	return  "import"
}

func ( this *Import )Save()( int64, error ){
	this.At = time.Now().Unix()
	this.Start = 0
	return Db(0).Insert(this)
}

func ( this *Import)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *Import)update( where string, field string )(int64, error){
	 return Db(0).Where(where).Cols(field).Update(this)
}
func ( this *Import )Count( where string )( int64, error ){
	var val Import
	return  Db(0).Where(where).Count(&val)
}

func ( this *Import)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}

func ( this *Import )where( fage, count, page int, where, field string )( []Import,error ){
	list := make( []Import, 0 )
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

func ( this *Import )StartList( count, page int )( []Import,error ){
	// swhere := fmt.Sprintf( "start = %d", this.Start )
	return this.where( 1, count, page, "", "" )
}
func( this *Import)DisPlay( count, page int)([]Import,error){
	swhere := fmt.Sprintf("start =%d",this.Start)
	return this.where( 1, count, page, swhere, "" )
}
func ( this *Import )List( val *Pager )error{
	val.TName = this.TableName()
	val.Total,_ = this.Count( val.ToWhere() )
	list, err := this.where( val.Rules, val.Count, val.Page-1, val.ToWhere(), "" )
	if err == nil {
		val.List = list
	}
	return err
}