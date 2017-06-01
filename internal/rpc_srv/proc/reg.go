package proc

import (
	"housekeeper/internal/cache"
	"housekeeper/internal/com/logger"
	"housekeeper/internal/com/cfg"
)

// For magpie(push server) registering
type Reg struct{}

type RegArgs struct {
	Hostname string
	Ip       string
	Port     string
}

type RegResponse struct {
	PushServerHb int
}

func NewReg() *Reg {
	return new(Reg)
}

// Register push server info when push server starts
func (this *Reg) RegPushServer(args *RegArgs, res *RegResponse) error {
	err := cache.NewPushServerCache().AddUnique(args.Ip, args.Port)
	if err != nil {
		logger.Error("rpc.reg.push.server", logger.Format(
			"err", err.Error(),
			"hostname", args.Hostname,
			"ip", args.Ip,
			"port", args.Port))
		return err
	}

	res.PushServerHb = cfg.C.RpcSrv.PushServerHb
	return nil
}


