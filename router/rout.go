package router

import (
	//"fmt"
	"time"
	"../api/proway"
	"../api/login"
	"../api/chann"
	"../api/pro_url"
	"../api/impfile"
	"../api/prodata"
	"../api/spread"
	"../api/chanat"
	"../api/methods"
	"../api/prothe"
	"../api/account"
	"../api/proexp"
	"../api/extension"
	"../router/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var routs *gin.Engine

func login_rout( rout_name string ){
	rout := routs.Group( rout_name )
	rout.GET( "login", login.UserLogin)					//用户登陆
	rout.GET( "add",chanat.ChanatAdd)					//首咨注册添加
	rout.POST("callback_head",chanat.CallBackHead)
	rout.POST("bd_add",account.BaiDuLook)				//采集百度数据
	rout.POST("sll_add",account.SllGather)				//采集360数据
	rout.POST("sg_add",account.SgGather)					//采集搜狗数据
	rout.POST("sm_add",account.SmGather)					//采集神马数据
    rout.GET("download",impfile.HanderDownload) 			//下载
	rout.POST( "gdtadd",chanat.ChanatAddGdt)				//首咨注册添加	推广药师
}

/*
 * 描述：考研首资入库
 ******************************************************************/
func extension_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.GET( "extension_add",extension.ChantAdd)			//首咨添加
	rout.POST( "extension_list",extension.ChantList)			//首咨列表
}
/*
 * 描述：广点通健康路由
 ******************/
func spread_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "healthy01",spread.SpreadHealthy)			//首咨注册添加-广点通推广健康
	rout.POST( "healthy02",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy03",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy04",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy05",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy06",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy07",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy08",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy09",spread.SpreadHealthy)			//首咨注册添加
	rout.POST( "healthy10",spread.SpreadHealthy)			//首咨注册添加
}
/*
 * 描述：广点通法考路由
 ******************/
func method_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "method01",spread.SpreadMethod)			//首咨注册添加-广点通推广法考
	rout.POST( "method02",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method03",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method04",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method05",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method06",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method07",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method08",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method09",spread.SpreadMethod)			//首咨注册添加
	rout.POST( "method10",spread.SpreadMethod)			//首咨注册添加
}
/*
 * 描述：广点通消防路由
 ******************/
func fire_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "fire01",spread.SpreadFire)			//首咨注册添加-广点通消防推广
	rout.POST( "fire02",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire03",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire04",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire05",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire06",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire07",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire08",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire09",spread.SpreadFire)			//首咨注册添加
	rout.POST( "fire10",spread.SpreadFire)			//首咨注册添加
}

/*
 * 描述：广点通二建路由
 ******************/
func second_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "second01",spread.SpreadSecond)			//首咨注册添加-广点通二建
	rout.POST( "second02",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second03",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second04",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second05",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second06",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second07",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second08",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second09",spread.SpreadSecond)			//首咨注册添加
	rout.POST( "second10",spread.SpreadSecond)			//首咨注册添加
}

/*
 * 描述：广点通注册会计路由
 ******************/
func notes_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "notes01",spread.SpreadNotes)			//首咨注册添加-广点通注册会计
	rout.POST( "notes02",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes03",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes04",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes05",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes06",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes07",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes08",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes09",spread.SpreadNotes)			//首咨注册添加
	rout.POST( "notes10",spread.SpreadNotes)			//首咨注册添加
}
/*
 * 描述：广点通中级会计路由
 ******************/
