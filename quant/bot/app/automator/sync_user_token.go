package automator

import (
	"bot/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/imroc/req/v3"
	"log"
	"time"
)

// SyncUserTokenAutomating 任务名称
const (
	SyncUserTokenAutomating = "Automating:SyncUserTokenAutomating"
)

// SyncUserTokenAutomatingPayload 任务所需数据的结构定义
type SyncUserTokenAutomatingPayload struct {
}

// SyncUserTokenAutomatingTask 任务定时器设置
func SyncUserTokenAutomatingTask() {
	payload, _ := json.Marshal(SyncUserTokenAutomatingPayload{})
	task := asynq.NewTask(SyncUserTokenAutomating, payload)
	// 设置脚本的执行频率
	//entryID, err := utils.AsynqScheduler.Register("@every 5s", task) // 调测用
	entryID, err := utils.AsynqScheduler.Register("@every 30s", task)
	if err != nil {
		fmt.Println("SyncUserTokenAutomatingTask, 注册失败, ", err)

	}
	fmt.Println("SyncUserTokenAutomatingTask, 已注册, ", entryID)

}

// HandleSyncUserTokenAutomatingTask 任务算法
func HandleSyncUserTokenAutomatingTask(ctx context.Context, t *asynq.Task) error {
	// 载入Payload
	var payload SyncUserTokenAutomatingPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	// 预取账号
	usr := utils.DBClient.Account.Query().FirstX(ctx)
	// 构造请求客户端
	requ := req.C().SetTimeout(30 * time.Second)
	// 响应解析
	var result LoginRes
	resp := requ.Post("https://api.shinwell.cn/api/v2/user/login").
		SetBodyJsonMarshal(
			map[string]string{
				"username": usr.Username,
				"password": usr.Password,
			},
		).
		SetSuccessResult(&result).
		Do()
	if resp.IsSuccessState() {
		err := utils.RedisClient.Set(ctx, "TOKEN", result.Data, time.Hour*8).Err()
		if err != nil {
			log.Println(err)
			return fmt.Errorf("缓存系统异常。")
		}
	} else if resp.Err != nil { // status `code >= 400` is considered as error
		// Must have been marshaled to errMsg if no error returned before
		log.Println("got error:", resp.String())
		fmt.Println("got error:", resp.String())
	} else {
		log.Println("unknown http status:", resp.Status)
		fmt.Println("unknown http status:", resp.Status)
	}
	return nil
}

// LoginRes 登录鉴权API返回报文解析
type LoginRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	TraceId string `json:"traceId"`
}
