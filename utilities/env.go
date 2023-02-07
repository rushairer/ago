package utilities

import (
	"os"
)

// 获取环境变量值，并设置默认值
func GetEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	os.Setenv(key, defaultValue)
	return defaultValue
}
