package app

import "gin-server-api/internal/environment"

func Env() environment.Environment {
	return environment.GetEnv()
}

func InitEnv(env string) error {
	return environment.SetEnvString(env)
}
