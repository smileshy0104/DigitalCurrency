package tools

import (
	"fmt"
	"math/rand"
)

// Rand4Num 生成4位随机数——作为验证码
func Rand4Num() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn = intn + 1000
	}
	return fmt.Sprintf("%d", intn)
}
