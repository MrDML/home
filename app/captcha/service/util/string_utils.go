package util

import (
	"math/rand"
	"strings"
	"time"
)


// GetRandNumber 获取数字
func GetRandNumber(size int) string{
	if size > 9 || size < 0 {
		size = 6
	}

	rand.Seed(time.Now().Unix())
	var collection  = []string{"1","2","3","4","5","6","7","8","9"}
	res := make([]string,size)
	for i := 0; i < size; i++ {
		index := rand.Intn(size)
		res[i] = collection[index]
	}
	return strings.Join(res,"")
}
