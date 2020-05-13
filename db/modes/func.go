package modes

import (
        "fmt"
        "github.com/gin-gonic/gin"
        "github.com/go-xorm/xorm"

	"../../db"
        "../../router/middleware"
)

// var Ur db.R
// func URedis( db int )error{
// 	return Ur.Init( db )
// }

func Db( index int )*xorm.Engine{
	return	db.UseHand(index)
}

/*
 * 描述: 打印Jwt结构
 *
 ***********************************************************************/
func PrintJwt( c *gin.Context,  fage string ){
        user := c.MustGet("claims").(*middleware.CustomClaims)
        fmt.Println("============\t\t", fage, "\t============================")
        fmt.Println("用户昵称:", user.Phone )
        fmt.Println("用户标示:", user.Id )
}

/*
 * 描述: 获取手机号
 *
 ***********************************************************************/
func GetPhone( c *gin.Context ) string {
        user := c.MustGet("claims").(*middleware.CustomClaims)
        return user.Phone
}

/*
 * 描述: 返回Jwt结构
 *
 ***********************************************************************/
func JwtUser(c *gin.Context) *middleware.CustomClaims {
        return c.MustGet("claims").(*middleware.CustomClaims)
}



func PIdWhere( id int64 )( int64, int64 ){
	if id == 0 {
                return 0, 9999999999
        }
        POne := id % 1000
        if POne != 0 {
		return id, id
	}
        PTow := ( id % 1000000 ) / 1000
        if PTow != 0 {
		return id, id + 1000
	}
        Care := ( id % 100000000 ) / 1000000
        if Care != 0 {
		return id, id + 1000000
	}
        return id, id + 100000000
}

type sql_where struct{
	Key   string    `json:"key"`	// Key 
	Fage  string	`json:"-"`	// 符号 =, >, <
	Tage  string	`json:"-"`	// 条件 or, and
	Val   string    `json:"value"`	// value
}

type Pager struct {
	Count	int		`json:"count"`		// 单页数量
	Page	int		`json:"page"`		// 页码
	Total	int64		`json:"total"`		// 本条件的总数量
	Coll    int		`json:"collation"`	// 排序字段
	Rules   int		`json:"rules"`		// true 从小到大排序，false 从大到小排序
	List    interface{}	`json:"list"`		// 数据列表
	TName	string		`json:"-"`		// 表名
	Where	[]sql_where	`json:"-"`		// 
}

func ( this *Pager )Add(  key, fage, tage, val string  ){
	this.Where = append( this.Where, sql_where{ key, fage, tage, val } )
}
func ( this *Pager )IdArea( field string, id int64 ){
	sid, eid := PIdWhere( id )
	this.Where = append( this.Where, sql_where{ field, ">=", "", fmt.Sprintf("%d",sid) } )
	this.Where = append( this.Where, sql_where{ field, "<", "and", fmt.Sprintf("%d",eid) } )
}
func ( this *Pager )GetValue( key string )string{
	for _, v := range this.Where {
		if  v.Key == key {
			return v.Val
		}
	}
	return ""
}
func ( this *Pager )Exist( key string )bool{
	for _, v := range this.Where {
		if v.Key == key {
			return true
		}
	}
	return false
}

func ( this *Pager )ToWhere()string{
	var key, where string
	if len( this.Where ) > 0 {
		key = fmt.Sprintf( "%s.%s", this.TName, this.Where[0].Key )
		where = fmt.Sprintf( "%s %s %s ",key, this.Where[0].Fage,this.Where[0].Val)
		for _, v := range this.Where[1:] {
			key = fmt.Sprintf( "%s.%s", this.TName, v.Key )
			where = fmt.Sprintf( "%s %s %s %s %s ", where, v.Tage, key, v.Fage, v.Val)
		}
	}
	return where
}
