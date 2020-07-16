package e

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	ERROR:          "Sever Error",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
