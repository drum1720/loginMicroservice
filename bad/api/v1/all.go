package v1

import (
	"github.com/gin-gonic/gin"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/bad/api/v1/middlewares"
	"loginMicroservice/bad/api/v1/models"
	"loginMicroservice/bad/api/v1/response"
	"time"
)

func AllApi(r *gin.Engine) {
	group := r.Group("v1")

	group.Use(middlewares.UserMiddleware)

	group.POST("start", func(c *gin.Context) {
		userId := c.GetInt64("UserId")
		res, err := posgresql.GetDbPoolConnect().Exec(c, "INSERT INTO \"TimeLog\" (\"UserId\", \"StartTime\", \"EndTime\") VALUES ($1,$2,NULL)", userId, time.Now())
		if err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 100))
			return
		}
		timeLogId := res.RowsAffected()
		c.JSON(200, timeLogId)
	})

	group.POST("end", func(c *gin.Context) {
		dbPoolConnects := posgresql.GetDbPoolConnect()
		userId := c.GetInt64("UserId")

		querryGET := `SELECT id, "UserId", "StartTime", "IsEnded" 
FROM "TimeLog" 
WHERE "UserId" = $1 
AND "EndTime" IS NULL 
AND "IsEnded" = false 
ORDER BY id DESC 
LIMIT 1`
		var record models.TimeLogEntry
		res := dbPoolConnects.QueryRow(c, querryGET, userId)
		err := res.Scan(&record.Id, &record.UserId, &record.StartTime, &record.IsEnded)

		if err != nil || record.Id == 0 {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 100))
			return
		}

		QuerryUPDATE := `UPDATE "TimeLog"
SET "EndTime" = $1,
    "IsEnded" = $2
WHERE id = $3`
		record.EndTime = time.Now()
		record.IsEnded = true
		record.CalculateMinutes()
		if _, err = dbPoolConnects.Exec(c, QuerryUPDATE, record.EndTime, record.IsEnded, record.Id); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 101))
			return
		}

		c.JSON(200, record)

	})
}
