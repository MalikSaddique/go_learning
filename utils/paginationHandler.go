package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationHandler(c *gin.Context) (page int, limit int) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, _ = strconv.Atoi(pageStr)
	limit, _ = strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return
}
