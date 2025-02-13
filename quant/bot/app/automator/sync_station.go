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

// SyncStationAutomating 任务名称
const (
	SyncStationAutomating = "Automating:SyncStationAutomating"
)

// SyncStationAutomatingPayload 任务所需数据的结构定义
type SyncStationAutomatingPayload struct {
}

// SyncStationAutomatingTask 任务定时器设置
func SyncStationAutomatingTask() {
	payload, _ := json.Marshal(SyncStationAutomatingPayload{})
	task := asynq.NewTask(SyncStationAutomating, payload)
	// 设置脚本的执行频率
	entryID, err := utils.AsynqScheduler.Register("@every 10s", task)
	if err != nil {
		fmt.Println("SyncStationAutomatingTask, 注册失败, ", err)

	}
	fmt.Println("SyncStationAutomatingTask, 已注册, ", entryID)

}

// HandleSyncStationAutomatingTask 任务算法
func HandleSyncStationAutomatingTask(ctx context.Context, t *asynq.Task) error {
	// 载入Payload
	var payload SyncStationAutomatingPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	// 预取Token
	acc := utils.DBClient.Account.Query().FirstX(ctx)
	token := acc.Token
	// 构造请求客户端
	requ := req.C().SetTimeout(30 * time.Second)
	// 响应解析
	var result StationListRes
	// 定位拉取进度
	pageNum := 1
	pageNumRedis, _ := utils.RedisClient.Get(ctx, "pageNum").Int()
	if pageNumRedis != 0 {
		pageNum = pageNumRedis + 1
	} else {
		err := utils.RedisClient.Set(ctx, "pageNum", 0, 0).Err()
		if err != nil {
			log.Println("缓存异常,%w", err)
		}
	}
	// 拉取数据范围 400，大约13000条数据
	if pageNum <= 413 {
		resp := requ.Post("https://api.shinwell.cn/api/v2/agency/station/list").
			SetContext(ctx).
			SetHeaders(map[string]string{
				"Token": token,
			}).
			SetBodyJsonMarshal(
				map[string]any{
					"pageNum":                    pageNum,
					"pageSize":                   100,
					"channelManagerUserIds":      []string{},
					"cityCode":                   nil,
					"cooperationTypes":           []string{},
					"country":                    nil,
					"fittingDirectSupply":        nil,
					"keywords":                   nil,
					"settlementAccountConfirmed": nil,
					"supportPersonalBusiness":    nil,
				},
			).
			SetSuccessResult(&result).
			Do()
		if resp.IsSuccessState() {
			for _, item := range result.Data.List {
				_, err := utils.DBClient.Station.Create().
					SetAgencyId(item.AgencyId).
					SetAgencyName(item.AgencyName).
					Save(ctx)
				if err != nil {
					log.Println("存储服务站数据异常,%w", err, item.AgencyId, item.AgencyName)
				}
			}
			err := utils.RedisClient.Incr(ctx, "pageNum").Err()
			if err != nil {
				log.Println("缓存异常,%w", err)
			}
			log.Println("第", pageNum, "页数据（每页100条）已入库。")
		} else if resp.Err != nil { // status `code >= 400` is considered as error
			// Must have been marshaled to errMsg if no error returned before
			log.Println("got error:", resp.String(), resp.Dump())

		} else {
			log.Println("unknown http status:", resp.Status, resp)

		}
	}
	return nil
}

// StationListRes 维保单API返回报文解析
type StationListRes struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    StationListResData `json:"data"`
}

// StationListResData 维保单API返回报文解析
type StationListResData struct {
	Total int                    `json:"total"`
	List  StationListResDataList `json:"list"`
}

// StationListResDataList 维保单API返回报文解析
type StationListResDataList []struct {
	AgencyId   int    `json:"agencyId"`
	AgencyName string `json:"agencyName"`
}
