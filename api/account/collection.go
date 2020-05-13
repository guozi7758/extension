package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"../../lib"
	// "fmt"
)

/*
 描述 ： 采集每天360数据
 路由 ：login/sll_add
 类型 ： GET
 ******************************************************************************/
func SllGather( c *gin.Context ){
	//str,_ := lib.SanLiuZero.GetToken()
	//lib.SanLiuZero.SllTime(str)
}

/*
 描述 ： 采集每天搜狗数据
 路由 ：login/sll_add
 类型 ： GET
 ******************************************************************************/
func SgGather( c *gin.Context ){
	//var sg lib.Envelope
	//err := sg.SgTime()
	//if err !=nil{
	//	c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"采集失败" })
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"err": 0,"msg":"ok" })

}
/*
 描述 ： 采集每天神马数据
 路由 ：login/sm_add
 类型 ： GET
 ******************************************************************************/
func SmGather( c *gin.Context ){
	//var sm lib.SmServer
	//err := sm.GetTaskId()
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"err": 1,"msg":err.Error() })
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg":"ok" })
}

func BaiDuLook( c *gin.Context)  {
	var bai lib.FieldsData
	baiduclick,_ := strconv.ParseInt( bai.Click_Time(), 10, 64)
	fmt.Println(baiduclick)
}
