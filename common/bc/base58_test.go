package bc

import (
	"fmt"
	"testing"
)

// TestBase58Encode 测试 Base58 编码和解码功能。
//
// 该测试用例验证以下逻辑：
//  1. 对字节切片 "asdasdasdasdasdasdas" 进行 Base58 编码
//  2. 将编码结果转换为字符串并打印
//  3. 对编码结果进行 Base58 解码
//  4. 将解码后的字节切片转换为字符串并打印
//
// 测试目标：
//   - 验证 Base58Encode 函数能否正确处理输入数据
//   - 验证 Base58Decode 函数能正确还原原始数据
//   - 检查编解码过程的完整性（原始数据与解码后数据应一致）
//
// 参数:
//
//	t : *testing.T  测试框架提供的测试上下文对象
func TestBase58Encode(t *testing.T) {
	// 对测试字符串进行 Base58 编码
	encode := Base58Encode([]byte("asdasdasdasdasdasdas"))
	// 打印编码结果（转换为可读字符串）
	fmt.Println(string(encode))

	// 对编码结果进行 Base58 解码
	decode := Base58Decode(encode)
	// 打印解码结果（应与原始输入一致）
	fmt.Println(string(decode))
}
