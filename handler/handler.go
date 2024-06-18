package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST Opening!",
	})
}
func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Opening!",
	})
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE Opening!",
	})
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "UPDATE Opening!",
	})
}
func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Opening!",
	})
}