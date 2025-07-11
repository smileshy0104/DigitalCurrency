package bc

import (
	"bytes"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// ReverseBytes 反转字节数组
// 参数:
//
//	data: 需要反转的字节切片
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// Base58Encode 将字节数组编码为Base58格式字符串
// 参数:
//
//	input: 待编码的原始字节切片
//
// 返回值:
//
//	[]byte: Base58编码结果的字节表示
func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	// 通过大数运算进行Base58转换
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	// 反转结果字节序列
	ReverseBytes(result)

	// 处理前导零字节: 每个前导零在Base58中表示为字母'1'
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

// Base58Decode 将Base58编码字符串解码为原始字节数组
// 参数:
//
//	input: Base58编码的字节切片
//
// 返回值:
//
//	[]byte: 解码后的原始字节切片
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	// 统计前导'1'字符数量(对应原始数据的前导零)
	for _, b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	// 跳过前导字符处理有效载荷
	payload := input[zeroBytes:]

	// 将Base58字符转换为大整数
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	// 注意: 原始实现中前导零处理被注释，实际使用需考虑恢复前导零
	// decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)
	return decoded
}
