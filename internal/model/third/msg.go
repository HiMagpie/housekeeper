package third

/**
 * Posted from third-part app for pushing.
 */
type Msg struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Logo    string `json:"logo"`
	Url     string `json:"url"`
	MsgConf MsgConf `json:"msg_conf"`
}

type MsgConf struct {
	Ring      int `json:"ring"`
	Vibrate   int `json:"vibrate"`
	Cleanable int `json:"cleanable"`
	Trans     int `json:"trans"`    // 传输方式: 1透传,2notify
	Begin     string `json:"begin"` // 发出通知的开始结束时间
	End       string `json:"end"`
}

func NewMsg() *Msg {
	return new(Msg)
}
