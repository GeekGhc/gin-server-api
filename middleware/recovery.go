package middleware

import "github.com/gin-gonic/gin"

type RecoveryWriter struct {
}

func (w *RecoveryWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func Recovery() gin.HandlerFunc {
	w := new(RecoveryWriter)
	return gin.RecoveryWithWriter(w)
}
