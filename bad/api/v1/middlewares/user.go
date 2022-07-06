package middlewares

import (
	"github.com/gin-gonic/gin"
	"loginMicroservice/bad/api/v1/response"
	"net/http"
)

func UserMiddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId int64 `binding:"required"`
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse("bad user supplied", 2))
		return
	}

	//todo:validation user need

	c.Set("UserId", req.UserId)
}
