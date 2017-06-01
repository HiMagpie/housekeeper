package models

/**
 * 第三方应用发过来的Payload格式
 */
type Payload struct {
	Cid string `json:"cid"`
	Ring bool `json:"ring"`
	Vibrate bool `json:"vibrate"`
	Cleanable bool `json:"cleanable"`
	Trans int `json:"trans"` // 传输方式: 1透传,2notify
	Begin string `json:"begin"` // 发出通知的开始结束时间
	End string `json:"end"`

	Title string `json:"title"`
	Text string `json:"text"`
	Logo string `json:"logo"`
	Url string `json:"url"`

	// 备注信息
	MsgId int64 `json:"msg_id"`
	Ctime int64 `json:"ctime"`
}
