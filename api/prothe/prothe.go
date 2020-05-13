package prothe

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
)

/*
 描述 ： 添加推广端
 
 类型 ： GET
 ******************************************************************************/
func ProtheAdd( c *gin.Context ){
	var prothe modes.Prothe
	prothe.Name = c.Query( "name" )
	prothe.Creator = c.Query("creator")
	if _, err := prothe.Save(); nil != err {
		fmt.Println( "Chanquery  err :", err.Error() )
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"添加推广端重复" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "推广端添加成功"})
}

/*
 描述 ： 修改推广端

 类型 ： GET
 ******************************************************************************/
func ProtheSet( c *gin.Context ) {
	var prothe modes.Prothe
	prothe.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) //必传
	prothe.Name = c.Query("name")  //必传
	prothe.Start ,_ = strconv.Atoi( c.Query("start") )  //必传
	if _, err := prothe.IdSet("name,start"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"修改推广端重复" })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}
/*
 *
 * 描述: 推广端列表
 *
 ********************************************************************/
func ProtheList( c *gin.Context ) {
	var prothe modes.Prothe
	// prothe.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := prothe.StartList( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}

//	描述: 按状态返回渠道列表
func DisPlay( c*gin.Context ){
	var prothe modes.Prothe
	prothe.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := prothe.DisPlay( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}


