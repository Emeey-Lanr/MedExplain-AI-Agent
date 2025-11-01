package routes

import (
	"ai-agent/handlers"

	"github.com/gin-gonic/gin"
)

func AgentRoute (rg *gin.RouterGroup){
   rg.POST("/medic",  handlers.Inquire)
}