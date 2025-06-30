package code_gen

import (
	"html/template"
	"log"
	"os"
	"strings"
)

// RpcCommon 存储生成RPC代码所需的公共信息
type RpcCommon struct {
	PackageName string // 包名
	GrpcPackage string // gRPC包名
	ModuleName  string // 模块名
	ServiceName string // 服务名
}

// Rpc 表示一个RPC方法的定义，包含方法名、请求类型和响应类型
type Rpc struct {
	FunName string // 方法名
	Req     string // 请求类型
	Resp    string // 响应类型
}

// RpcResult 包含生成RPC代码所需的所有信息
type RpcResult struct {
	RpcCommon RpcCommon // 公共信息
	Rpc       []Rpc     // RPC方法列表
	ParamList []string  // 参数列表
}

// GenZeroRpc 根据提供的RpcResult对象生成ZeroRPC客户端代码
func GenZeroRpc(result RpcResult) {
	// 创建并解析模板文件 client.tpl
	t := template.New("client.tpl")
	tmpl, err := t.ParseFiles("./client.tpl")
	log.Println(err)

	// 检查 ./gen 目录是否存在，如果不存在则创建它
	_, err = os.Stat("./gen")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./gen", 0666) // 创建目录，权限为0666
		}
	}

	// 初始化参数列表，并填充请求和响应类型
	var pl []string
	for _, v := range result.Rpc {
		// 添加唯一的请求类型到参数列表
		if !isContain(pl, v.Req) {
			pl = append(pl, v.Req)
		}
		// 添加唯一的响应类型到参数列表
		if !isContain(pl, v.Resp) {
			pl = append(pl, v.Resp)
		}
	}
	result.ParamList = pl

	// 创建输出文件
	fileName := "./gen/" + strings.ToLower(result.RpcCommon.ServiceName) + ".go"
	file, err := os.Create(fileName)
	defer file.Close()
	log.Println(err)

	// 执行模板，将数据写入文件
	err = tmpl.Execute(file, result)
	log.Println(err)
}

// isContain 检查字符串str是否存在于字符串切片pl中
func isContain(pl []string, str string) bool {
	for _, p := range pl {
		if p == str {
			return true // 如果存在，返回true
		}
	}
	return false // 否则返回false
}
