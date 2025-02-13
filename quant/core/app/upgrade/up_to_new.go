package upgrade

import (
	"core/ent"
	"core/locale"
	"core/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostUpToHandlerRecordField struct {
	GID string `json:"gid"`
}

// ItemField  标准化订单详情字段
type ItemField struct {
	GID      string `json:"gid"`
	Quantity int    `json:"quantity"`
}

func UpToHandler(c *gin.Context) {
	//var req PostUpToHandlerRecordField
	var res locale.ResponseField
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	res = locale.Result.BadInput
	//	res.Message = "提交的内容有误。请核实后重新提交"
	//	c.JSON(http.StatusBadRequest, res)
	//	return
	//}
	// 核心业务逻辑
	err := UpdateProgram(c, utils.DBClient)
	if err != nil {
		fmt.Print(err)
		res = locale.Result.BadInput
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = locale.Result.Ok
	res.Message = "升级脚本执行成功。"
	c.JSON(http.StatusOK, res)
}

func UpdateProgram(ctx *gin.Context, client *ent.Client) error {
	// 仅管理员操作

	return nil
}
