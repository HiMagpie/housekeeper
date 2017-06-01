package rpc_srv

import (
	"net/rpc"
	"housekeeper/internal/rpc_srv/proc"
	"net"
	"log"
	"housekeeper/internal/com/logger"
	"time"
	"net/rpc/jsonrpc"
	"fmt"
)

// Run rpc server and register receivers
func Start(port int) {

	srv := rpc.NewServer()
	proc.AddProc(srv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("rpc Listen err: ", err.Error())
	}

	// Deal requests
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Error("rpc.accept", logger.Format(
					"err", err.Error(),
				))
				time.Sleep(time.Millisecond * 100)
				continue
			}

			go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()

	logger.Info("rpc_srv", "rpc on port: " , port)
}
