package tasks

import "housekeeper/internal/task/workers"

func Start()  {
	go workers.DealMsgCidStatus()

	go workers.ClearPushedMsgs()

	go workers.ValidCidTask()

	go workers.CleanPushServer()
}
