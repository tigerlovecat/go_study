package errno

var (
	// Common errors
	OK            = &Errno{Status: 1, Info: "OK"}
	SverOld       = &Errno{Status: 9, Info: "当前软件版本过低，请及时更新软件"}
	ParamsEmpty   = &Errno{Status: 10, Info: "参数为空"}
	SignError     = &Errno{Status: 11, Info: "签名错误"}
	BanError      = &Errno{Status: 12, Info: "该Ip 加入到黑名单"}
	SmsError      = &Errno{Status: 13, Info: "发生短信错误"}
	SmsLimitError = &Errno{Status: 14, Info: "该Ip 已限制发送短信"}
	EmailError    = &Errno{Status: 15, Info: "发送邮件错误"}
	AckError      = &Errno{Status: 16, Info: "Ack机制错误返回信息"}
	DbFoundError  = &Errno{Status: 17, Info: "数据库查不到数据"}
	ParamsError   = &Errno{Status: 18, Info: "参数异常"}
	GetError      = &Errno{Status: 19, Info: "获取异常"}
	GetNews       = &Errno{Status: 20, Info: "阅读公告消息"}
	TokenError    = &Errno{Status: 299, Info: "登录过期，请重新登录"}

	// 网络数据库错误
	InternalServerError = &Errno{Status: 10001, Info: "Internal server error"}
	ErrValidation       = &Errno{Status: 10002, Info: "Validation failed."}
	ErrDatabase         = &Errno{Status: 10003, Info: "数据库操作失败."}
	ErrLianLe           = &Errno{Status: 10004, Info: "请求盛天接口失败："}
	ErrResourceDetail   = &Errno{Status: 10005, Info: "请求资源详情接口失败："}
	ErrResourceNotFound = &Errno{Status: 10006, Info: "暂无资源"}

	// user errors
	ErrUserNotFound      = &Errno{Status: 20102, Info: "用户名或密码错误"}
	ErrTokenInvalid      = &Errno{Status: 20103, Info: "token 校验失败"}
	ErrPasswordIncorrect = &Errno{Status: 20104, Info: "密码错误"}
	ErrTokenError        = &Errno{Status: 20105, Info: "生成 token 失败"}
	ErrSendSmsError      = &Errno{Status: 20106, Info: "验证码重复发送"}
	ErrRegistered        = &Errno{Status: 20107, Info: "该手机已注册过"}
	ErrUnLogin           = &Errno{Status: 20108, Info: "该手机未注册过"}
	ErrValidRegister     = &Errno{Status: 20109, Info: "校验码已失效，请重新注册"}
	ErrValidFindPassword = &Errno{Status: 20110, Info: "校验码已失效，请重新找回"}
	ErrEmailUnLogin      = &Errno{Status: 20111, Info: "该邮箱未注册过"}
	ErrNickNameUsed      = &Errno{Status: 20112, Info: "用户名重复,不能注册"}
	ErrPassword          = &Errno{Status: 20113, Info: "密码长度必须是6至20位,并且包含大小写字母及数字"}
	ErrSTPassword        = &Errno{Status: 20114, Info: "密码为8到20个字符"}
	ErrRouterWsClose     = &Errno{Status: 20115, Info: "路由器websocket未连接"}
	ErrTokenDel		     = &Errno{Status: 20116, Info: "token 校验失败"}

	// sms errors
	ErrSmsExpire = &Errno{Status: 30001, Info: "验证码已过期"}
	ErrSmsValid  = &Errno{Status: 30002, Info: "验证码校验失效"}

	// email errors
	ErrEmailExpire = &Errno{Status: 31001, Info: "邮件验证码已过期"}
	ErrEmailValid  = &Errno{Status: 31002, Info: "邮件验证码失效"}

	// task errors
	ErrSetTaskCache            = &Err{Status: 40001, Info: "任务信息缓存失败"}
	ErrTaskToQueue             = &Errno{Status: 40002, Info: "添加下载任务进入队列失败"}
	ErrRouterSnValidate        = &Errno{Status: 40003, Info: "SN检测不通过 "}
	ErrDealLoadTask            = &Errno{Status: 40004, Info: "任务不存在或已删除"}
	ErrWaitTaskMax             = &Errno{Status: 40005, Info: "正在下载的任务数过多，请稍后再试"}
	ErrGetWaitTaskNum          = &Errno{Status: 40006, Info: "获取下载中及等待中的任务数失败"}
	ErrTaskPush                = &Errno{Status: 40007, Info: "任务推送失败"}
	ErrGetTaskInfoCache        = &Err{Status: 40008, Info: "获取任务详情缓存失败"}
	ErrAddTaskToLoadingTask    = &Errno{Status: 40009, Info: "添加任务到下载队列失败"}
	ErrDelTaskFromLoadingQueue = &Errno{Status: 40010, Info: "将任务从下载中队列移除失败"}
	ErrDelTaskRequest          = &Errno{Status: 40011, Info: "任务请求中，请稍后再试"}
	ErrPauseTaskRequest        = &Errno{Status: 40012, Info: "只能暂停下载中的任务"}
	ErrResumeTaskRequest       = &Errno{Status: 40013, Info: "任务状态异常"}
	ErrTaskLoadRepeat          = &Errno{Status: 40014, Info: "下载队列中已有此资源"}
	ErrGetTaskRepeat           = &Err{Status: 40015, Info: "根据path查询任务失败"}
	ErrDelTaskUpdateStatus     = &Errno{Status: 40016, Info: "修改任务状态失败"}
	ErrDelTask                 = &Errno{Status: 40017, Info: "删除任务失败"}
	ErrTaskQueuePush           = &Err{Status: 40018, Info: "任务推送失败"}

	// device errors
	ErrDeviceFound          = &Errno{Status: 50001, Info: "设备不存在"}
	ErrDeviceValidate       = &Errno{Status: 50002, Info: "设备存在不合法绑定关系"}
	ErrDeviceMacValidate    = &Errno{Status: 50002, Info: "设备Mac匹配异常"}
	ErrDeviceBind           = &Errno{Status: 50003, Info: "该设备已经绑定过，不能重复绑定"}
	ErrDeviceBindFail       = &Errno{Status: 50004, Info: "设备绑定失败"}
	ErrDeviceUnBind         = &Errno{Status: 50005, Info: "该设备不属于你，不能解绑"}
	ErrSetSpeedLimit        = &Errno{Status: 50006, Info: "设置速度限制失败"}
	ErrSetSpeedLimitTimeout = &Errno{Status: 50007, Info: "设置速度限制超时"}
	ErrDeviceSnMac          = &Errno{Status: 50008, Info: "设备SN和Mac绑定不合法"}
	ErrSearchLocal          = &Errno{Status: 50009, Info: "下载模块搜索本地资源失败"}
	ErrDeviceUid            = &Errno{Status: 50010, Info: "该设备不属于你"}
	ErrDeviceGetInfo        = &Err{Status: 50011, Info: "获取连接设备详情失败"}
	ErrSSID		            = &Errno{Status: 50012, Info: "路由器名称不能超过10个中文"}

	// storage  errors
	ErrPushStorageScan    = &Errno{Status: 60001, Info: "下发扫描目录命令失败"}
	ErrGetStoragePathInfo = &Errno{Status: 60002, Info: "分页获取文件列表失败"}
	ErrDelStoragePathI    = &Errno{Status: 60003, Info: "删除文件失败"}
	ErrGetStoragePar      = &Errno{Status: 60004, Info: "该设备没有外接存储"}
	ErrDefaultPath        = &Errno{Status: 60005, Info: "外接设备数据异常，请重新插拔外接设备"}
	ErrSetDefaultPath     = &Errno{Status: 60006, Info: "设置分区失败"}
	ErrGetDeviceInfo      = &Errno{Status: 60007, Info: "获取设备信息失败"}
	ErrGetTaskInfo        = &Errno{Status: 60008, Info: "获取任务信息失败"}
	ErrDelLoadingTask     = &Errno{Status: 60009, Info: "不能删除正在下载的任务文件"}
	ErrNullStorage	      = &Errno{Status: 60010, Info: "该设备没有外接存储"}
	ErrRouterConnection	  = &Errno{Status: 60011, Info: "路由器已关闭连接"}
)
