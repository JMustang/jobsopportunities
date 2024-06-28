package handler

import (
	"net/http"

	"github.com/JMustang/jobsopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]

func ListOpeningHandler(ctx *gin.Context) {
	openins := []schemas.Opening{}

	if err := db.Find(&openins).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openins)
}
