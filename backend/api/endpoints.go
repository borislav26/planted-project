package api

import (
	"backend/builder"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTotalIntersection(services *builder.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		totalOverlap, overlapPrecentage, err := services.MainService.TotalOverlap()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, map[string]any{
			"totalOverlap":      totalOverlap,
			"overlapPercentage": overlapPrecentage,
		})
	}
}

func getCommonZipcodes(services *builder.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		zipcodes, err := services.MainService.MostCommonZipcodes()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, zipcodes)
	}
}

func getOverlapBetweenGroups(services *builder.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		combinations, err := services.MainService.AllCombinationsOverlap()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, combinations)
	}
}

func getAllUsersWithNumberOfElements(services *builder.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := services.MainService.AllUsersWithNumberOfElements()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
