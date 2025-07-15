package domain

// CaptchaDomain 验证码验证的业务逻辑结构体
type CaptchaDomain struct {
}

// NewCaptchaDomain 创建一个新的CaptchaDomain实例
func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}

// vaptchaReq 验证码验证请求的结构体
type vaptchaReq struct {
	Id        string `json:"id"`
	Secretkey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}

// vaptchaRsp 验证码验证响应的结构体
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

// Verify 验证验证码的正确性。
// 该方法通过发送POST请求到服务器，验证给定的验证码(token)是否有效。
// 参数:
//
//	server: 验证码服务的URL。
//	vid: 验证码的ID，用于标识验证码。
//	key: 安全密钥，用于服务器验证请求的合法性。
//	token: 客户端生成的验证码token。
//	scene: 验证码场景，用于区分不同的验证码使用场景。
//	ip: 客户端的IP地址，用于增加额外的安全验证层。
//
// 返回值:
//
//	验证码验证结果，true表示验证成功，false表示验证失败。
func (d *CaptchaDomain) Verify(
	server string,
	vid string,
	key string,
	token string,
	scene int,
	ip string) bool {
	//// 发送一个post请求
	//resp, err := tools.Post(server, &vaptchaReq{
	//	Id:        vid,
	//	Secretkey: key,
	//	Token:     token,
	//	Scene:     scene,
	//	Ip:        ip,
	//})
	//if err != nil {
	//	logx.Error(err)
	//	return false
	//}
	result := &vaptchaRsp{}
	//err = json.Unmarshal(resp, result)
	//if err != nil {
	//	logx.Error(err)
	//	return false
	//}
	// 判断验证结果是否成功
	return result.Success == 1
}
