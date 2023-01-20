package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkData struct {
	Link string `json:"link"`
}

func DisplayRoot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello to my page",
	})
}

func ProcessRoot(ctx *gin.Context) {
	//fmt.Println(io.ReadAll(ctx.Request.Body))
	var newLink LinkData
	if err := ctx.ShouldBindJSON(&newLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newLink)
}
