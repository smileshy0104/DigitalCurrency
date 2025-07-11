package bc

import (
	"fmt"
	"testing"
)

// TestWallet_GetAddress 测试 Wallet 的地址获取功能
//
// 该测试用例验证以下行为：
//  1. 成功创建新钱包实例
//  2. 正确获取钱包的测试地址
//  3. 正确获取钱包的私钥
//
// 测试流程：
//   - 初始化新钱包
//   - 调用 GetTestAddress() 获取地址并打印
//   - 调用 GetPriKey() 获取私钥并打印
//
// 注意：该测试主要验证核心功能可用性，实际使用时应检查输出格式和内容
func TestWallet_GetAddress(t *testing.T) {
	// 初始化新钱包
	wallet, err := NewWallet()
	if err != nil {
		panic(err)
	}
	// 获取钱包的测试地址
	address := wallet.GetTestAddress()
	fmt.Println(string(address))
	priKey := wallet.GetPriKey()
	fmt.Println(priKey)
}
