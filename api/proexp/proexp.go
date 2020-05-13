package proexp

import (
	"../../db/modes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
*推广数据报表
*
**********************/
func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}
//推广数据报表 POST

func ExpList( c *gin.Context){
	var val modes.TgData
	var page modes.Pager
	BuildPager( c,&page)
	if c.PostForm("start_time") != "" {
			page.Add("at",">=","and",c.PostForm("start_time"))
	}
	if c.PostForm("p_id") != "" {
		page.Add("pid","=","and",c.PostForm("p_id"))
	}
	if c.PostForm("end_time") != ""{
			page.Add("at","<=","and",c.PostForm("end_time"))
	}
	if c.PostForm("channelid") != ""{
			page.Add("start","=","and",c.PostForm("channelid"))
	}
	if err := val.Lookups(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})
}
