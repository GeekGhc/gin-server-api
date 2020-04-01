package http

import (
	"gin-server-api/internal/logger"
	"github.com/astaxie/beego/validation"
)

// logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logger.Info(err.Key,err.Message)
	}
	return
}
