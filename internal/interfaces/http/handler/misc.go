package handler

import (
	"github.com/gin-gonic/gin"
	"nekoacm-client/internal/application/service"
	"nekoacm-client/internal/interfaces/http/vo"
	"net/http"
)

// 生成笑话
func GenerateJoke(c *gin.Context) {
	p, err := service.TellJoke()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}
