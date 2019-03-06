package constvar

const (
	DefaultLimit   = 50
	BanIpLimit     = 10 // 现在IP错误密码次数
	UserTokenLimit = 10 // 用户token数量限制
	SmsCodeLimit   = 20 // 限制每天每个IP、用户发送短信数量
	StatusOk       = 1  //webscoket 正确响应code
	StatusErr      = 2  //webscoket 错误响应code
	AckLimit       = 3  //消息推送最大重试次数

	// 配置图片URL前缀
	PicPathPrefix = "http://tu.jstucdn.com/ftp/"

	/**** Websocket MType 相关配置 start ******/
	DownloadInitRequestMtype            = 1    //websocket下载模块初始化登录请求
	DownloadInitResponeMtype            = 1001 //websocket下载模块初始化登录响应
	AppInitRequestMtype                 = 2    //websocketAPP初始化登录请求
	AppInitResponeMtype                 = 1002 //websocketAPP初始化登录响应
	MtypeDownloadHeartRequest           = 201  //websocket下载模块心跳请求
	DownloadHeartResponeMtype           = 1201 //websocket下载模块心跳响应
	AppHeartRequestMtype                = 202  //websocket APP心跳请求
	AppHeartResponeMtype                = 1202 //websocket App心跳响应
	MtypePushTaskTodDownloadRequest     = 203  //websocket推送下载中队列至下载模块请求
	MtypePushTaskTodDownloadRespone     = 1203 //websocket推送下载中队列至下载模块响应
	MtypeDelTaskTodDownloadRequest      = 204  //websocket删除任务推送至下载模块请求
	MtypeDelTaskTodDownloadRespone      = 1204 //websocket删除任务推送至下载模块响应
	MtypeAppGetJoinDeviceInfoRequest    = 210  //websocketAPP 获取设备详情请求
	MtypeAppGetJoinDeviceInfoResponse   = 1210 //websocketAPP 获取设备详情响应
	MtypeAppGetRouterSpeedInfoRequest   = 101  //websocketAPP 获取路由器上下行速度信息请求
	MtypeAppGetRouterSpeedInfoResponse  = 1101 //websocketAPP 获取路由器上下行速度信息响应
	MtypeAppGetTaskListRequest          = 103  //websocketAPP 获取下载任务列表请求
	MtypeAppGetTaskListResponse         = 1103 //websocketAPP 获取下载任务列表响应
	MtypeRouterStorageScanRequest       = 102  //websocket  2.3.2	平台下发扫描目录命令给路由器请求
	MtypeRouterStorageScanResponse      = 1102 //websocketAPP 2.3.2	平台下发扫描目录命令给路由器响应
	MtypeRouterDelPathRequest           = 212  //websocket  2.3.2	平台下发扫描目录命令给路由器请求
	MtypeRouterDelPathResponse          = 1212 //websocketAPP 2.3.2	平台下发扫描目录命令给路由器响应
	MTypeRouterSetSpeedLimitRequest     = 209  //websocket  4.5	设置下载上传速度限制请求
	MTypeRouterSetSpeedLimitResponse    = 1209 //websocketAPP 4.5	设置下载上传速度限制响应
	MTypeRouterStorageOutNoticeRequest  = 205  //websocketAPP 4.1	外接设备插拔通知请求
	MTypeRouterStorageOutNoticeResponse = 1205  //websocketAPP 4.1	外接设备插拔通知响应
	MTypeGetStorageParRequest           = 105  //websocketAPP 7.1	获取外接存储分区信息请求
	MTypeGetStorageParResponse          = 1105 //websocketAPP 7.1	获取外接存储分区信息响应
	MTypeTaskDealNoticeRequest          = 207  //websocketAPP 4.3	任务通知请求
	MTypeAPPNoticeRequest               = 208  //websocketAPP 2.12	app通用通知请求
	MTypeAPPNoticeResponse              = 1208 //websocketAPP 2.12	app通用通知响应
	MTypeGetSearchLocalResponse         = 1104 //websocketAPP 3.2	下载模块搜索本地资源文件响应
	MTypeGetSearchLocalRequest          = 104  //websocketAPP 3.2	下载模块搜索本地资源文件请求
	MTypeSetPartitionResponse           = 1211 //websocketAPP 8.1	设置分区响应
	MTypeSetPartitionRequest            = 211  //websocketAPP 8.1	设置分区请求
	MTypeAppStorageNoticeRequest        = 206  //websocketAPP  4.2 外接存储拔插通知
	/**** Websocket MType 相关配置 end ******/


	/**** time 相关配置 start ******/
	TimeFormatD       = "2006-01-02"
	TimeFormatS       = "2006-01-02 15:04:05"
	TimeFormatSecond  = 1
	TimeFormatMin     = 60
	TimeFormatHour    = 60 * 60
	TimeFormatCST     = 60 * 60 * 8
	TimeFormatDay     = 60 * 60 * 24
	TimeFormatWeek    = 60 * 60 * 24 * 7
	TimeFormatMonth   = 60 * 60 * 24 * 30
	TimeFormatYear    = 60 * 60 * 24 * 365
	TimeFormatTenYear = 60 * 60 * 24 * 365 * 10
	/**** time 相关配置 end ******/



	// redis 相关配置

	/**** 前缀 start ******/
	RouterMqRedisPrefix = "mq:"
	RouterRedisPrefix   = "router:"
	UserRedisPrefix     = "user:"
	/**** 前缀 end ******/


	/**** 公用 start ******/
	RedisUserBanIP    = "public:ban:ip:"
	RedisSmsBanIP     = "public:sms:ban:ip:"
	RedisSmsSend      = "public:sms:send:ip:"
	RedisEmailSend    = "public:email:send:ip:"
	RedisUniqueCode   = "public:unique:code:"
	RedisUserLoginErr = "public:login:password:"
	RedisGroupInfo    = "public:group:info:"
	RedisAnnounce     = "public:announce:"
	/**** 公用 end   ******/


	/**** 用户相关 start ******/
	RedisUserToken    = ":token"         // 用户Token 缓存数据
	RedisUserInfo     = ":info"          // 用户基本信息 缓存数据
	RedisDeviceBind   = ":bind:device"   // 用户绑定的路由器 缓存数据
	RedisAnnounceRead = ":announce:read" // 用户公告阅读记录
	/**** 用户相关 end   ******/


	/**** 路由器相关 start ******/
	RedisWsMidFix                     = ":mid"       					// websocket请求参数mid缓存键
	RedisRouterInfo                   = ":info:base"					// 路由器SN与Mac缓存键前缀
	RedisRouterPasswordInfo           = ":info:password"				// 路由器密码数据
	RedisRouterWIFIInfo               = ":info:wifi"					// WIFI信息缓存
	RedisRouterLimitSpeed             = ":info:limitSpeed"              // 路由器限速
	RedisStorageDownloadPath          = ":storage:partitionDPath"       // 外接设备下载目录
	RedisSearchLocalFilesCacheKey     = ":temp:marker:searchLocal:"     // 搜索本地资源文件列表
	RedisSearchLocalFilesPageCacheKey = ":temp:marker:searchLocal:page" // 搜索本地资源文件列表
	/**** 路由器相关 end ******/

)
