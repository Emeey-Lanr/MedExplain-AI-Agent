package main

import (
	"ai-agent/config"
	"ai-agent/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main ()  {

	r := gin.Default()

    config.Init()

	if err := godotenv.Load(); err != nil{
		log.Println("Error loading.env:", err) // change to println in production for your app not crash
	}

	agentRoute := r.Group("/a2a")
	routes.AgentRoute(agentRoute)



	r.Run(":8080")
}
