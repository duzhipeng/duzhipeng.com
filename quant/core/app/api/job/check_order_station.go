package job

import (
	"core/ent"
	"core/ent/station"
	"core/locale"
	"core/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCheckOrderStationRecordField  输入校验
type GetCheckOrderStationRecordField struct {
}

// GetCheckOrderStationRecordHandler 查询API
func GetCheckOrderStationRecordHandler(c *gin.Context) {
	// Check input.
	var req GetCheckOrderStationRecordField
	var res locale.ResponseField
	if err := c.ShouldBindQuery(&req); err != nil {
		res = locale.Result.BadInput
		res.Message = "提交的内容有误。请核实后重新提交。"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// 核心业务逻辑
	data, err := QueryCheckOrderStationRecords(c, utils.DBClient, req)
	if err != nil {
		res = locale.Result.NotFound
		res.Data = err.Error()
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = locale.Result.Ok
	res.Data = data

	c.JSON(http.StatusOK, res)
}

// QueryCheckOrderStationRecords 查询账户数据
func QueryCheckOrderStationRecords(ctx *gin.Context, client *ent.Client, req GetCheckOrderStationRecordField) (any, error) {
	allOrder := client.Order.Query().AllX(ctx)
	for _, o := range allOrder {
		station, err := client.Station.Query().Where(station.AgencyId(o.DispatchedStationId)).First(ctx)
		if err != nil || station == nil {
			return nil, fmt.Errorf("不匹配,%q,%q,%d", o.VehiclePlateNo, o.MaintOrderNo, o.DispatchedStationId)
		}
	}
	return nil, nil
}
