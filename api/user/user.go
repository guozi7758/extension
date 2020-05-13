package user

import (
//	"fmt"
//	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
//	"../../lib"
)

/*
 * 描述: 用户详情
 * 路由: user/info
 * 类型: GET
 *
 ***********************************************************************/
func UserInfo( c *gin.Context ){
	var user modes.UserInfo
	user.Phone = modes.GetPhone( c )
	if _, err := user.Get(); nil != err {
		c.JSON(http.StatusOK, gin.H{"error_code": 1,"msg": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"error_code": 0,"info": user })

}