func centre_rout(rout_name string)  {
	rout := routs.Group( rout_name )
	rout.POST( "centre01",spread.SpreadCentre)			//首咨注册添加-广点通中级会计
	rout.POST( "centre02",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre03",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre04",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre05",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre06",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre07",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre08",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre09",spread.SpreadCentre)			//首咨注册添加
	rout.POST( "centre10",spread.SpreadCentre)			//首咨注册添加
}

/*
 * 描述：渠道路由
 ******************/
func chann_rout( rout_name string)  {
	//chann
	rout :=routs.Group( rout_name )
	rout.Use( middleware.JWTAuth())
	rout.GET( "add",	chann.ChannAdd )				// 渠道新增
	rout.GET( "set",	chann.ChannSet)					// 修改渠道名称和状态
	rout.POST( "s_list",chann.ChannList)				// 渠道列表

	rout.GET( "wayadd",	proway.ProwayAdd ) 			// 推广方式新增
	rout.GET( "wayset",	proway.ProwaySet)			// 修改推广方式名称和状态
	rout.POST( "way_list",proway.ProwayList)			// 推广方式列表

	rout.GET( "metadd",	methods.MethodsAdd) 		// 营销方式新增
	rout.GET( "metset",	methods.MethodsSet)			// 修改营销方式名称和状态
	rout.POST( "met_list",methods.MethodsList)		// 营销方式列表
	//chann
	rout.GET( "theadd",	prothe.ProtheAdd) 			// 推广端新增
	rout.GET( "theset",	prothe.ProtheSet)			// 修改推广端名称和状态
	rout.POST( "the_list",prothe.ProtheList)			// 推广端列表
	
	rout.POST( "qd_list",chann.DisPlay)				//渠道条件显示
	rout.POST( "tg_list",proway.DisPlay)				//推广方式条件显示
	rout.POST( "yx_list",methods.DisPlay)			//营销方式条件显示
	rout.POST( "tgd_list",prothe.DisPlay)			//推广端条件显示
}

/*
*描述：url管理
*******************/
func url_rout( rout_name string )  {
	rout :=routs.Group(rout_name)
	rout.Use(middleware.JWTAuth())
	rout.POST("add",	pro_url.ProIdAdd)			//url添加
	rout.GET( "set" ,	pro_url.ProIdSet)		//url修改
	rout.POST( "p_list",pro_url.ProjectList)		//项目加项目id
	rout.POST( "lookup",pro_url.ProIdLookup)		//url管理列表加搜索
	rout.POST( "personal",pro_url.UserInfo)		//个人信息
	rout.GET("pass_set",pro_url.PassInfoSet)		//修改个人密码

	//推广账户
	rout.GET( "accadd" ,account.AccountAdd)		//推广账户的添加 
	rout.GET( "accset",account.AccountSet)		//推广账户修改
	rout.POST("acclist",account.AccountList)		//推广账户修改
}

/*
*描述：推广数据导入
*******************/
func impfile_rout( rout_name string )  {
	rout :=routs.Group(rout_name)
	rout.Use(middleware.JWTAuth())
	rout.POST("import",impfile.ImportFile)				//推广数据文件导入
	rout.POST("analysisfile",impfile.AnalysisFile)		//推广数据文件解析
	rout.POST("export_chant",impfile.Filess)				//首咨导出
}

/*
 *描述：推广数据统计
 *
 ********************/
 func pro_rout( rout_name string ) {
	rout := routs.Group(rout_name)
	rout.Use(middleware.JWTAuth())
	rout.GET("first",chanat.First)					//工作台数据 （实时）
	rout.POST("imadd",chanat.ImportAdd)				//导入数据添加		//pro_data
	rout.POST("imlist",chanat.ImportList)			//导入数据列表
	rout.GET( "set",prodata.ProdSet )				// 推广统计修改
	rout.POST("u_list",prodata.ProdList)				//推广统计列表显示
	rout.POST("lookup",chanat.ChanList)				//推广首咨列表
	rout.GET("del",chanat.ChanatDel)					//首咨删除
 }
/*
 *描述 ：推广消费
 *
 ****************/	//exp
func proexp_rout( rout_name string ) {
	rout := routs.Group(rout_name)
	rout.Use(middleware.JWTAuth())
	rout.POST("lookup",	proexp.ExpList)	//推广消费列表
	rout.GET("time",chanat.Time)			//工作台数据对比
	rout.GET("cost",chanat.RealTime)		//实时的花费
	rout.GET("click",chanat.ClickTime)	//实时的点击数
	rout.GET("sm",chanat.SeMatest)		//测试神马
	rout.GET("first",chanat.FirstCost)	//根据渠道ＩＤ和时间戳来获取首咨数量   消费
	rout.GET("sll",chanat.SllInfo)		//测试360数据

}
/*
	描述：百度推广
 */
func baidu_rout(rout_name string) {
	rout := routs.Group(rout_name)
	rout.POST("bd01",spread.BaiduAdd)		//百度推广
	rout.POST("bd02",spread.BaiduAdd)		//百度推广
	rout.POST("bd03",spread.BaiduAdd)		//百度推广
	rout.POST("bd04",spread.BaiduAdd)		//百度推广
	rout.POST("bd05",spread.BaiduAdd)		//百度推广
	rout.POST("bd06",spread.BaiduAdd)		//百度推广
	rout.POST("bd07",spread.BaiduAdd)		//百度推广
	rout.POST("bd08",spread.BaiduAdd)		//百度推广
	rout.POST("bd09",spread.BaiduAdd)		//百度推广
	rout.POST("bd010",spread.BaiduAdd)		//百度推广
	rout.POST("bd011",spread.BaiduAdd)		//百度推广
	rout.POST("bd012",spread.BaiduAdd)		//百度推广
	rout.POST("bd013",spread.BaiduAdd)		//百度推广
	rout.POST("bd014",spread.BaidufkAdd)		//百度推广
	rout.POST("bd015",spread.BaidufkAdd)		//百度推广

}

func init() {
	routs = gin.Default()
	// routs = gin.New()
	Init()
	routs.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	proexp_rout("exp")
	impfile_rout("impfile")
	login_rout( "login" )
	chann_rout( "chann" )
	extension_rout( "extension" )
	url_rout("tg_url")
	pro_rout("pro_data")
	spread_rout("spread")
	method_rout("method")
	fire_rout("fire")
	second_rout("second")
	notes_rout("notes")
	centre_rout("centre")
	baidu_rout("baidu")
	routs.Run(Run.RunPort)

}
