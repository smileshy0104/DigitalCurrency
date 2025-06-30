package code_gen

import "testing"

// TestGenStruct 测试GenStruct函数，用于验证结构体代码生成的正确性。
// 该测试用例通过调用GenStruct函数来生成一个名为"Member"的结构体代码。
func TestGenStruct(t *testing.T) {
	GenModel("member", "Member")
}

// TestGenRpc 测试GenRpc函数，用于验证RPC相关代码生成的正确性。
// 该测试用例设置了RPC服务的公共信息和具体的RPC方法信息，并调用GenZeroRpc函数生成相应的代码。
func TestGenRpc(t *testing.T) {
	// 初始化RPC服务的公共信息，如包名、模块名等。
	rpcCommon := RpcCommon{
		PackageName: "mclient",
		ModuleName:  "Market",
		ServiceName: "Market",
		GrpcPackage: "market",
	}

	// 下面三行是RPC方法的示例，展示了RPC方法的定义方式，但并未实际使用。
	// rpc FindSymbolThumb(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolThumbTrend(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolInfo(MarketReq) returns(ExchangeCoin);

	// 定义一个具体的RPC方法，包括方法名、请求和响应的消息类型。
	rpc1 := Rpc{
		FunName: "FindSymbolThumbTrend",
		Resp:    "SymbolThumbRes",
		Req:     "MarketReq",
	}

	// 可以继续定义其他RPC方法，但目前被注释掉了。
	//rpc2 := Rpc{
	//	FunName: "FindSymbolThumbTrend",
	//	Resp:    "SymbolThumbRes",
	//	Req:     "MarketReq",
	//}
	//rpc3 := Rpc{
	//	FunName: "FindSymbolInfo",
	//	Resp:    "ExchangeCoin",
	//	Req:     "MarketReq",
	//}

	// 创建一个RPC方法列表，并将定义的RPC方法添加到列表中。
	rpcList := []Rpc{}
	rpcList = append(rpcList, rpc1)

	// 构建包含RPC服务公共信息和RPC方法列表的结果对象。
	result := RpcResult{
		RpcCommon: rpcCommon,
		Rpc:       rpcList,
	}

	// 调用GenZeroRpc函数生成RPC相关的代码。
	GenZeroRpc(result)
}
