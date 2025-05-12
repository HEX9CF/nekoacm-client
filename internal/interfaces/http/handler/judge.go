package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/application/service"
	"nekoacm-client/internal/interfaces/http/vo"
	"net/http"
)

// 提交评测
func JudgeSubmit(c *gin.Context) {
	var req dto.Submission

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	j, err := service.Submit(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", j))
}
