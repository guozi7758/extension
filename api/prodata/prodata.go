package prodata

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
)
/*
 测试 ： 推广数据
 路由 : prodata/add
 类型 ： GET
*/


func ProdList( c *gin.Context ){
	var prod modes.ProData
	prod.Start ,_ = strconv.Atoi( c.PostForm("start") )											
	count, _ := strconv.Atoi( c.PostForm("count") )
	page, _  := strconv.Atoi( c.PostForm("page") )

	list, err := prod.IdList( count, page - 1 )
	if nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "list":list})
}

func ProdSet( c *gin.Context ){
	var prod modes.ProData
	prod.Id ,_  = strconv.ParseInt(c.Query("id"),10,64) 			//必传
	prod.ClickNum ,_ = strconv.ParseInt(c.Query("clicknum"),10,64)  //必传
	prod.ShowNum ,_ = strconv.ParseInt(c.Query("shownum"),10,64) 	//必传
	prod.Submit ,_  = strconv.ParseInt(c.Query("submit"),10,64) 	//必传
	prod.Start ,_  = strconv.Atoi( c.Query("start") )				//必传
	fmt.Println( prod )
	if _, err := prod.IdSet("click_num,show_num,submit,start"); nil !=err {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg": err.Error() })
		return 
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"修改成功"})
}

