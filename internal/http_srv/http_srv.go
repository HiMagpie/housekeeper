package http_srv

import (
	"housekeeper/internal/http_srv/router"
	"net/http"
	"os"
	"fmt"
	"housekeeper/internal/com/logger"
)

// Run http server and set api handlers
func Start(port int) {

	router.AddRouters()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("HTTP ListenAndServe err: ", err)
			os.Exit(1)
		}
	}()

	logger.Info("http_srv", "HTTP server on port:" , port)
}
