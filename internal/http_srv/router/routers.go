package router

import (
	"net/http"
	"housekeeper/internal/http_srv/controllers"
)

/**
 * 添加路由
 */
func AddRouters()  {
	// 客户端注册
	http.HandleFunc("/registration/check-in", controllers.CheckIn)
	http.HandleFunc("/registration/get-servers", controllers.GetServers)

	// 第三方发过来的期望推送的消息
	http.HandleFunc("/receiver/recv-multi", controllers.RecvMulti)
}
