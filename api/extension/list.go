package extension

import (
	"../../db/modes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}
/*
 *测试 ： 推广列表
 *路由 : extension/extension_list
 *类型 ：POST
**************************************************************/

func ChantList( c *gin.Context){
	var val modes.SetChanat
	var page modes.Pager
	BuildPager( c,&page)
	if c.PostForm("phone") != "" {
		page.Add( "phone", fmt.Sprintf("like '%s%%'", c.PostForm("phone") ), "and", "" )
	}
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
	if c.PostForm("pro_id") != ""{
		page.Add("pro_id","=","and",c.PostForm("pro_id"))
	}
	if c.PostForm("channelid") != ""{
		page.Add("channelid","=","and",c.PostForm("channelid"))
	}
	if c.PostForm("start") != "" {
		page.Add("start","=","and",c.PostForm("start"))
	}
	if err := val.ChanatLists(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})
}