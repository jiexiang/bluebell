package controller

type ResCode int64

const (
	UserInfo = "UserInfo"
)

const (
	CodeServerBusy ResCode = 1000 + iota
)

var codeMsgMap = map[ResCode]string{
	CodeServerBusy: "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
