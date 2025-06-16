package middlewares

import (
	"alist/internal/conf"
	"alist/internal/errs"
	"alist/internal/setting"
	"alist/server/common"

	"github.com/gin-gonic/gin"
)

func SearchIndex(c *gin.Context) {
	mode := setting.GetStr(conf.SearchIndex)
	if mode == "none" {
		common.ErrorResp(c, errs.SearchNotAvailable, 500)
		c.Abort()
	} else {
		c.Next()
	}
}
