package utilities

import (
	"testing"

	"gotest.tools/assert"
)

// 获取环境变量值，并设置默认值
func Test_GetEnv(t *testing.T) {
	defaultValue := "this value is empty"

	goPath := GetEnv("GOPATH", defaultValue)
	assert.Assert(t, goPath != defaultValue)

	emptyEnvValue := GetEnv("EMPTY_ENV", defaultValue)
	assert.Equal(t, emptyEnvValue, defaultValue)
}
