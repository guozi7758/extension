DROP TABLE IF EXISTS channel;
CREATE TABLE IF NOT EXISTS channel(
	id      int unsigned AUTO_INCREMENT PRIMARY KEY, -- id,
	name    varchar(24)   NOT NULL UNIQUE,		 -- 渠道名称,
	start   tinyint       NOT NULL DEFAULT 0,	 -- 渠道状态 0 启用  1 禁用
	at	int unsigned  NOT NULL,  		 -- 创建时间,	
	creator	varchar(24)   NOT null 			 -- 创建
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;	 -- 推广渠道;

DROP TABLE IF EXISTS proid;
CREATE TABLE IF NOT EXISTS proid(
	id        int unsigned AUTO_INCREMENT PRIMARY KEY ,-- 'id',
	ursine	  varchar(64)  NOT NULL, 			-- 'url后缀',
	urlone	  varchar(64)  NOT NULL,			-- "url域名"
	channelid int unsigned NOT NULL,             -- '渠 道 ID',
 	miso_id   int unsigned NOT NULL,             -- '项 目 ID',
	account   varchar(64)  NOT NULL DEFAULT '',  -- '推广账户',
	at        int unsigned NOT NULL  ,           -- '创建时间',
	methods   tinyint      NOT NULL DEFAULT 0 ,  -- '营销方式', -- 0 电销 1 微信
	pro_the   tinyint      NOT NULL DEFAULT 0 ,  -- '推 广 端', -- 0 PC 1 移动浏览器 2 小程序 3 APP
	pro_way   tinyint      NOT NULL DEFAULT 0 ,  -- '推广方式', -- 0 搜索 1 信息流 2 APP
	key_file  varchar(256) NOT NULL ,            -- '文 件URL',
	c_name    varchar(24) 	NOT NULL ,           -- '创建人',
	code 	  varchar(64)	NOT NULL ,			 -- 唯一哈希值
	start     tinyint       NOT NULL DEFAULT 0   -- '状   态'        /*--- 0 启用  1 禁用*/ 
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=1 ; -- ='URL管理表';



DROP TABLE IF EXISTS cha_nat;
CREATE TABLE cha_nat (
	id        int unsigned    NOT NULL AUTO_INCREMENT,   -- '机会id',
	phone     bigint unsigned NOT NULL UNIQUE,           -- '手 机 号',
	channelid int unsigned    NOT NULL,                  -- '渠 道 ID',
	courses   tinyint         NOT NULL DEFAULT 0,        -- '1自然注册 2乐语注册 3其他',
-- 	s_type    tinyint         NOT NULL DEFAULT 0         -- '来源类型', -- 0 搜索推广 1 乐语推广 2 微信推广 3 自然流量
-- 	source_id int unsigned    NOT NULL DEFAULT 0         -- '来 源 ID',
    url_code   varchar(48)    NOT NULL,                  -- 'url'
    pro_way   tinyint         NOT NULL DEFAULT 0,        -- '推广方式', -- 1 搜索 2 信息流 3 APP  4乐语推广 5自然注册
    pro_the   tinyint         NOT NULL DEFAULT 0,        -- '推 广 端', -- 1 PC 2 移动浏览器 3 小程序 4 APP
	education tinyint         NOT NULL DEFAULT 0,        -- '学    历', -- 1 无学历, 2 中专 3 大专， 4 本科 5 研究生
	major	  varchar(100)	  NOT NULL DEFAULT 0, 	     -- '备注',
	city      int             NOT NULL DEFAULT 0,        -- 'AD  CODE',
	uid       int unsigned    NOT NULL DEFAULT 0,        -- '指派坐席',
	uname     varchar(35)     NOT NULL DEFAULT '',       -- '坐席姓名',
	at        int unsigned    NOT NULL,                  -- '创建时间',
	pro_id	  int unsigned	  NOT NULL, 				    -- '项 目 ID',
	pro_name  varchar(24)	  NOT NULL, 				    -- '项目名称',
	start	  int unsigned    NOT NULL DEFAULT 0,	    -- '状态( 0 未分配 1 已分配 2 已成交)',
	PRIMARY KEY(id),
	INDEX cha_nat( pro_id,uid )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ;-- ='推广机会属性表';

DROP TABLE IF EXISTS file_data;
CREATE TABLE file_data (
  id		int unsigned	NOT NULL AUTO_INCREMENT, -- 'excel表ID',
  miso_id	int unsigned	NOT NULL,   		-- '咨询项目id',
  state		tinyint		NOT NULL DEFAULT 0, 	-- '0 没解析 1解析',
  at		int unsigned	NOT NULL, 		-- '入库时间',
  key_file	varchar(128)	NOT NULL DEFAULT '', 	-- '文件url',
  uid       	int unsigned	NOT NULL,                -- '员 工 ID',
  PRIMARY KEY ( id )
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ;-- ='推广导入excel数据表';


/*推广端*/
DROP TABLE IF EXISTS prothe; 
CREATE TABLE IF NOT EXISTS prothe(
	id      int  unsigned AUTO_INCREMENT PRIMARY KEY, -- 'id',
	name    varchar(24)   NOT NULL UNIQUE,		 	 -- '推广端',
	start   tinyint       NOT NULL DEFAULT 0,		 -- '推广端状态',  /*--- 0 启用  1 禁用*/
	at		int unsigned  NOT NULL,  			 	 -- '创建时间',	
	creator	varchar(24)   NOT NULL 				 	 -- '创建人'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;-- ='推广端';
/*营销方式*/
DROP TABLE IF EXISTS methods; 
CREATE TABLE IF NOT EXISTS methods(
	id      int  unsigned AUTO_INCREMENT PRIMARY KEY, -- 'id',
	name    varchar(24)   NOT NULL UNIQUE,			 -- '营销方式',
	start   tinyint       NOT NULL DEFAULT 0,		 -- '营销方式状态',  /*--- 0 启用  1 禁用*/
	at		int unsigned  NOT NULL,  			 	 -- '创建时间',	
	creator	varchar(24)   NOT NULL 					 -- '创建人'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;-- ='营销方式';
/*推广方式*/
DROP TABLE IF EXISTS proway; 
CREATE TABLE IF NOT EXISTS proway(
	id      int  unsigned AUTO_INCREMENT PRIMARY KEY, -- 'id',
	name    varchar(24)   NOT NULL UNIQUE,			 -- '推广方式',
	start   tinyint       NOT NULL DEFAULT 0,		 -- '推广方式状态',  /*--- 0 启用  1 禁用*/
	at		int unsigned  NOT NULL,  				 -- '创建时间',	
	creator	varchar(24)   NOT NULL 					 -- '创建人'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1; -- ='推广方式';

/*推广账户*/
DROP  TABLE IF EXISTS account;
CREATE TABLE  IF NOT EXISTS  account(
    id          int  unsigned AUTO_INCREMENT,                  -- 'id',
    uid        	int  unsigned   NOT NULL,                      -- '用户id',
	project		int 			NOT NULL DEFAULT 0,			  -- '项目id',
  	name     	varchar(70)  	NOT NULL,            	  	  -- '推广账户',
    channelid  	int   			NOT NULL DEFAULT 0,            -- '渠 道id',
	proway		int 			NOT	NULL DEFAULT 0,			  -- '推广方式',
	prothe		int 			NOT NULL DEFAULT 0,			  -- '推广端',
    url     	varchar(255)    NOT NULL,                  	  -- 'url',
    start       tinyint         NOT NULL DEFAULT 0,		      -- '状态',  /*--- 0 启用  1 禁用*/
	at			int	unsigned	NOT NULL, 					  -- '时间',
    PRIMARY KEY ( id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ; -- ='推广账户';

/*数据推广数据存放*/
DROP  TABLE IF EXISTS sics_data;
CREATE TABLE  IF NOT EXISTS  sics_data(
    id          int  unsigned AUTO_INCREMENT,                  -- 'id',
	name		varchar(64),				 				 	  -- '账户名称',
	at			int	unsigned	NOT NULL, 					  -- '插入时间',
	start_at	varchar(24)		NOT NULL,					  -- '统计时间'	
	impression 	int	unsigned	NOT NULL DEFAULT 0,			      -- '展示次数',
	click 		int unsigned		NOT NULL DEFAULT 0,			      -- '点击数量',
	cost		int unsigned		NOT NULL DEFAULT 0,			      -- '所花费用',
	ctr 		int	unsigned	NOT NULL DEFAULT 0,			          -- '点击率',
	cpc			int unsigned		NOT	NULL DEFAULT 0,			      -- '平均点击价格',
	cpm			int unsigned		NOT	NULL DEFAULT 0,			      -- '单元出价',
	start 		tinyint 		NOT NULL DEFAULT 0, 			  -- '类别', /*1.百度  2.360  3.搜狗 4.神马*/
	pid         int unsigned    NOT NULL,                       -- '项目id'   -- 追加
    PRIMARY KEY ( id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ; -- ='百度数据推广数据存放';-



DROP  TABLE IF EXISTS import;
CREATE TABLE  IF NOT EXISTS  import(
    id          int  	unsigned 	AUTO_INCREMENT,                 -- 'id',
	pid			int 	unsigned	NOT NULL,						-- 项目id
	c_name		varchar(24)			NOT NULL,						-- 操作人
	at 			int		unsigned 	NOT NULL,						-- 插入时间 
	start 		tinyint 			NOT NULL	DEFAULT 0,			-- 是否执行/*0未执行	1已执行*/
	im_count	int 	unsigned 	NOT NULL	DEFAULT 0,			-- 导入总条数
	wrong		int		unsigned 	NOT NULL	DEFAULT 0,			-- 错误条数
	actual		int 	unsigned 	NOT NULL 	DEFAULT 0,			-- 实际执行条数
	exec_at		int		unsigned 	NOT	NULL 	DEFAULT 0,			-- 执行时间 
	file_exec	varchar(256)		NOT NULL,						-- 文件
	code 		varchar(124)		NOT NULL	DEFAULT 0,						-- hashcode 
    PRIMARY KEY ( id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ; 								-- '导入execl';





/*
 * 描述：项目表
 *
 *	id =  00     00       000     
 *	==================================================================================
 *           大区  事业部     项目
 *
 *	state 0 启用  1 禁用
 *		
 * 报考条件: 报考条件为HTML文件，ID是表名
 * 
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS project;
CREATE TABLE IF NOT EXISTS project (
	id    int unsigned	NOT NULL,		-- 项 目 id
	name  varchar(24) 	NOT NULL UNIQUE,     	-- 项目名称
	start tinyint     	NOT NULL DEFAULT 0,  	-- 状    态
	PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO project ( id,name,start )VALUES( 1000000, '郑州', 0);
INSERT INTO project ( id,name,start )VALUES( 1001000, '医学事业', 0);
INSERT INTO project ( id,name,start )VALUES( 1001001, '职业药师', 0);
INSERT INTO project ( id,name,start )VALUES( 1001002, '健康管理', 0);
INSERT INTO project ( id,name,start )VALUES( 1002000, '财会事业', 0);
INSERT INTO project ( id,name,start )VALUES( 1002001, '法考', 0);
INSERT INTO project ( id,name,start )VALUES( 1003000, '建造事业', 0);
INSERT INTO project ( id,name,start )VALUES( 1003001, '一级消防', 0);
