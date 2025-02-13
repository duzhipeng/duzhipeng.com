package account

import (
	"core/ent"
	"core/locale"
	"core/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PostAccountRecordField  新增时输入校验
type PostAccountRecordField struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PostAccountRecordHandler 新增条目API
func PostAccountRecordHandler(c *gin.Context) {
	// Check input.
	var req PostAccountRecordField
	var res locale.ResponseField
	if err := c.ShouldBindJSON(&req); err != nil {
		res = locale.Result.BadInput
		res.Message = "内容有误或格式错误。"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// 核心业务逻辑
	err := CreateAccountRecord(c, utils.DBClient, req)
	if err != nil {
		res = locale.Result.BadInput
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = locale.Result.Ok
	res.Message = "新账户已创建。"
	c.JSON(http.StatusOK, res)
}

// CreateAccountRecord 新增账户
func CreateAccountRecord(ctx *gin.Context, client *ent.Client, req PostAccountRecordField) error {
	_, err := client.Account.Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		Save(ctx)
	if err != nil {
		log.Println("新增账户数据失败。", err.Error())
		return fmt.Errorf("新增账户数据失败")
	}
	return nil
}
