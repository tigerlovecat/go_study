package e

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"math"
	"math/big"
	"math/rand"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"web_framework/config"
	"web_framework/pkg/constvar"
	"web_framework/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"

	yp "github.com/yunpian/yunpian-go-sdk/sdk"

	log "github.com/sirupsen/logrus"
	"web_framework/model"
	)

// 封装日志调用
func Log(logType string, args ...interface{}) {

	// 如果正式上线之后，系统会将日志以每天分一个文件存在根目录下面的文件 eg：log/error/20190114.log
	// 创建日志目录
	if config.C.RunMode == "release" {
		logPath := "log/"
		if logType == "err" || logType == "error" || logType == "Err" {
			logPath += "error/"
		}else {
			logPath += "info/"
		}
		res, err := MakeDir(logPath)
		if err != nil || !res {
			log.Error("创建日志目录 [log/] 失败")
		}
		dateStr := time.Now().Format("20060102")
		filePtah := logPath + dateStr + ".log"
		file, err := os.OpenFile(filePtah, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0666)
		if err == nil {
			DelLogPath()
			log.SetOutput(file)
		}
	}

	switch logType {
	case "info", "Info":
		log.Info(args...)
	case "err", "error", "Err":
		log.Error(args...)
	case "bug", "debug":
		log.Debug(args...)
	case "params", "param":
		params, _ := json.Marshal(args)
		log.Info("参数为：", string(params))
	default:
		log.Warn(args...)
	}
}

func DelLogPath() error {
	logPath := "log/"
	dataStr := time.Now().AddDate(0, 0, -7).Format("20060102")
	delErrPath :=  logPath + "error/" + dataStr + ".log"
	delInfoPath :=  logPath + "info/" + dataStr + ".log"
	os.RemoveAll(delErrPath)
	os.RemoveAll(delInfoPath)
	return nil
}

