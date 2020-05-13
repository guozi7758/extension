package login

import (
	"time"
//	"strconv"
	  "../../router/middleware"
        "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"net/http"
	"../../db/modes"
//	"../../lib"
)

/*
 * 描述: 创建用户TOKEN
 *
 ***********************************************************************/
func create_token( Id,GId,PId,RId  int64,Name,Phone string ) (string, error){
        j := &middleware.JWT{
                []byte("aldie33!@Dd$$%G&&#asd$^asd&*x%a%……334*"),
        }
        claims := middleware.CustomClaims{
                Id,   // 用户Id
		GId,   // 组ID
		PId,   // 项目ID
		RId,   // 角色ID
		Name,  // 姓名
		Phone, // 用户手机号
                jwt.StandardClaims{
                        NotBefore: time.Now().Unix(),                     // 签名生效时间
                        ExpiresAt: time.Now().Unix() + ( 1 * 86400 ),    // 过期时间 一天
                        Issuer:    "user",                                // 签名的发行者
                },
        }
        return j.CreateToken(claims)
}

/*
 * 描述: 用户登陆
 * 路由: login/login
 * 类型: GET
 *
 ***********************************************************************/
func UserLogin( c *gin.Context ){
        var user modes.Employee
        user.Phone = c.Query( "phone" )
        if err := c.Query("phone");err == ""{
        	c.JSON(http.StatusOK,gin.H{"err":2,"msg":"用戶名不能为空"})
        	return
		}
        if fage, err := user.Get(); nil == err {
		if !fage {
			c.JSON(http.StatusOK, gin.H{"err": 2,"msg": "此账户未授权给推广系统" })
			return
		}
        }else{
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
		return
	}
        if user.Pass !=  c.Query( "pass" ){
                c.JSON(http.StatusOK, gin.H{"err": 1,"msg": "密码不正确" })
		return
        }
	token,_ := create_token( user.Id, user.TeamId, user.ProjId, user.RoleId, user.Name, user.Phone)
        c.JSON(http.StatusOK, gin.H{"err": 0, "token": token ,"info": user })
}



