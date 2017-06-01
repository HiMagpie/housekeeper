package main

import (
	"runtime"
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/rpc_srv"
	"housekeeper/internal/http_srv"
	"housekeeper/internal/task"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Run rpc server
	go rpc_srv.Start(cfg.C.RpcSrv.Port)

	// Run http server
	go http_srv.Start(cfg.C.HttpSrv.Port)

	// Start background task dealing
	go tasks.Start()

	select {}
}
