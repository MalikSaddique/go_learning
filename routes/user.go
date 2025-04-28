package routes

import (
	"fmt"
	"net/http"

	"github.com/MalikSaddique/go_learning/models"
	"github.com/MalikSaddique/go_learning/utils"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm/utils"
)

// protected handler
// SaveResult godoc
// @Summary      Protected endpoint for saving analyzed data
// @Description  Requires JWT. Analyzes a text file and saves the result to the database.
// @Tags         protected
// @Security     BearerAuth
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      401
// @Failure      500
// @Router       /protected [get]
func (r *Router) SaveResult(c *gin.Context) {
	var req *models.Result

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Failed to bind JSON: %v\n", err)
		fmt.Println("Received JSON body:", c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := r.UserService.SaveResults(c, req)

	if result == nil {
		return
	}
	c.JSON(200, result)
}

// GetResult godoc
// @Summary      Get Analysis Result (Paginated)
// @Description  Retrieves paginated analyzed text data from the database for the given user ID
// @Tags         protected
// @Security     BearerAuth
// @Produce      json
// @Param        user_id path string true "User ID"
// @Param        page query int false "Page number"
// @Param        limit query int false "Limit per page"
// @Success      200  "Result data retrieved"
// @Failure      404  "Result not found"
// @Failure      500  "Database connection failed"
// @Router       /getdata/{user_id} [get]
func (r *Router) GetResult(c *gin.Context) {
	page, limit := utils.PaginationHandler(c)

	results, err := r.UserService.GetUserResults(c, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(404, gin.H{"error": "No results found"})
		return
	}

	c.JSON(200, gin.H{
		"page":    page,
		"limit":   limit,
		"results": results,
	})
}
