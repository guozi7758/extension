package account

import (
	"../../db/modes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 描述 ： 添加推广账户

 类型 ： GET
 ******************************************************************************/
func AccountAdd( c *gin.Context ){
	var account modes.Account
	var err error
	jwt_user := modes.JwtUser( c )
	account.Name = c.Query( "name" )
	boole,err := account.Get()
	if boole == true{
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"账户已存在" })
		return
	}
	account.ChannelId ,_ = strconv.ParseInt(c.Query("channelid"),10,64)
	account.ProJect ,_ = strconv.ParseInt(c.Query("project"),10,64)
	account.Prothe ,_ = strconv.ParseInt(c.Query("prothe"),10,64)
	account.Proway ,_ = strconv.ParseInt(c.Query("proway"),10,64)
	account.Url = c.Query("url")
	account.UId  = jwt_user.Id
	if _, err = account.Save(); nil != err{
		c.JSON(http.StatusOK, gin.H{"err": 2,"msg": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "推广账户添加成功"})
	
}
//推广账户修改  GET
func AccountSet( c *gin.Context ) {
	var account modes.Account
	account.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) //必传
	account.ProJect ,_  = strconv.ParseInt(c.Query("project"),10,64) //项目id 
	account.Name = c.Query("name")  //推广账户
	account.ChannelId ,_  = strconv.ParseInt(c.Query("channelid"),10,64)
	account.Proway ,_  = strconv.ParseInt(c.Query("proway"),10,64)
	account.Prothe ,_  = strconv.ParseInt(c.Query("prothe"),10,64)
	account.Start ,_ = strconv.Atoi( c.Query("start") )  //必传
	account.Url = c.Query("url")
	if _, err := account.IdSet("name,start,project,channelid,proway,prothe,url"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}

//推广账户显示 POST


func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}
//推广账户显示 POST

func AccountList( c *gin.Context){
	var val modes.SetAccount
	var page modes.Pager
	jwt_user := modes.JwtUser( c )
	BuildPager( c,&page)
		uid :=strconv.FormatInt(jwt_user.Id,10)
	if 	uid != "" {	
		page.Add( "uid", "=", "and", uid)
	}
	if err := val.Lookups(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})
}

func Baidutime( c *gin.Context ){
	//var timess lib.FieldsData
	//time  := fmt.Sprintf("%d-%02d-%02d",time.Now().Year(),time.Now().Month() ,time.Now().Day()-1 )
	//
	////百度第一个账号
	//if err := timess.Time(time,time);err !=nil {
	//	c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"采集失败" })
	// 	return
	//}
	//c.JSON(http.StatusOK, gin.H{"err": 0,"msg":"ok" })

}