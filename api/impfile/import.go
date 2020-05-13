package impfile

import (
	"../../db/modes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
	"../../lib"
)

/*
 测试 ： url导入推广数据文件
 路由 :  impfile/ImportFile
 类型 ： POST
*/
func ImportFile( c *gin.Context ){
	var imp_file modes.FileData
	jwt_user := modes.JwtUser(c)
	imp_file.UId = jwt_user.Id
	imp_file.ProId ,_ = strconv.ParseInt(c.PostForm("impo_id"),10,64)
	file, err := c.FormFile("imp_file")
	if err != nil{
		fmt.Println( "url error:", err.Error() )
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg": "地址有误" })
		return
	}
	imp_file.KeyFile = fmt.Sprintf( "./upload/%d_%s", time.Now().Unix(),file.Filename )
	if err := c.SaveUploadedFile(file,imp_file.KeyFile); err !=nil{
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
		return
	}
	if _, err := imp_file.Save(); nil != err {
		fmt.Println( "Chanquery  err :", err.Error() )
		os.Remove(imp_file.KeyFile)
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"入库失败" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "文件导入成功"})
}

/*
 测试 ： url解析推广数据文件
 路由 :  impfile/ImportFile
 类型 ： POST
*/
func AnalysisFile( c *gin.Context ){
	var filedata lib.FileData
	if _, err :=os.Stat(c.PostForm("file_url")); err !=nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":fmt.Sprintf("%s:文件读取失败",c.PostForm("file_url")) })
		return
	}
	list, err :=filedata.ParsingFile(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg":"文件读取失败" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "文件解析成功","data":list})
}

func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}

func Filess(c *gin.Context){
	var val modes.ChaNat
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
	if path,err := val.ToExcelFile(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}else{
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%d.xlsx",time.Now().Day()))
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(path)
		os.Remove(path)
	}
}
