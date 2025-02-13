package app

import (
	"bot/app/automator"
	"github.com/hibiken/asynq"
)

// InitTaskMonitor 初始化客户端：添加所有定时任务
func InitTaskMonitor() {
	// 注册定时任务
	//automator.SyncUserTokenAutomatingTask()
	//automator.SyncStationAutomatingTask()
	automator.SyncOrderImgAutomatingTask()
}

// InitTaskMux 初始化服务端：注册所有任务算法
func InitTaskMux() *asynq.ServeMux {
	// mux
	mux := asynq.NewServeMux()
	// 注册任务算法
	//mux.HandleFunc(automator.SyncUserTokenAutomating, automator.HandleSyncUserTokenAutomatingTask)
	//mux.HandleFunc(automator.SyncStationAutomating, automator.HandleSyncStationAutomatingTask)
	mux.HandleFunc(automator.SyncOrderImgAutomating, automator.HandleSyncOrderImgAutomatingTask)
	return mux
}
