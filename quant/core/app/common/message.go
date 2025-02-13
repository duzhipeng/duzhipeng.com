package common

import (
	"context"
	"core/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/thanhpk/randstr"
	"log"
	"time"
)

// SendSMS 处理并发送短信
func SendSMS(ctx *gin.Context, enum string, phone []string, content ...string) error {
	switch enum {
	case "VERIFY": // 发送验证码。验证码短信不进队列，直接发出
		// 检查是否已存在待验证的验证码
		singlePhone := phone[0] // 仅支持单发
		val, err := utils.RedisClient.Exists(ctx, singlePhone+":message:verify").Result()
		if err != nil {
			log.Println(err)
			return fmt.Errorf("验证码缓存模块异常。")
		}
		if val == 1 {
			return fmt.Errorf("验证码正在发送途中。请耐心等待3分钟后再试。")
		}
		// 生成及缓存验证码
		token := randstr.Dec(4)
		// ID：1978590 你的验证码为：V-{1}，有效期：10分钟。若非你本人操作请忽略不管。
		msgTempId := "1978590"
		contentStr := []string{token}
		err = utils.SendSMSAtTencentCloud([]string{singlePhone}, msgTempId, contentStr)
		if err != nil {
			return fmt.Errorf("短信通道异常。")
		}
		// 验证码过期时间统一为10分钟
		err = utils.RedisClient.Set(ctx, singlePhone+":message:token", token, 10*time.Minute).Err()
		if err != nil {
			return fmt.Errorf("验证码已过期。请重新申请。")
		}
		// 验证码重发时间统一为180秒
		err = utils.RedisClient.Set(ctx, singlePhone+":message:verify", token, 180*time.Second).Err()
		if err != nil {
			return fmt.Errorf("验证码正在发送途中。请耐心等待3分钟后再试。")
		}
		return nil
	default:
		return nil
	}
}

// SMSCaptchaVerify 验证码核验算法
func SMSCaptchaVerify(phone string, token string) (bool, error) {
	ctx := context.Background()
	// 获取验证码
	val, err := utils.RedisClient.Get(ctx, phone+":message:token").Result()
	if errors.Is(err, redis.Nil) {
		return false, fmt.Errorf("验证码错误。")
	} else if err != nil {
		log.Println(err)
		return false, fmt.Errorf("验证码模块异常。")
	}
	// 比对验证码
	if val == token {
		return true, nil
	} else {
		return false, fmt.Errorf("验证码错误。")
	}
}
