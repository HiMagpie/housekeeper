package proc

import "net/rpc"

func AddProc(srv *rpc.Server) {
	srv.Register(NewReg())
}
