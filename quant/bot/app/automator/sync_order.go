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

// SyncOrderAutomating 任务名称
const (
	SyncOrderAutomating = "Automating:SyncOrderAutomating"
)

// SyncOrderAutomatingPayload 任务所需数据的结构定义
type SyncOrderAutomatingPayload struct {
}

// SyncOrderAutomatingTask 任务定时器设置
func SyncOrderAutomatingTask() {
	payload, _ := json.Marshal(SyncOrderAutomatingPayload{})
	task := asynq.NewTask(SyncOrderAutomating, payload)
	// 设置脚本的执行频率
	entryID, err := utils.AsynqScheduler.Register("@every 15s", task)
	if err != nil {
		fmt.Println("SyncOrderAutomatingTask, 注册失败, ", err)

	}
	fmt.Println("SyncOrderAutomatingTask, 已注册, ", entryID)

}

// HandleSyncOrderAutomatingTask 任务算法
func HandleSyncOrderAutomatingTask(ctx context.Context, t *asynq.Task) error {
	// 载入Payload
	var payload SyncOrderAutomatingPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	// 预取Token
	acc := utils.DBClient.Account.Query().FirstX(ctx)
	token := acc.Token
	// 构造请求客户端
	requ := req.C().SetTimeout(30 * time.Second)
	// 响应解析
	var result OrderListRes
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
	// 拉取数据范围 202，大约10000条数据
	if pageNum <= 203 {
		resp := requ.Post("https://api.shinwell.cn/api/v2/maint/maint-order/list").
			SetContext(ctx).
			SetHeaders(map[string]string{
				"Token": token,
			}).
			SetBodyJsonMarshal(
				map[string]any{
					"pageNum":                pageNum,
					"pageSize":               50,
					"vehicleTeamCountryCode": "CHN",
					"maintOrderStatusList":   []string{"QUOTE_ACCEPTED"},
					"from":                   1,
					"filterType":             "ALL",
					"countOnly":              false,
				},
			).
			SetSuccessResult(&result).
			Do()
		if resp.IsSuccessState() {
			for _, item := range result.Data.List {
				_, err := utils.DBClient.Order.Create().
					SetMaintOrderNo(item.MaintOrderNo).
					SetVehiclePlateNo(item.VehiclePlateNo).
					SetVehicleTeamName(item.VehicleTeamName).
					SetMaintRequestType(item.MaintRequestType).
					SetStationName(item.StationName).
					SetDispatchedStationId(item.DispatchedStationId).
					Save(ctx)
				if err != nil {
					log.Println("存储维保单数据异常,%w", err, item.MaintOrderNo, item.VehiclePlateNo)
				}
			}
			err := utils.RedisClient.Incr(ctx, "pageNum").Err()
			if err != nil {
				log.Println("缓存异常,%w", err)
			}
			log.Println("第", pageNum, "页数据（每页50条）已入库。")
		} else if resp.Err != nil { // status `code >= 400` is considered as error
			// Must have been marshaled to errMsg if no error returned before
			log.Println("got error:", resp.String(), resp.Dump())

		} else {
			log.Println("unknown http status:", resp.Status, resp)

		}
	}
	return nil
}

// OrderListRes 维保单API返回报文解析
type OrderListRes struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    OrderListResData `json:"data"`
}

// OrderListResData 维保单API返回报文解析
type OrderListResData struct {
	Total int                  `json:"total"`
	List  OrderListResDataList `json:"list"`
}

// OrderListResDataList 维保单API返回报文解析
type OrderListResDataList []struct {
	MaintOrderNo        string `json:"maintOrderNo"`
	VehiclePlateNo      string `json:"vehiclePlateNo"`
	VehicleTeamName     string `json:"vehicleTeamName"`
	MaintRequestType    string `json:"maintRequestType"` // MAINT_RESERVATION 预约维保 EMERGENCY_RESCUE 紧急有缘
	DispatchedStationId int    `json:"dispatchedStationId"`
	StationName         string `json:"stationName"`
}
