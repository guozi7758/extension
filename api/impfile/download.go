package impfile

import (
	//"encoding/json"
	//"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/tools/godoc/util"
	//"io/ioutil"
	//"net/http"
	//"strconv"
	//"strings"
	_ "strings"
)

func HanderDownload(c *gin.Context) {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A2", "我要下载一个excel文件")
	xlsx.SetCellValue("Sheet1", "A1", "有没有看到我帅气的脸庞")

	//保存文件方式
	//_ = xlsx.SaveAs("./aaa.xlsx")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = xlsx.Write(c.Writer)

}