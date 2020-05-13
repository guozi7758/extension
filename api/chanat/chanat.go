package chanat

import (
	"../../db/cache"
	"../../db/modes"
	"../../lib"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	//"time"
)
func Split( str, char string )[]string{
	re := regexp.MustCompile( fmt.Sprintf("[%s]", char ) )
	split := re.Split( str, -1)
	set := []string{}
	for i := range split {
		set = append(set, split[i])
	}
	return set
}
/*
 测试 ： 推广数据广点通
 路由 : login/ChanatAddGdt
 类型 ：GET
*/
func ChanatAddGdt(c *gin.Context)  {
	var val lib.ChanntGdt
	var channe modes.ChaNat
	err := c.BindJSON(&val)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"数据有误"})
		return
	}
	for _,row := range val.Data.Data{
			if strings.Contains(row.Label, "手机") == true {
				channe.Phone, _ = strconv.ParseInt(row.Value,10,64)
				if fage, err := channe.Get(); nil == err {
					if fage{
						c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"报名成功"})
						return
					}
				}
				var juhe lib.PhoneArea
				phoness := strconv.FormatInt(channe.Phone,10)
				adcode, err :=juhe.Area(phoness)
				if err !=nil {
					fmt.Println("手机号解析错误")
				}
				channe.City = adcode
				continue
			}
			channe.Major = fmt.Sprintf("%s--%s:%s", channe.Major,row.Label, row.Value )
	}
	channe.Pro_id = 1001001
	channe.ProWay = 6
	channe.ChannelId = 18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "执业药师"
	//fmt.Println(channe.Url_code)
	//return
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":err.Error()})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}
/*
 测试 ： 推广数据
 路由 : login/add
 类型 ：GET
*/

func  ChanatAdd( c *gin.Context) {
	var chanat modes.ChaNat
	chanat.Phone ,_ = strconv.ParseInt(c.Query("phone"),10,64)
	if fage, err := chanat.Get(); nil == err {
		if fage{
			c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"报名成功"})

			return
		}
	}
	var juhe lib.PhoneArea
	phoness := strconv.FormatInt(chanat.Phone,10)
	adcode, err :=juhe.Area(phoness)
	if err !=nil {
		fmt.Println("手机号解析错误")
	}
	chanat.Url_code  = c.Query("url_code")
	chanat.City = adcode
	chanat.ProWay ,_ =strconv.ParseInt(c.Query("pro_way"),10,64)
	chanat.ProThe ,_ =strconv.ParseInt(c.Query("pro_the"),10,64)
	chanat.ChannelId ,_ =strconv.ParseInt(c.Query("channelid"),10,64)
	chanat.Courses = c.Query("courses")
	chanat.Education ,_ = strconv.ParseInt(c.Query("education"),10,64)
	chanat.Major = c.Query("major")
	chanat.Pro_id ,_ = strconv.ParseInt(c.Query("pro_id"),10,64)
	chanat.Pro_name = c.Query("pro_name")

	if _, err := chanat.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":err.Error()})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "报名成功"})
}
//首咨删除
func ChanatDel(c *gin.Context ){
	var val modes.ChaNat
	var err error
	val.Id,_ = strconv.ParseInt(c.Query("id"),10,64)
	if c.Query("id") == "" {
		c.JSON(http.StatusOK, gin.H{"err": 1, "msg":"参数错误"})
		return
	}
	if _, err :=val.Get();err !=nil{
		c.JSON(http.StatusOK,gin.H{"err":2, "msg": err.Error() })
		return
	}
	start := val.Start
	if start == 1 {
		c.JSON(http.StatusOK, gin.H{"err": 3,"msg":"首咨已分配不可删除" })
		return
	}
	if _, err = val.Del(); nil != err{
		c.JSON(http.StatusOK, gin.H{"err":4,"msg": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "首咨删除成功"})

}
func BuildPager( c *gin.Context, page *modes.Pager ){
	page.Count,_	= strconv.Atoi( c.PostForm("count") ) // 单页数量
	page.Page,_	= strconv.Atoi( c.PostForm("page") )  // 页码
	page.Coll,_     = strconv.Atoi( c.PostForm("coll") )  // 排序字段
	page.Rules,_	= strconv.Atoi( c.PostForm("rule") )  // 排序规则
}

//推广列表
func ChanList( c *gin.Context){
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
	fmt.Println(page)

	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})
}

