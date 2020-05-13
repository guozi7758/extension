package chann

import (
	 "fmt"
	 "strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
)

/*
 描述 ： 添加推广渠道
 路由 :  chann/add
 类型 ： GET
 ******************************************************************************/
func ChannAdd( c *gin.Context ){
	var chans modes.Channel
	chans.Name = c.Query( "name" )
	chans.Creator = c.Query("creator")
	if _, err := chans.Save(); nil != err {
		fmt.Println( "Chanquery  err :", err.Error() )
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"添加渠道重复" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "渠道添加成功"})
}

/*
 描述 ： 修改推广渠道
 路由 :  chann/set
 类型 ： GET
 ******************************************************************************/
func ChannSet( c *gin.Context ) {
	var chans modes.Channel
	chans.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) //必传
	chans.Name = c.Query("name")  //必传
	chans.Start ,_ = strconv.Atoi( c.Query("start") )  //必传
	if _, err := chans.IdSet("name,start"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1,"msg":"修改渠道重复" })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}
/*
 *
 * 描述:渠道列表
 *
 ********************************************************************/
func ChannList( c *gin.Context ) {
	var chans modes.Channel
	// chans.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := chans.StartList( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}

//	描述: 按状态返回渠道列表
func DisPlay( c*gin.Context ){
	var chans modes.Channel
	chans.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := chans.DisPlay( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}

