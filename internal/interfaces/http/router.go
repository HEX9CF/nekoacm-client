package http

import (
	"github.com/gin-gonic/gin"
	"nekoacm-client/internal/interfaces/http/handler"
	"nekoacm-client/internal/interfaces/http/middlewares"
	"nekoacm-client/internal/interfaces/http/vo"
	"nekoacm-client/pkg/config"
	"net/http"
)

func InitRoute() error {
	conf := config.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, vo.RespOk("NekoACM 启动成功", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, vo.RespError("404 Not Found", nil))
	})

	// 使用中间件
	ginServer.Use(middlewares.TokenAuth())

	// 初始化路由
	apiRoute := ginServer.Group("/api")
	{
		apiRoute.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, vo.RespOk("NekoACM 服务可用", nil))
		})
	}

	initProblemRoute(apiRoute)
	initTestcaseRoute(apiRoute)
	initSolutionRoute(apiRoute)
	initJudgeRoute(apiRoute)
	initChatRoute(apiRoute)
	initMiscRoute(apiRoute)

	// 启动服务
	err := ginServer.Run(":" + conf.Port)
	if err != nil {
		return err
	}
	return nil
}

func initProblemRoute(rg *gin.RouterGroup) {
	pr := rg.Group("/problem")

	pr.POST("/parse", handler.ParseProblem)
	pr.POST("/translate", handler.TranslateProblem)
	pr.POST("/generate", handler.GenerateProblem)
}

func initTestcaseRoute(rg *gin.RouterGroup) {
	tc := rg.Group("/testcase")

	tc.POST("/generate", handler.GenerateTestcase)
}

func initSolutionRoute(rg *gin.RouterGroup) {
	s := rg.Group("/solution")

	s.POST("/generate", handler.GenerateSolution)
}

func initJudgeRoute(rg *gin.RouterGroup) {
	j := rg.Group("/judge")

	j.POST("/submit", handler.JudgeSubmit)
}

func initChatRoute(rg *gin.RouterGroup) {
	c := rg.Group("/chat")

	c.POST("/assistant", handler.ChatAssistant)
}

func initMiscRoute(rg *gin.RouterGroup) {
	m := rg.Group("/misc")

	m.GET("/joke", handler.GenerateJoke)
}
