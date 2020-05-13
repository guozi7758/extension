package spread
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
	描述：广点通推广健康
*/
func SpreadHealthy(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
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
	channe.Pro_id = 1001002
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "健康管理师"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}

func RegexpPhone(phone string) string  {
	re, _ := regexp.Compile(`\d+`)
	//查找符合正则的第一个
	all := re.FindAll([]byte(phone),-1)
	var regphone string
	for _,item := range all {
		regphone = string(item)
	}
	return regphone
}
/*
	描述：广点通推广法考
*/
func SpreadMethod(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
	var val lib.ChanntGdt
	var channe modes.ChaNat
	err := c.BindJSON(&val)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"数据有误"})
		return
	}
	for _,row := range val.Data.Data{
		if strings.Contains(row.Label, "手机") == true{
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
	channe.Pro_id = 1002001
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "法律资格考试"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}
/*
	描述：广点通消防推广
 */
func SpreadFire(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
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
	channe.Pro_id = 1003001
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "一级注册消防"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}
/*
 描述：广点通二建
 */
func SpreadSecond(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
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
	channe.Pro_id = 1003002
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "二级建造师"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}

/*
	描述：广点通注册会计
*/

func SpreadNotes(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
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
	channe.Pro_id = 1002002
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "注册会计师"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}
/*
	描述：广点通中级会计
 */

func SpreadCentre(c *gin.Context)  {
	//url := fmt.Sprintln(c.Request.URL)
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
	channe.Pro_id = 1002003
	channe.ProWay = 6
	channe.ChannelId =18
	channe.ProThe = 6
	channe.Courses = fmt.Sprintln(val.Data.AccountID)
	channe.Education = 1
	urlcode := Split(fmt.Sprintln(val.Data.URL),"?")
	ucode :=fmt.Sprintln(urlcode[0])
	channe.Url_code  = ucode
	channe.Pro_name = "中级会计师"
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})
}

func  BaiduAdd(c *gin.Context)  {			//百度药师信息流
	url := fmt.Sprintln(c.Request.URL)
	var channe modes.ChaNat
	var bd lib.ChanntBaiDu
	err := c.BindJSON(&bd)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"数据有误"})
		return
	}
	for _,row := range bd.Data{
		if strings.Contains(row.Name,"手机") == true{
			phone := RegexpPhone(row.Value)
			channe.Phone,_ = strconv.ParseInt(phone,10,64)
			if fage,err := channe.Get();nil == err {
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
		channe.Major = fmt.Sprintf("%s--%s:%s", channe.Major,row.Name, row.Value )
	}
	channe.Pro_id = 1001001
	channe.ProWay = 6
	channe.ChannelId =3
	channe.ProThe = 6
	channe.Courses = "1"
	channe.Education = 1
	channe.Url_code  = url
	channe.Pro_name = "执业药师"
	fmt.Println(channe)
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})

}


func  BaidufkAdd(c *gin.Context)  {				//百度法考信息流
	url := fmt.Sprintln(c.Request.URL)
	var channe modes.ChaNat
	var bd lib.ChanntBaiDu
	err := c.BindJSON(&bd)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": 1, "msg":"数据有误"})
		return
	}
	for _,row := range bd.Data{
		if strings.Contains(row.Name,"手机") == true{
			phone := RegexpPhone(row.Value)
			channe.Phone,_ = strconv.ParseInt(phone,10,64)
			if fage,err := channe.Get();nil == err {
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
		channe.Major = fmt.Sprintf("%s--%s:%s", channe.Major,row.Name, row.Value )
	}
	channe.Pro_id = 1002001
	channe.ProWay = 6
	channe.ChannelId =3
	channe.ProThe = 6
	channe.Courses = "1"
	channe.Education = 1
	channe.Url_code  = url
	channe.Pro_name = "法律资格考试"
	fmt.Println(channe)
	if _, err := channe.Save(); nil != err {
		c.JSON(http.StatusOK,gin.H{"err": 2, "msg":"报名成功"})
		return
	}
	var inc cache.FirstInfo
	if  err := inc.In(); nil != err{
		fmt.Println("计数失败")
	}
	c.JSON(http.StatusOK,gin.H{"err":0, "msg":"报名成功,稍后会有老师联系你"})

}


