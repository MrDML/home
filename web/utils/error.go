package utils

type errorCode string

const (
	RECODE_OK        errorCode = "0"
	RECODE_DBERR     errorCode = "4001"
	RECODE_NODATA    errorCode = "4002"
	RECODE_DATAEXIST errorCode = "4003"
	RECODE_DATAERR   errorCode = "4004"

	RECODE_SESSIONERR errorCode = "4101"
	RECODE_LOGINERR   errorCode = "4102"
	RECODE_PARAMERR   errorCode = "4103"
	RECODE_USERONERR  errorCode = "4104"
	RECODE_ROLEERR    errorCode = "4105"
	RECODE_PWDERR     errorCode = "4106"
	RECODE_USERERR    errorCode = "4107"
	RECODE_SMSERR     errorCode = "4108"
	RECODE_MOBILEERR  errorCode = "4109"

	RECODE_REQERR    errorCode = "4201"
	RECODE_IPERR     errorCode = "4202"
	RECODE_THIRDERR  errorCode = "4301"
	RECODE_IOERR     errorCode = "4302"
	RECODE_SERVERERR errorCode = "4500"
	RECODE_UNKNOWERR errorCode = "4501"
)

var recodeText = map[errorCode]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_USERONERR:  "用户已经注册",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_SMSERR:     "短信失败",
	RECODE_MOBILEERR:  "手机号错误",
}

func RecodeText(code errorCode) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
