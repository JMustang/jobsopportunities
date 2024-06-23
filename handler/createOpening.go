package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
	}
}
