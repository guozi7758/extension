package methods

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
)

/*
 描述 ： 添加营销方式
 
 类型 ： GET
 ******************************************************************************/
func MethodsAdd( c *gin.Context ){
	var methods modes.Methods
	methods.Name = c.Query( "name" )
	methods.Creator = c.Query("creator")
	if _, err := methods.Save(); nil != err {
		fmt.Println( "Chanquery  err :", err.Error() )
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"添加营销方式重复" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "营销方式添加成功"})
}

/*
 描述 ： 修改营销方式

 类型 ： GET
 ******************************************************************************/
func MethodsSet( c *gin.Context ) {
	var methods modes.Methods
	methods.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) //必传
	methods.Name = c.Query("name")  //必传
	methods.Start ,_ = strconv.Atoi( c.Query("start") )  //必传
	if _, err := methods.IdSet("name,start"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"修改营销方式重复" })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}
/*
 *
 * 描述: 营销方式列表
 *
 ********************************************************************/
func MethodsList( c *gin.Context ) {
	var methods modes.Methods
	// methods.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := methods.StartList( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}
//按状态返回列表
func DisPlay( c*gin.Context ){
	var methods modes.Methods
	methods.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := methods.DisPlay( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}



