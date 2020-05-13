package proway

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
)

/*
 描述 ： 添加推广方式

 类型 ： GET
 ******************************************************************************/
func ProwayAdd( c *gin.Context ){
	var proway modes.Proway
	proway.Name = c.Query( "name" )
	proway.Creator = c.Query("creator")
	if _, err := proway.Save(); nil != err {
		fmt.Println( "Chanquery  err :", err.Error() )
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"添加推广方式重复" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "推广方式添加成功"})
}

/*
 描述 ： 修改推广方式

 类型 ： GET
 ******************************************************************************/
func ProwaySet( c *gin.Context ) {
	var proway modes.Proway
	proway.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) //必传
	proway.Name = c.Query("name")  //必传
	proway.Start ,_ = strconv.Atoi( c.Query("start") )  //必传
	if _, err := proway.IdSet("name,start"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"修改推广方式重复" })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}
/*
 *
 * 描述: 按状态返回推广方式列表
 *
 ********************************************************************/
func ProwayList( c *gin.Context ) {
	var proway modes.Proway
	// proway.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := proway.StartList( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}


//	描述: 按状态返回渠道列表
func DisPlay( c*gin.Context ){
	var proway modes.Proway
	proway.Start ,_ = strconv.Atoi( c.PostForm("start") )
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := proway.DisPlay( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}



