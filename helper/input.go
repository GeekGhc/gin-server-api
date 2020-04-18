package helper


// GetParam
func GetParam(ctx *gin.Context, name string, def ...string) string {
	if value, exists := ctx.GetQuery(name); exists {
		return value
	}

	if value, exists := ctx.GetPostForm(name); exists {
		return value
	}

	if len(def) > 0 {
		return def[0]
	} else {
		return ""
	}
}
