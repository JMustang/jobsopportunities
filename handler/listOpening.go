package handler

import (
	"net/http"

	"github.com/JMustang/jobsopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openins := []schemas.Opening{}

	if err := db.Find(&openins).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openins)
}
