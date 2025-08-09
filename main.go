package main

import (
	"log"
	"net/http"

	"github.com/Djonsinere/web_api_DeepSeek_chat.git/src"
	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func CheckErr(err error) {
	if err != nil {
		log.Print("[[main]]: ", err)
	}
}

func chatHandler(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userMsg := req.Message

	resp, err := src.Ai_send_request("user", userMsg)
	CheckErr(err)

	var reply string

	if err == nil {
		reply = resp
	} else {
		reply = "Error, try again later"
	}

	c.JSON(http.StatusOK, ChatResponse{Reply: reply})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.POST("/api/chat", chatHandler)

	r.Run(":8080")
}
