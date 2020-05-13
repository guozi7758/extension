package extension

import (
	"../../db/modes"
	"../../lib"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 *测试 ： 推广数据
 *路由 : extension/extension_add
 *类型 ：GET
**************************************************************/

func  ChantAdd( c *gin.Context) {
	var chant modes.ChaNat
	if c.Query("phone") == ""{
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"参数错误"})
		return
	}
	chant.Phone ,_ = strconv.ParseInt(c.Query("phone"),10,64)
	if fag, err := chant.Get(); nil == err {
		if fag {
			c.JSON(http.StatusOK,gin.H{"err": 0, "msg":"报名成功,稍后会有老师联系您！"})
			return
		}
	}else{
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":err.Error()})
		return
	}
	var jude lib.PhoneArea
	phones := strconv.FormatInt(chant.Phone,10)
	accede, err := jude.Area(phones)
	if err !=nil {
		fmt.Println("手机号解析错误")
	}
	chant.Url_code  = c.Query("url_code")
	chant.City = accede
	chant.ProWay ,_ =strconv.ParseInt(c.Query("pro_way"),10,64)
	chant.ProThe ,_ =strconv.ParseInt(c.Query("pro_the"),10,64)
	chant.ChannelId ,_ =strconv.ParseInt(c.Query("channelid"),10,64)
	chant.Courses = c.Query("courses")
	chant.Education ,_ = strconv.ParseInt(c.Query("education"),10,64)
	chant.Major = c.Query("major")
	chant.Pro_id ,_ = strconv.ParseInt(c.Query("pro_id"),10,64)
	chant.Pro_name = c.Query("pro_name")
	if _, err := chant.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "报名成功,稍后会有老师联系您！"})
}