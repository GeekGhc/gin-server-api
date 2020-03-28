package environment

import (
	"fmt"
	"strings"
	"sync"
)

// 支持的环境
const (
	DevEnvironment Environment = "dev"
	QaEnvironment  Environment = "qa"
	PreEnvironment Environment = "pre"
	PrdEnvironment Environment = "prd"
)

// 环境的类型定义
type Environment string

// ----------------------------------------
//  全局运行环境（默认开发环境）
//  由于运行环境在运行期间的唯一性，所以这里使用
//  私有的全局变量
// --
var gEnvironment = DevEnvironment

// ----------------------------------------
//  运行环境变更时互斥锁
// ----------------------------------------
var gEnvironmentWriteMutex = &sync.Mutex{}

// 获取全局的运行环境
func GetEnv() Environment{
	return gEnvironment
}

//设置全局的运行环境
func SetEnv(env Environment) error {
	gEnvironmentWriteMutex.Lock()
	defer gEnvironmentWriteMutex.Unlock()

	if env.IsValid() {
		gEnvironment = env
		return nil
	} else {
		return fmt.Errorf("the given envronemnt %s is invalid\n", env)
	}
}

// 设置全局运行环境的字符串形式
func SetEnvString(env string) error {
	return SetEnv(Environment(strings.ToLower(env)))
}

// 获取全局运行环境的字符串形式
func (e Environment) String() string {
	return string(e)
}

// 检查全局运行环境是否支持并有效
func (e Environment) IsValid() bool {
	switch e {
	case DevEnvironment, QaEnvironment, PreEnvironment, PrdEnvironment:
		return true
	default:
		return false
	}
}

// 检查当前的全局运行环境是否是给定的值
func (e Environment) Is(env Environment) bool {
	return e == env
}

// 检查当前的全局运行环境是都是在给定的值范围内
func (e Environment) In(envs ...Environment) bool {
	for i, j := 0, len(envs); i < j; i++ {
		if e.Is(envs[i]) {
			return true
		}
	}
	return false
}

//检查当前的全局运行环境是否是开发环境
func (e Environment) IsDev() bool {
	return e.Is(DevEnvironment)
}

//检查当前的全局运行环境是否是测试环境
func (e Environment) IsQa() bool {
	return e.Is(QaEnvironment)
}

//检查当前的全局运行环境是否是预发布环境
func (e Environment) IsPre() bool {
	return e.Is(PreEnvironment)
}

//检查当前的全局运行环境是否是生产环境
func (e Environment) IsPro() bool {
	return e.Is(PrdEnvironment)
}
