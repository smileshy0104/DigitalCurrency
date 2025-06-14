package tools

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

// 默认常量定义盐值长度、迭代次数和密钥长度
const (
	defaultSaltLen    = 64    // 默认盐值长度为64字节
	defaultIterations = 10000 // 默认PBKDF2迭代次数为10000次
	defaultKeyLen     = 128   // 默认生成的密钥长度为128字节
)

// 默认使用的哈希算法为 SHA-512
var defaultHashFunction = sha512.New

// Options 结构体用于自定义 PBKDF2 参数，包括盐值长度、迭代次数、密钥长度和哈希函数
type Options struct {
	SaltLen      int              // 自定义盐值长度
	Iterations   int              // 自定义迭代次数
	KeyLen       int              // 自定义密钥长度
	HashFunction func() hash.Hash // 自定义哈希函数（如 sha512.New）
}

// generateSalt 生成指定长度的随机盐值，使用字母数字字符
func generateSalt(length int) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	salt := make([]byte, length)
	rand.Read(salt)
	for key, val := range salt {
		salt[key] = alphanum[val%byte(len(alphanum))]
	}
	return salt
}

// Encode 对原始密码进行加密编码
// 参数:
//
//	rawPwd string - 原始密码
//	options *Options - 自定义参数选项，如果为 nil 则使用默认参数
//
// 返回值:
//
//	string - 生成的盐值
//	string - 加密后的密码（十六进制格式）
func Encode(rawPwd string, options *Options) (string, string) {
	if options == nil {
		// 使用默认配置
		salt := generateSalt(defaultSaltLen)
		encodedPwd := pbkdf2.Key([]byte(rawPwd), salt, defaultIterations, defaultKeyLen, defaultHashFunction)
		return string(salt), hex.EncodeToString(encodedPwd)
	}
	// 使用自定义配置
	salt := generateSalt(options.SaltLen)
	encodedPwd := pbkdf2.Key([]byte(rawPwd), salt, options.Iterations, options.KeyLen, options.HashFunction)
	return string(salt), hex.EncodeToString(encodedPwd)
}

// Verify 验证原始密码是否与已加密的密码匹配
// 参数:
//
//	rawPwd string - 原始密码
//	salt string - 之前生成的盐值
//	encodedPwd string - 已加密的密码（十六进制格式）
//	options *Options - 自定义参数选项，如果为 nil 则使用默认参数
//
// 返回值:
//
//	bool - 是否匹配
func Verify(rawPwd string, salt string, encodedPwd string, options *Options) bool {
	if options == nil {
		// 使用默认配置验证
		return encodedPwd == hex.EncodeToString(pbkdf2.Key([]byte(rawPwd), []byte(salt), defaultIterations, defaultKeyLen, defaultHashFunction))
	}
	// 使用自定义配置验证
	return encodedPwd == hex.EncodeToString(pbkdf2.Key([]byte(rawPwd), []byte(salt), options.Iterations, options.KeyLen, options.HashFunction))
}
