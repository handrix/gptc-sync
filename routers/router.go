package routers

import (
	"github.com/gin-gonic/gin"
	"gptc-sync/pkg/setting"
	v1 "gptc-sync/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("block_number", v1.CurrentNodeBlockNumber)
		apiv1.GET("nodes", v1.TotalNodeList)
		apiv1.GET("hash_rate", v1.HashRate)
	}

	return r
}