// 创建目录
func MakeDir(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	// 不存在目录就创建目录
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUrlName(url string) (name, ext string) {
	filenameWithSuffix := path.Base(url)
	fileSuffix := path.Ext(filenameWithSuffix)
	filename := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	lowerSuffix := strings.ToLower(fileSuffix)
	return filename, lowerSuffix
}

// 校验APP版本号
func CheckVersion(currentSver string) error {

	if config.C.RunMode == "debug" {
		return nil
	}
	if currentSver == "" {
		return errno.SverOld
	}
	sVer, _ := strconv.Atoi(currentSver)
	if sVer < config.C.SVer {
		return errno.SverOld
	}
	return nil
}

// 校验P4Pclient版本号
func CheckP4PVersion(currentSver string) error {

	if config.C.RunMode == "debug" {
		return nil
	}
	if currentSver == "" {
		return errno.SverOld
	}
	sVer, _ := strconv.Atoi(currentSver)
	if sVer < config.C.P4pclientSver {
		return errno.SverOld
	}
	return nil
}

// MD5 加密
func Md5Str(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 校验签名 app 端
func CheckSign(sign string, args... string) error {

	if config.C.RunMode == "debug" {
		return nil
	}

	var signString, md5Str string
	for _, v := range args {
		signString += v
	}
	md5Str = Md5Str(signString)
	if md5Str == sign || strings.ToUpper(md5Str) == sign {
		return nil
	}

	return errno.SignError
}

// trim 掉多余的空格和换行符
func Trim(str string) (newStr string) {

	// 将字符串的转换成[]rune
	strList := []rune(str)
	lth := len(strList)
	star := 0
	end := lth - 1
	for i := 0; i < lth; i++ {
		if star == i {
			if string(strList[i:i+1]) == " " {
				star++
			}
		} else {
			if string(strList[i:i+1]) == " " {
				end = i
			}
		}
	}

	if star < end {
		newStr = string(strList[star:end])
	}
	return
}

// 参数为空校验
func CheckEmptyParams(c *gin.Context, checkList []string) (err error, params map[string]string) {

	// 校验为空的 参数
	var emptyCheck []string

	params = make(map[string]string)
	for _, v := range checkList {
		if c.Query(v) == "" {
			emptyCheck = append(emptyCheck, v)
		} else {
			params[v] = c.Query(v)
		}
	}

	if len(emptyCheck) > 0 {
		errno.ParamsEmpty.Info = " 参数为空: {" + strings.Join(emptyCheck, ",") + "}"
		Log("Warn", errno.ParamsEmpty.Info)
		return errno.ParamsEmpty, params
	}

	return nil, params
}

// 获取当前时间的微秒数
func MicroTime() int64 {
	return time.Now().UnixNano() / 1000000
}

// 生成随机数
func RandInt(min, max int) int {

	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// 邮箱校验
func IsEmail(email string) bool {
	matched, _ := regexp.MatchString("\\w[-\\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\\.)+[A-Za-z]{2,14}$", email)
	return matched
}

// 手机号校验
func IsPhone(phone string) bool {
	matched, _ := regexp.MatchString("0?(13|14|15|18)[0-9]{9}", phone)
	return matched
}

// 校验密码是否正确
func CheckPassword(check, password, salt string) bool {
	if password == Md5Str(Md5Str(check)+salt) {
		return true
	}
	return false
}

// 人人账号的密码规则：密码长度必须是6至20位,并且包含大小写字母及数字
func YYetsPasswordCheck(password string) bool {

	if len(password) > 20 || len(password) < 6 {
		fmt.Println("[Err] YYets Password Check: Error 0")
		return false
	}

	reg := regexp.MustCompile("[\\d]+?")
	if ! reg.MatchString(password) {
		fmt.Println("[Err] YYets Password Check: Error 1")
		return false
	}

	reg = regexp.MustCompile("[a-z]+?")
	if ! reg.MatchString(password) {
		fmt.Println("[Err] YYets Password Check:Error 2")
		return false
	}

	reg = regexp.MustCompile("[A-Z]+?")
	if ! reg.MatchString(password) {
		fmt.Println("[Err] YYets Password Check:Error 3")
		return false
	}

	return true
}

// 盛天路由器管理密码校验
func STPasswordCheck(password string) bool {

	if len(password) > 20 || len(password) < 8 {
		fmt.Println("[Err] YYets Password Check: Error 4")
		return false
	}

	reg := regexp.MustCompile("[\\w]+?")
	if ! reg.MatchString(password) {
		fmt.Println("[Err] YYets Password Check: Error 5")
		return false
	}

	return true
}

// 生成密码
func CreatePassword(password, salt string) string {
	return Md5Str(Md5Str(password) + salt)
}

func GetRandomNum(lenght int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	result = []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机字符串
func GetRandomStr(lenght int) string {
	str := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	bytes := []byte(str)
	var result []byte
	result = []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// IP地址转化为整型
func IpToInt(ip string) int64 {

	ret := big.NewInt(0)
	matched, _ := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", ip)
	if matched == false {
		return ret.Int64()
	}

	ret.SetBytes(net.ParseIP(ip).To4())
	if ret.Int64() > 2147483647 {
		return 2147483647
	}
	return ret.Int64()
}

// 整型IP地址转化为192.168.1.1格式
func IpToString(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// 生成唯一ID
func GetUuid() string {
	Uuid := uuid.Must(uuid.NewV4()).String()
	Uuid = strings.Replace(Uuid, "-", "", -1)
	return Uuid
}

// 发送短信 【人人影视】 【YYeTs】
func SendSMS(phone string, area int) (int, error) {

	// 生成随机 短信码
	codeNum := RandInt(1000, 9999)
	code := strconv.Itoa(codeNum)

	var smsCode string
	if area == 86 || area == 852 || area == 886 || area == 853 {
		smsCode = "您的验证码是" + code
	} else if area == 82 {
		smsCode = "당신의 인증코드는 " + code + " 입니다"
	} else if area == 81 {
		smsCode = "あなたの認証番号は " + code
	} else {
		smsCode = "Your verification code is " + code
	}

	client := yp.New(config.C.YpApiKey)
	param := yp.NewParam(2)
	param[yp.MOBILE] = phone
	param[yp.TEXT] = smsCode
	r := client.Sms().SingleSend(param)

	if r.Code == 0 {
		return codeNum, nil
	} else {
		return 0, errno.SmsError
	}
}

func EmailFindPassWord(email, code string) error {
	emailHtml := "\r\n" +
		"<p>尊敬的用户：</p> \r\n" + "<p></p>" +
		"<p align='left'>&nbsp;&nbsp;&nbsp;您好，收到这个邮件是因为有人使用您的邮箱进行密码找回。如果不是您本人操作，请忽略。<br/> 输入该验证码修改帐号密码：<strong>" + code + "</strong></p>" +
		"<p align='left'> </p>" +
		"<p align='left'>字幕组非常感谢您的使用。</p>"

	err := SendExMail(email, emailHtml, "字幕组,邮箱密码找回")
	return err
}

func SendExMail(to, email, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.C.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", email)
	d := gomail.NewDialer(config.C.Email.Host, config.C.Email.Port, config.C.Email.User, config.C.Email.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}

	return nil
}

func PointFormat(point int64) (str string) {

	i := 0
	Units := map[int]string{0: "", 1: "万", 2: "亿", 3: "兆"}

Loop:
	for {
		if math.Abs(float64(point)) >= 10000 {
			point = point / 10000
			i++
			if i == 3 {
				break Loop
			}
		} else {
			break Loop
		}
	}

	newSize := ParseString(float64(point), 0)
	str = string(newSize) + Units[i]

	return
}

// 字节数转GB TB单位
func SizeFormat(byteSizeInt64 int64, pre int) (str string) {

	i := 0
	Units := map[int]string{0: "B", 1: "KB", 2: "MB", 3: "GB", 4: "TB", 5: "PB"}
	byteSize := float64(byteSizeInt64)
Loop:
	for {
		if math.Abs(byteSize) >= 1024 {
			byteSize = byteSize / 1024
			i++
			if i == 5 {
				break Loop
			}
		} else {
			break Loop
		}
	}

	newSize := ParseString(float64(byteSize), pre)
	str = string(newSize) + Units[i]

	return
}

// 速度转化
func SpaceFormat(byteSize float64) (str string) {

	i := 0
	Units := map[int]string{0: "B/s", 1: "KB/s", 2: "MB/s", 3: "G/s", 4: "T/s"}

Loop:
	for {
		if math.Abs(byteSize) >= 1024 {
			byteSize = byteSize / 1024
			i++
			if i == 4 {
				break Loop
			}
		} else {
			break Loop
		}
	}
	newSize := fmt.Sprintf("%.1f", byteSize)
	str = string(newSize) + Units[i]

	return
}

// 相对路径转绝对图片路径  2018/1224/b_273cde81d79dc4851bade79ca0cb0061.jpg
func GetPicPath(url string) string {

	matched, _ := regexp.MatchString("^((https|http|ftp|rtsp|mms)?:\\/\\/)[^\\s]+", url)
	if !matched {
		return constvar.PicPathPrefix + url
	}
	return url
}

//获取当前时间到明天凌晨的秒数差
func GetNowToNextDayUnix() (int64, error) {

	//获取明天凌晨的时间字符串
	nextDayStr := time.Now().Format("2006-01-02") + " 23:59:59"
	//时间字符串格式化
	nextDayTime, err := time.Parse("2006-01-02 15:04:05", nextDayStr)
	if err != nil {
		return 0, err
	}
	nextDayUnix := nextDayTime.Unix()
	//秒数差
	unixDiff := nextDayUnix - time.Now().Unix()

	return unixDiff, nil
}

// 时间格式转为时间戳
func StrToTime(t string) (unixTime int64) {
	timeTemp, err := time.Parse(constvar.TimeFormatS, t)
	if err == nil {
		unixTime = timeTemp.Unix() - constvar.TimeFormatCST
	}
	return
}

// 时间格式转为周
func StrTimeToWeek(t string) (result string) {
	var timeStr string
	timeStr = t + " 00:00:00"
	week := time.Unix(StrToTime(timeStr), 0).Weekday().String()
	switch week {
	case "Sunday":
		result = t + " 周日"
	case "Monday":
		result = t + " 周一"
	case "Tuesday":
		result = t + " 周二"
	case "Wednesday":
		result = t + " 周三"
	case "Thursday":
		result = t + " 周四"
	case "Friday":
		result = t + " 周五"
	case "Saturday":
		result = t + " 周六"
	default:
		result = t
	}

	return
}

// 时间戳转为时间字符串
func TimeToStr(t int64) (strTime string) {
	strTime = time.Unix(t, 0).Format(constvar.TimeFormatS)
	return
}

// 时间戳 输出为当前时间前 eg: 2分钟之前
func TimeFormatShow(t int64, tail string) (strTime string) {
	nowTime := time.Now().Unix()
	diff := nowTime - t
	if diff < constvar.TimeFormatHour { // 分钟 < 60 * 60
		strTime = strconv.FormatInt(diff/constvar.TimeFormatMin, 10) + "分钟" + tail
	} else if diff < constvar.TimeFormatDay { // 小时 < 60 * 60 * 24
		strTime = strconv.FormatInt(diff/constvar.TimeFormatHour, 10) + "小时" + tail
	} else if diff < constvar.TimeFormatWeek { // 天   < 60 * 60 * 24 * 7
		strTime = strconv.FormatInt(diff/constvar.TimeFormatDay, 10) + "天" + tail
	} else if diff < constvar.TimeFormatMonth { // 周   < 60 * 60 * 24 * 30
		strTime = strconv.FormatInt(diff/constvar.TimeFormatWeek, 10) + "周" + tail
	} else if diff < constvar.TimeFormatYear { // 月   < 60 * 60 * 24 * 365
		strTime = strconv.FormatInt(diff/constvar.TimeFormatMonth, 10) + "月" + tail
	} else if diff < constvar.TimeFormatTenYear { // 年   < 60 * 60 * 24 * 365 * 10
		strTime = strconv.FormatInt(diff/constvar.TimeFormatYear, 10) + "年" + tail
	} else { // 盘古开荒前
		strTime = "盘古开荒" + tail
	}

	return
}

// 时间戳 输出为当前时间前 (仅一天以内)
func TimeFormatOneDay(t int64, tail string) (strTime string) {
	nowTime := time.Now().Unix()
	diff := nowTime - t
	if diff < constvar.TimeFormatHour { // 分钟 < 60 * 60
		strTime = strconv.FormatInt(diff/constvar.TimeFormatMin, 10) + "分钟" + tail
	} else if diff < constvar.TimeFormatDay { // 小时 < 60 * 60 * 24
		strTime = strconv.FormatInt(diff/constvar.TimeFormatHour, 10) + "小时" + tail
	} else {
		strTime = TimeToStr(t)
	}

	return
}

// 时间字符串 输出为当前时间前 eg: 2分钟之前
func DateTimeToStr(dateTime string, showType int) (returnTimeStr string) {
	dateTime = strings.Replace(dateTime, "T", " ", -1)
	returnTimeStr = strings.Replace(dateTime, "+08:00", "", -1)
	if showType == 1 {
		return returnTimeStr
	}
	timeData := StrToTime(returnTimeStr)
	returnTimeStr = TimeFormatShow(timeData, "前")
	return returnTimeStr
}

//速度单位转换
func SpeedUnitConversion(speed float64) string {

	//if speed < 1024*1024{
	//	string := strconv.FormatFloat( speed / 1024, 'f', 1, 64)
	//	return string + "Kb/s"
	//}else{
	//	string := strconv.FormatFloat( speed / (1024*1024), 'f', 1, 64)
	//	return string + "Mb/s"
	//}
	return SpaceFormat(speed)
}

//速度单位转换
func StorageSpaceConversion(space int64) string {
	if space > 1024*1024 && space <= 1024*1024*1024 {
		return strconv.FormatInt(space/int64(1024*1024), 10) + "M"
	} else if space > 1024*1024*1024 && space <= 1024*1024*1024*1024 {
		return strconv.FormatInt(space/int64(1024*1024*1024), 10) + "G"
	} else if space > 1024*1024*1024*1024 {
		return strconv.FormatInt(space/int64(1024*1024*1024*1024), 10) + "T"
	} else {
		return "小于1M"
	}
}

//获取剧集名称
func GetZimuzuSeason(channal string, season, episode int) string {
	var (
		seasonStr  string
		episodeStr string
	)
	if model.TaskChannelMovie != channal {
		if episode != 0 {
			episodeStr = "第" + strconv.Itoa(episode) + "集"
		}
		switch season {
		case 0:
			seasonStr = "前传"
		case 101:
			seasonStr = "单剧"
		case 102:
			seasonStr = "MINI剧"
		case 103:
			seasonStr = "周边资源"
		default:
			seasonStr = "第" + strconv.Itoa(season) + "季"
		}

		return seasonStr + " " + episodeStr
	}else{
		return ""
	}
}

//根据URL获取资源名称
func GetResNameByUrl(url string)string{
	var (
		channal string
		season , epsiode int
	)
	urlReg := "(N=)(.*?)(\\|)"
	reg1 := regexp.MustCompile(urlReg)
	result := reg1.FindString(url)
	if len(result) == 0{
		return url
	}
	result = strings.Replace(result, "N=", "", 1 )
	nameSilce := strings.Split(result , ".")

	seasonEpisodeReg := "(?i)S(\\d+)\\s*E(\\d+)"
	reg2 := regexp.MustCompile(seasonEpisodeReg)
	seasonEpisode := reg2.FindString(url)
	if len(seasonEpisode) == 0 {
		return nameSilce[0]
	}

	seasonEpisodeNumReg := "\\d+"
	reg3 := regexp.MustCompile(seasonEpisodeNumReg)
	seasonEpisodeNum := reg3.FindAllString(result,2)

	season, _ = strconv.Atoi(seasonEpisodeNum[0])
	epsiode, _ = strconv.Atoi(seasonEpisodeNum[1])
	seasonStr := GetZimuzuSeason(channal , season , epsiode)

	resName := nameSilce[0] + " " + seasonStr

	return resName
}

// float64 转字符串保留两个小数
func ParseString(float642 float64, pre int) string {
	formatStr := "%.2f"
	if pre <= 0 {
		return strconv.Itoa(int(math.Floor(float642 + 0/5)))
	}else {
		formatStr = "%."+strconv.Itoa(pre)+"f"
	}
	return fmt.Sprintf(formatStr, float642)
}

func TransPinYin(str string) string {
	var trans string
	dict := pinyin.NewDict()
	trans = dict.Convert(str, "").None()
	return trans
}

func ReplaceString(str string) string {

	src := strings.Replace(str, "\\n", " ", -1)
	return src
}