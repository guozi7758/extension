package pro_url

import (
	"../../db/modes"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net"
	"strings"
	"net/http"
	"strconv"
	//"fmt"

)

/*
 测试 ： url管理
 路由 :  pro_url/add
 类型 ： GET
*/
func ProIdAdd( c *gin.Context ){
	var val modes.ProId
	jwt_user := modes.JwtUser(c)
	a :=0
	val.Code = strconv.Itoa(a)
	val.CName =jwt_user.Name
	val.ChannelId ,_ = strconv.ParseInt(c.PostForm("channel_id"),10,64)
	val.MisoId ,_ = strconv.ParseInt(c.PostForm("miso_id"),10,64)
	val.Account	= c.PostForm("account")
	val.Methods ,_ = strconv.Atoi( c.PostForm("methods") )
	val.ProThe ,_ = strconv.Atoi( c.PostForm("prothe") )
	val.ProWay ,_ = strconv.Atoi( c.PostForm("proway") )
	//url	:= c.PostForm("ursine")					//文件地址
	val.Urlone = c.PostForm("urlone")			//域名
	val.KeyFile = c.PostForm("file")			//图片地址
	comma := strings.Index(val.KeyFile, "/")
	keyfile := fmt.Sprintf(val.KeyFile[comma:])
	val.Ursine	=	val.Urlone + keyfile
	fmt.Println(val.Ursine)

	if _, err := val.Save();nil != err {
		c.JSON(http.StatusOK,gin.H{"err":1 , "msg":err.Error()})
		return
	}
	//fileip := "106.13.79.158"
	//conn,err := net.Dial("ip:icmp",val.Urlone)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//add := conn.RemoteAddr()
	//ip ,_:=fmt.Println(add)
	//urlip := strconv.Itoa(ip)
	//if urlip != fileip{
	//	c.JSON(http.StatusOK, gin.H{"err": 2,"msg": "此域名未解析到本服务器"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "ok"})
}
//修改
func ProIdSet( c *gin.Context ){
	var pro_url modes.ProId
	jwt_user := modes.JwtUser( c )
	pro_url.Id ,_ = strconv.ParseInt(c.Query("id"),10,64)
	pro_url.Ursine	= c.PostForm("ursine")
	pro_url.ChannelId ,_ = strconv.ParseInt(c.PostForm("channel_id"),10,64)
	pro_url.MisoId ,_ = strconv.ParseInt(c.PostForm("miso_id"),10,64)
	pro_url.Account	= c.PostForm("account")
	pro_url.Methods ,_ = strconv.Atoi( c.PostForm("methods") )
	pro_url.ProThe ,_ = strconv.Atoi( c.PostForm("prothe") )
	pro_url.ProWay ,_ = strconv.Atoi( c.PostForm("proway") )
	pro_url.CName  =  jwt_user.Name
	pro_url.KeyFile =  c.PostForm("file")
	if _, err := pro_url.IdSet("ursine,account,channelid,miso_id,methods,prothe,proway,uid,file"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})

}
//项目 id 返回项目	name 
func ProjectList( c *gin.Context ){
	var val modes.Project
	id,_ := strconv.ParseInt( c.PostForm( "id" ), 10, 64 )
	list, err := val.IdList(id )
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
	}
        c.JSON(http.StatusOK, gin.H{"err": 0,"list": list })
}

func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}
//url管理列表展示加搜索
func ProIdLookup(  c *gin.Context  ) {
	var val modes.SetProId
	var page modes.Pager
	BuildPager(c,&page)
	if c.PostForm("start_time") != "" {
		page.Add("at",">","and",c.PostForm("start_time"))
	}
	if c.PostForm("end_time") != ""{
		page.Add("at","<","and",c.PostForm("end_time"))
	}
	if c.PostForm("proway") != ""{
		page.Add("pro_way","=","and",c.PostForm("proway"))
	}
	if c.PostForm("prothe") != ""{
		page.Add("pro_the","=","and",c.PostForm("prothe"))
	}
	if c.PostForm("miso_id") != ""{
		page.Add("miso_id","=","and",c.PostForm("miso_id"))
	}
	if c.PostForm("channelid") != ""{
		page.Add("channelid","=","and",c.PostForm("channelid"))
	}
	if c.PostForm("methods") != ""{
		page.Add("methods","=","and",c.PostForm("methods"))
	}

	if err := val.Lookups(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})

}


//个人信息
func UserInfo(c *gin.Context){
	var personal modes.SetPersonal
	jwt_user := modes.JwtUser( c )
	uid  := jwt_user.Id
	list ,err := personal.Personal(uid)
	if nil != err{ 
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0,"list":list})

}
/*
	修改密码
 */
func PassInfoSet(c *gin.Context)  {
	var val modes.Employee
	jwt_user := modes.JwtUser(c)
	val.Id = jwt_user.Id
	val.Get()
	if val.Pass != c.Query("pass"){
		c.JSON(http.StatusOK,gin.H{"err":1,"msg":"密码错误"})
		return
	}
	Passone := c.Query("passone")
	if Passone == "" {
		c.JSON(http.StatusOK,gin.H{"err":2 , "msg": "密码不可为空"})
		return
	}
	Passtwo := c.Query("passtwo")
	if Passtwo == "" {
		c.JSON(http.StatusOK,gin.H{"err":3 , "msg": "新密码不可为空"})
		return
	}
	if Passone != Passtwo {
		c.JSON(http.StatusOK,gin.H{"err":4 , "msg": "检查密码与新密码是否相同"})
		return
	}
	val.Pass = Passtwo
	fmt.Println("*************",val)
	if  _,err := val.IdSet("pass");nil != err {
		c.JSON(http.StatusOK,gin.H{"err":5 , "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0,"msg": "密码修改成功"})

}


 
