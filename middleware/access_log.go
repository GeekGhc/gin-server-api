package middleware

import (
	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		path := ctx.Request.URL.Path
		//start := time.Now()
		ctx.Next()

		// 跳过健康检查
		if path == "/ping" {
			return
		}

		//stop := time.Since(start)
		//latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		//hostName, err := os.Hostname()
		//if err != nil {
		//	hostName = "unknown"
		//}
		//app.Log().WithFields(logrus.Fields{
		//	"latency":       latency,
		//	"hostname":      hostName,
		//	"method":        ctx.Request.Method,
		//	"status_code":   ctx.Writer.Status(),
		//	"client_ip":     ctx.ClientIP(),
		//	"user_agent":    ctx.Request.UserAgent(),
		//	"referer":       ctx.Request.Referer(),
		//	"data_length":   ctx.Writer.Size(),
		//	"path":          path,
		//	"response_size": ctx.Writer.Size(),
		//	"content_type":  ctx.GetHeader("Content-Type"),
		//	"params":        helper.GetParams(ctx),
		//}).Info("access log")
	}
}