/*
 *描述:导入数据添加 
 *类型：POST
 *******************************************************************/
 func ImportAdd(c *gin.Context){
	var val modes.Import
	jwt_user := modes.JwtUser(c)
	val.Code = strconv.Itoa(0)
	val.CName =jwt_user.Name			
	val.Pid ,_ = strconv.ParseInt(c.PostForm("pid"),10,64)
	val.FileExec = c.PostForm("file")
	fmt.Println("导入数据信息",val)
	if _, err := val.Save();nil != err {
		c.JSON(http.StatusOK,gin.H{"err":1 , "msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0,"msg": "ok"})
 }

 /*
 *描述:导入数据列表
 *类型：POST	
 *******************************************************************/
 func ImportList( c *gin.Context){
	var val modes.Import
	var page modes.Pager
	BuildPager( c,&page)
	if c.PostForm("start_time") != "" {
			page.Add("at",">","and",c.PostForm("start_time"))
	}
	if c.PostForm("end_time") != ""{
			page.Add("at","<","and",c.PostForm("end_time"))
	}
	if c.PostForm("pid") != ""{
		page.Add("pid","=","and",c.PostForm("pid"))
	}
	if err := val.List(&page);err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1 ,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0 ,"data" :page})
}

//时实工作台首咨数据
func First(c *gin.Context) {
	var data cache.AdminData
	if err := data.Get();nil!=err {
		c.JSON(http.StatusOK,gin.H{"err":1,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "info": data })
}
//工作台数据对比
func Time(c *gin.Context) {
	var val modes.Gather
	s_time ,_ :=strconv.ParseInt(c.Query("s_time"),10,64)
	e_time ,_ :=strconv.ParseInt(c.Query("e_time"),10,64)
	list,err := val.Get(s_time,e_time)
	if nil != err {
		c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
		return
	}
	c.JSON(http.StatusOK,gin.H{"err": 0, "info":list})
}

//工作台实时数据(cost)
//strconv.ParseInt(string, 10, 64)
func RealTime(c *gin.Context){
	//var redis cache.SMInfo
	//redis.Get()
	//var bai lib.FieldsData
	//var soso lib.AutoGenerated
	//var bcosts int64
	//var sllcost int64
	//timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	//times := time.Unix(GetHourZero(time.Now().Unix()), 0).Format(timeLayout)
	//baiducost,_ := strconv.ParseInt( bai.Cost_Time(), 10, 64)					//实时百度消费
	//sosocost,_ := strconv.ParseInt( soso.SoSo_Coust(times), 10, 64)			//实时搜狗消费
	////str,_ := lib.SanLiuZero.GetToken()
	////sllcost,_ := strconv.ParseInt(lib.SanLiuZero.SllTime(str), 10, 64)			//实时360消费
	//for i, _ := range lib.SanLiuZeros {
	//	bcosts,_ = lib.SanLiuZeros[i].SllTime()
	//	sllcost += bcosts
	//}
	//cost := 0 + sosocost + sllcost + 0										//sum实时数据
	//fmt.Println("soso","360",sosocost,sllcost)
	c.JSON(http.StatusOK,gin.H{"err": 0, "cost":0})
}
//工作台实时数据(click)

func ClickTime(c *gin.Context){
	//var redis cache.SMInfo
	//redis.Get()
	//println(redis.Click)
	//var bai lib.FieldsData
	//var soso lib.AutoGenerated
	//var bclick int64
	//var sllclick int64
	//timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	//times := time.Unix(GetHourZero(time.Now().Unix()), 0).Format(timeLayout)
	//baiduclick,_ := strconv.ParseInt( bai.Click_Time(), 10, 64) 			//百度实时click
	//sosoclick,_ := strconv.ParseInt( soso.SoSo_Click(times), 10, 64)				//搜狗实时click
	////str,_ := lib.SanLiuZero.GetToken()
	////sllclick,_ := strconv.ParseInt(lib.SanLiuZero.Click_SllTime(str), 10, 64) 		//实时360click
	//for i, _ := range lib.SanLiuZeros {
	//	_,bclick = lib.SanLiuZeros[i].SllTime()
	//	sllclick += bclick
	//}
	//click := 0 + sosoclick +sllclick + 0											//sum实时click
	//fmt.Println("soso","360",sosoclick,sllclick )
	//c.JSON(http.StatusOK,gin.H{"err": 0, "click":click})

	c.JSON(http.StatusOK,gin.H{"err": 0, "click":0})
}
func GetHourZero(checkTime int64 ) int64 {
	return checkTime - (checkTime % 3600)
}

func SeMatest(c *gin.Context)  {
	tack_id := lib.SmZero.GetTaskId()
	file_data := lib.SmZero.GetFileId( tack_id )
	smclick,smcost := lib.SmZero.GetFileCilck( file_data )
	fmt.Println(smclick,smcost)


}
//根据渠道ＩＤ和时间戳来获取首咨数量和消费
func FirstCost(c *gin.Context) {
	var val modes.BaiduFirst
	var su modes.BaiduSum
	s_time ,_ :=strconv.ParseInt(c.Query("s_time"),10,64)
	e_time ,_ :=strconv.ParseInt(c.Query("e_time"),10,64)
	channelid ,_ := strconv.ParseInt(c.Query("channelid"),10,64)
	if channelid == 0{
		list,err := su.BaiduSum(s_time,e_time)
		if nil != err {
			c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
			return
		}
		c.JSON(http.StatusOK,gin.H{"err": 0, "info":list})
	}else{
		list,err := val.BaiduCountone(s_time,e_time,channelid)
		if nil != err {
			c.JSON(http.StatusOK, gin.H{"err": 1,"msg": err.Error() })
			return
		}
		c.JSON(http.StatusOK,gin.H{"err": 0, "info":list})
	}
}
	//360数据
func SllInfo(c *gin.Context)  {
	for i, _ := range lib.SanLiuZero {
		err := lib.SanLiuZero[i].SllTime()
		if err != nil {
			fmt.Println("采集数据失败",err.Error())
		}
	}

	c.JSON(http.StatusOK,gin.H{"err": 0, "info":"Ok"})
}

func CallBackHead(c * gin.Context)  {
	fmt.Println("接口回传数据",c.Request)
	fmt.Println("测试今日头条数据回调")
}

