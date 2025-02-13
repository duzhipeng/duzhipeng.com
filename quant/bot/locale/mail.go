package locale

var MailMode = map[string]string{
	"ORDINARY":   "平信",
	"REGISTERED": "挂号信",
	"AIR_MAIL":   "国际航空挂号信",
	"SF_EXPRESS": "顺丰速运",
	"JDL":        "京东物流",
}

var MailStatus = map[string]string{
	"PENDING":          "确认中",
	"RECEIVED":         "寄存中",
	"PROCESSING":       "处理中",
	"AWAITING_PAYMENT": "待支付",
	"READY_TO_SEND":    "即将发出",
	"SEND_OUT":         "已发出",
	"DESTROYED":        "已销毁",
	"CANCELLED":        "已取消",
}
