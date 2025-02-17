package utils

import (
	"bot/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"log"
)

var SMSClient *sms.Client
var SMSRequest *sms.SendSmsRequest

func InitTencentSMS(c config.Configs) {
	/* 必要步骤：
	 * 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	 * 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	 * 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	 * 以免泄露密钥对危及你的财产安全。
	 * CAM密匙查询: https://console.cloud.tencent.com/cam/capi*/
	credential := common.NewCredential(
		// os.Getenv("TENCENTCLOUD_SECRET_ID"),
		// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		c.SMS.SecretID,
		c.SMS.SecretKey,
	)
	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()

	/* SDK默认使用POST方法。
	 * 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求 */
	cpf.HttpProfile.ReqMethod = "POST"

	/* SDK有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	// cpf.HttpProfile.ReqTimeout = 5

	/* 指定接入地域域名，默认就近地域接入域名为 sms.tencentcloudapi.com ，也支持指定地域域名访问，例如广州地域的域名为 message.ap-guangzhou.tencentcloudapi.com */
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	/* SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段 */
	cpf.SignMethod = "TC3-HMAC-SHA256"

	/* 实例化要请求产品(以sms为例)的client对象
	 * 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量 */
	/* 实例化要请求产品(以sms为例)的client对象
	 * 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，支持的地域列表参考 https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8 */
	SMSClient, _ = sms.NewClient(credential, "ap-guangzhou", cpf)

	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	 * 你可以直接查询SDK源码确定接口有哪些属性可以设置
	 * 属性可能是基本类型，也可能引用了另一个数据结构
	 * 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明 */

	SMSRequest = sms.NewSendSmsRequest()

	/* 基本类型的设置:
	 * SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	 * SDK提供对基本类型的指针引用封装函数
	 * 帮助链接：
	 * 短信控制台: https://console.cloud.tencent.com/smsv2
	 * 腾讯云短信小助手: https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81 */

	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	SMSRequest.SmsSdkAppId = common.StringPtr("1400614553")

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看 */
	SMSRequest.SignName = common.StringPtr("先验科技")
	/* 国际/港澳台短信 SenderId: 国内短信填空，默认未开通，如需开通请联系 [message helper] */
	SMSRequest.SenderId = common.StringPtr("")
}

func SendSMSAtTencentCloud(phoneNumber []string, tid string, content []string) error {
	/* 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	SMSRequest.SessionContext = common.StringPtr("")
	/* 短信码号扩展号: 默认未开通，如需开通请联系 [message helper] */
	SMSRequest.ExtendCode = common.StringPtr("")
	/* 模板参数: 若无模板参数，则设置为空*/
	SMSRequest.TemplateParamSet = common.StringPtrs(content)
	/* 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看 */
	SMSRequest.TemplateId = common.StringPtr(tid)
	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	SMSRequest.PhoneNumberSet = common.StringPtrs(phoneNumber)

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := SMSClient.SendSms(SMSRequest)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Println("An API error has returned: ", err)
		return fmt.Errorf("an API error has returned")
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		log.Println("No API error has returned:  ", err)
		return fmt.Errorf("an API error has returned")
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	log.Printf("%s\n", b)
	log.Println(phoneNumber, tid, content)
	return nil
}
