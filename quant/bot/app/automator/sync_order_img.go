package automator

import (
	"bot/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/imroc/req/v3"
	"github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"time"
)

// SyncOrderImgAutomating 任务名称
const (
	SyncOrderImgAutomating = "Automating:SyncOrderImgAutomating"
)

// SyncOrderImgAutomatingPayload 任务所需数据的结构定义
type SyncOrderImgAutomatingPayload struct {
}

// SyncOrderImgAutomatingTask 任务定时器设置
func SyncOrderImgAutomatingTask() {
	payload, _ := json.Marshal(SyncOrderImgAutomatingPayload{})
	task := asynq.NewTask(SyncOrderImgAutomating, payload)
	// 设置脚本的执行频率
	entryID, err := utils.AsynqScheduler.Register("@every 10s", task)
	if err != nil {
		fmt.Println("SyncOrderImgAutomatingTask, 注册失败, ", err)

	}
	fmt.Println("SyncOrderImgAutomatingTask, 已注册, ", entryID)

}

// HandleSyncOrderImgAutomatingTask 任务算法
func HandleSyncOrderImgAutomatingTask(ctx context.Context, t *asynq.Task) error {
	// 常量设置
	//ossUrl := "https://oss.shinwell.cn/"        // 图片服务器地址
	filePath := "/Users/adu/Downloads/666/3/"  // 任务文件夹地址
	fileName := "中通智运数字科技有限公司_202403-对账单.xlsx" // 文件名
	sheetName := "明细表"

	// 读取文件及获取单号
	f, err := excelize.OpenFile(filePath + fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 确定表格行数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rowsNum := len(rows)

	// 拿到指定列的全部值
	cellData := make(map[string]string)
	for num := range rowsNum {
		rowNum := strconv.Itoa(num)
		cell, err := f.GetCellValue(sheetName, "A"+rowNum)
		if err != nil {
			fmt.Println(err)
		}
		if len(cell) == 18 {
			cellData[cell] = rowNum
		}
	}
	fmt.Println("全部的单号：", cellData)

	// 载入Payload
	var payload SyncOrderImgAutomatingPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	// 预取Token
	acc := utils.DBClient.Account.Query().FirstX(ctx)
	token := acc.Token
	// 构造请求客户端
	requ := req.C().SetTimeout(30 * time.Second)
	// 响应解析
	var result OrderImgListRes

	// 按单号拉取相册照片到本地
	//imgUrlList := make([]string, rowsNum)
	for orderNum, rowNum := range cellData {
		getUrl := "https://api.shinwell.cn/api/v2/maint/gallery/" + orderNum + "/WORK_IMAGE/media/list"
		resp := requ.Get(getUrl).
			SetContext(ctx).
			SetHeaders(map[string]string{
				"Token": token,
			}).SetSuccessResult(&result).
			Do()
		if resp.IsSuccessState() {
			for index, item := range result.Data.MediaList {
				_, err = requ.R().SetOutputFile(filePath + orderNum + "/" + item.MediaKey + ".jpg").Get("https://oss.shinwell.cn/" + item.MediaKey)
				if resp.IsSuccessState() {
					// 图片插入指定的列
					file, err := os.ReadFile(filePath + orderNum + "/" + item.MediaKey + ".jpg")
					if err != nil {
						fmt.Println(err)
						return err
					}
					if err := f.AddPictureFromBytes(sheetName, "S"+rowNum, &excelize.Picture{
						Extension: ".jpg",
						File:      file,
						Format:    &excelize.GraphicOptions{AutoFit: true},
					}); err != nil {
						fmt.Println(err)
						return err
					}

				} else {
					log.Println("出错:", orderNum, index)
					return err
				}
			}
		} else if resp.Err != nil { // status `code >= 400` is considered as error
			// Must have been marshaled to errMsg if no error returned before
			log.Println("got error:", resp.String(), resp.Dump())

		} else {
			log.Println("unknown http status:", resp.Status, resp)

		}
	}
	// 保存文件
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	log.Println("一轮完成！")
	return nil
}

// OrderImgListRes 维保单API返回报文解析
type OrderImgListRes struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    OrderImgResData `json:"data"`
}

// OrderImgResData 维保单API返回报文解析
type OrderImgResData struct {
	MediaList OrderMediaListResData `json:"mediaList"`
}

// OrderMediaListResData 维保单API返回报文解析
type OrderMediaListResData []struct {
	MediaKey string `json:"mediaKey"`
}
