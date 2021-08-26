package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryChainCode(c *gin.Context)  {
	defer func() {
		if recover() != nil {
			obj :=make(map[string]interface{},16)
			obj["status"] = 400
			obj["msg"] = "出现异常了"
			obj["data"] = nil
			c.JSON(http.StatusBadGateway,obj)
		}
	}()
	panic("自动报错")
}